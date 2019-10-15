package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

const PARSE_DIR = "db/1_parses"

func main() {
	if len(os.Args) != 2+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: path to corpus (.yaml or .csv .txt file)
		Argument 2: hostname:port for freeling server`)
	}
	corpusPath := os.Args[1]
	freelingHostAndPort := os.Args[2]

	googleTranslateApiKey := os.Getenv("GOOGLE_TRANSLATE_API_KEY")
	if googleTranslateApiKey == "" {
		log.Fatalf("Specify GOOGLE_TRANSLATE_API_KEY environment variable")
	}

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

	outputs := parsing.ParsePhrasesWithFreeling(phrases, freelingHostAndPort)

	for _, output := range outputs {
		parsing.SaveParse(output.Phrase, output.ParseJson, PARSE_DIR)
		fmt.Fprintf(os.Stderr, "%s\n", PARSE_DIR+"/"+output.Phrase+".json")
	}
}
