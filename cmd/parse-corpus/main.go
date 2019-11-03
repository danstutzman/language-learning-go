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
		Argument 1: path to corpus (.txt file)`)
	}
	corpusPath := os.Args[1]

	python3Path := os.Getenv("PYTHON3_PATH")
	if python3Path == "" {
		log.Fatalf("Specify PYTHON3_PATH environment variable")
	}

	var phrases []parsing.Phrase
	if strings.HasSuffix(corpusPath, ".txt") {
		phrases = parsing.ListPhrasesInCorpusTxt(corpusPath)
	} else {
		log.Fatalf("Unrecognized extension for path '%s'", corpusPath)
	}

	var phraseL2s []string
	for _, phrase := range phrases {
		phraseL2s = append(phraseL2s, spacy.Uncapitalize1stLetter(phrase.L2))
	}

	spacy.ParsePhrasesWithSpacyCached(phraseL2s, python3Path, PARSE_DIR)
}
