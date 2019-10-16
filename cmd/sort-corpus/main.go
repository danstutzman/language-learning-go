package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const PARSE_DIR = "db/1_parses"

// Summary of a constituent
type Summary struct {
	firstTokenBegin int
	sexp            []string
	verbTypes       []string
}

func main() {
	if len(os.Args) != 1+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: path to corpus (.yaml or .csv or .txt file)`)
	}
	corpusPath := os.Args[1]

	var phrases []string
	if strings.HasSuffix(corpusPath, ".yaml") {
		phrases = parsing.ListPhrasesInCorpusYaml(corpusPath)
	} else if strings.HasSuffix(corpusPath, ".csv") {
		phrases = parsing.ListPhrasesInCorpusCsv(corpusPath)
	} else if strings.HasSuffix(corpusPath, ".txt") {
		phrases = parsing.ListPhrasesInCorpusTxt(corpusPath)
	} else {
		log.Fatalf("Unrecognized extension for path '%s'", corpusPath)
	}

	for _, phrase := range phrases {
		output := parsing.LoadSavedParse(phrase, PARSE_DIR)

		summary := summarizePhrase(output.Phrase, output.Parse)

		//fmt.Printf("%-20s %s\n", strings.Join(summary.sexp, ""), phrase)

		if hasNonEasyVerbType(summary.verbTypes) {
			fmt.Printf("%-20s %s\n", strings.Join(summary.verbTypes, " "), phrase)
		}
	}
}

var EASY_VERB_TYPES = map[string]bool{
	"AI": true,
	"MG": true,
	"MI": true,
	"MN": true,
	"MP": true,
	"SI": true,
}

func hasNonEasyVerbType(verbTypes []string) bool {
	for _, verbType := range verbTypes {
		isEasy := EASY_VERB_TYPES[verbType]
		if !isEasy {
			return true
		}
	}
	return false
}

func summarizePhrase(phrase string, parse parsing.Parse) Summary {
	summaries := []Summary{}
	for _, sentence := range parse.Sentences {
		tokenById := map[string]parsing.Token{}
		for _, token := range sentence.Tokens {
			tokenById[token.Id] = token
		}

		for _, constituent := range sentence.Constituents {
			summaries = append(summaries,
				summarizeConstituent(constituent, tokenById, true)...)
		}
	}

	sort.SliceStable(summaries, func(i, j int) bool {
		return summaries[i].firstTokenBegin < summaries[j].firstTokenBegin
	})

	superSummary := Summary{
		firstTokenBegin: summaries[0].firstTokenBegin,
		sexp:            concatSexps(summaries, true),
		verbTypes:       concatVerbTypes(summaries),
	}
	return superSummary
}

func concatSexps(summaries []Summary, isSentence bool) []string {
	concattedSexps := []string{}
	if !isSentence && len(summaries) > 1 {
		concattedSexps = append(concattedSexps, "(")
	}
	for _, summary := range summaries {
		concattedSexps = append(concattedSexps, summary.sexp...)
	}
	if !isSentence && len(summaries) > 1 {
		concattedSexps = append(concattedSexps, ")")
	}
	return concattedSexps
}

func concatVerbTypes(summaries []Summary) []string {
	allVerbTypes := []string{}
	for _, summary := range summaries {
		allVerbTypes = append(allVerbTypes, summary.verbTypes...)
	}
	return allVerbTypes
}

func summarizeConstituent(constituent parsing.Constituent,
	tokenById map[string]parsing.Token, isSentence bool) []Summary {

	if len(constituent.Children) == 0 {
		token := tokenById[constituent.Token]
		shortTag := token.Tag[0:1]

		verbTypes := []string{}
		if strings.HasPrefix(token.Tag, "V") {
			verbTypes = []string{token.Tag[1:3]}
		}

		if shortTag == "F" { // discard punctuation
			return []Summary{}
		}
		return []Summary{{
			firstTokenBegin: mustAtoi(token.Begin),
			sexp:            []string{shortTag},
			verbTypes:       verbTypes,
		}}
	}

	summaries := []Summary{}

	for _, child := range constituent.Children {
		summaries = append(summaries,
			summarizeConstituent(child, tokenById, false)...)
	}

	if len(summaries) == 0 {
		return []Summary{}
	}

	sort.SliceStable(summaries, func(i, j int) bool {
		return summaries[i].firstTokenBegin < summaries[j].firstTokenBegin
	})

	superSummary := []Summary{{
		firstTokenBegin: summaries[0].firstTokenBegin,
		sexp:            concatSexps(summaries, isSentence),
		verbTypes:       concatVerbTypes(summaries),
	}}

	return superSummary
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
