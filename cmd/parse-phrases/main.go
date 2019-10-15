package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) != 2+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: path to stories.yaml
		Argument 2: hostname:port for freeling server`)
	}
	storiesYamlPath := os.Args[1]
	freelingHostAndPort := os.Args[2]

	googleTranslateApiKey := os.Getenv("GOOGLE_TRANSLATE_API_KEY")
	if googleTranslateApiKey == "" {
		log.Fatalf("Specify GOOGLE_TRANSLATE_API_KEY environment variable")
	}

	parseDir := "db/1_parses"

	phrases := parsing.ImportStoriesYaml(storiesYamlPath, parseDir)

	outputs := parsing.ParsePhrasesWithFreeling(phrases, freelingHostAndPort)

	for _, output := range outputs {
		parsing.SaveParse(output.Phrase, output.ParseJson, parseDir)
		fmt.Fprintf(os.Stderr, "%s\n", parseDir+"/"+output.Phrase+".json")
	}

	writeSubphraseTranslations(outputs, googleTranslateApiKey)
}

func writeSubphraseTranslations(outputs []parsing.Output,
	googleTranslateApiKey string) {

	subphrasesSet := map[string]bool{}
	for _, output := range outputs {
		for _, sentence := range output.Parse.Sentences {
			tokenById := map[string]parsing.Token{}
			for _, token := range sentence.Tokens {
				tokenById[token.Id] = token
			}

			for _, constituent := range sentence.Constituents {
				collectSubphrasesFromConstituent(constituent, tokenById, subphrasesSet)
			}
		}
	}

	subphrases := []string{}
	for subphrase, _ := range subphrasesSet {
		subphrases = append(subphrases, subphrase)
	}

	translations := parsing.TranslateToEnglish(subphrases, googleTranslateApiKey)

	l1ByL2 := map[string]string{}
	for i, subphrase := range subphrases {
		l1ByL2[subphrase] = translations[i]
	}

	translationsYaml, err := yaml.Marshal(&l1ByL2)
	if err != nil {
		panic(err)
	}

	translationsPath := "db/1_translations.yaml"
	err = ioutil.WriteFile(translationsPath, translationsYaml, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(translationsPath)
}

func collectSubphrasesFromConstituent(constituent parsing.Constituent,
	tokenById map[string]parsing.Token, subphrasesSet map[string]bool) {

	tokens := getTokensForConstituent(constituent, tokenById)

	l2 := ""
	for i, token := range tokens {
		if i > 0 && mustAtoi(token.Begin) > mustAtoi(tokens[i-1].End) {
			l2 += " "
		}
		l2 += token.Form
	}
	subphrasesSet[l2] = true

	for _, child := range constituent.Children {
		collectSubphrasesFromConstituent(child, tokenById, subphrasesSet)
	}
}

func getTokensForConstituent(constituent parsing.Constituent,
	tokenById map[string]parsing.Token) []parsing.Token {

	tokens := []parsing.Token{}
	if constituent.Token != "" {
		tokens = append(tokens, tokenById[constituent.Token])
	}
	for _, child := range constituent.Children {
		tokens = append(tokens, getTokensForConstituent(child, tokenById)...)
	}
	sort.SliceStable(tokens, func(i, j int) bool {
		return mustAtoi(tokens[i].Begin) < mustAtoi(tokens[j].Begin)
	})
	return tokens
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
