package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"bitbucket.org/danstutzman/language-learning-go/internal/spacy"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

const PARSE_DIR = "db/1_parses"

func main() {
	if len(os.Args) != 1+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: path to corpus (.txt or .csv file)`)
	}
	corpusPath := os.Args[1]

	python3Path := os.Getenv("PYTHON3_PATH")
	if python3Path == "" {
		log.Fatalf("Specify PYTHON3_PATH environment variable")
	}

	var phrases []parsing.Phrase
	if strings.HasSuffix(corpusPath, ".txt") {
		phrases = parsing.ListPhrasesInCorpusTxt(corpusPath)
	} else if strings.HasSuffix(corpusPath, ".csv") {
		phrases = parsing.ListPhrasesInCorpusCsv(corpusPath)
	} else {
		log.Fatalf("Unrecognized extension for path '%s'", corpusPath)
	}

	var phraseL1s []string
	for _, phrase := range phrases {
		phraseL1s = append(phraseL1s, spacy.Uncapitalize1stLetter(phrase.L1))
	}
	if len(phraseL1s) > 0 {
		spacy.ParsePhrasesWithSpacyCached(phraseL1s, python3Path, PARSE_DIR, "en")
	}

	var phraseL2s []string
	for _, phrase := range phrases {
		phraseL2s = append(phraseL2s, spacy.Uncapitalize1stLetter(phrase.L2))
	}
	if len(phraseL2s) > 0 {
		spacy.ParsePhrasesWithSpacyCached(phraseL2s, python3Path, PARSE_DIR, "es")
	}
}
