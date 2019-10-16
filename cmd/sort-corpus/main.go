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
		fmt.Printf("%-20s %s\n", strings.Join(summary.sexp, ""), phrase)
	}
}

// Summary of a constituent
type Summary struct {
	firstTokenBegin int
	sexp            []string
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

	concattedSexps := []string{}
	for _, summary := range summaries {
		concattedSexps = append(concattedSexps, summary.sexp...)
	}

	superSummary := Summary{
		firstTokenBegin: summaries[0].firstTokenBegin,
		sexp:            concattedSexps,
	}
	return superSummary
}

func summarizeConstituent(constituent parsing.Constituent,
	tokenById map[string]parsing.Token, isSentence bool) []Summary {

	if len(constituent.Children) == 0 {
		token := tokenById[constituent.Token]
		shortTag := token.Tag[0:1]
		if shortTag == "F" { // discard punctuation
			return []Summary{}
		}
		return []Summary{{
			firstTokenBegin: mustAtoi(token.Begin),
			sexp:            []string{shortTag},
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

	superSummary := []Summary{{
		firstTokenBegin: summaries[0].firstTokenBegin,
		sexp:            concattedSexps,
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
