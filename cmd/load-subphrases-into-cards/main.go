package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: path to stories.yaml
		Argument 2: path to sqlite3 database file`)
	}
	storiesYamlPath := os.Args[1]
	dbPath := os.Args[2]

	parseDir := "db/1_parses"

	// Set mode=rw so it doesn't create database if file doesn't exist
	connString := fmt.Sprintf("file:%s?mode=rw", dbPath)
	dbConn, err := sql.Open("sqlite3", connString)
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	db.AssertCardsHasCorrectSchema(dbConn)
	db.AssertCardsMorphemesHasCorrectSchema(dbConn)
	db.AssertMorphemesHasCorrectSchema(dbConn)
	theModel := model.NewModel(dbConn)

	phrases := parsing.ImportStoriesYaml(storiesYamlPath, parseDir)

	for _, phrase := range phrases {
		output := parsing.LoadSavedParse(phrase, parseDir)

		errors := importPhrase(output.Phrase, output.Parse, theModel)
		for _, err := range errors {
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		}
	}
}

func importPhrase(phrase string, parse parsing.Parse,
	theModel *model.Model) []error {

	errors := make([]error, len(parse.Sentences))
	for sentenceNum, sentence := range parse.Sentences {
		// Uncapitalize first token
		for i, token := range sentence.Tokens {
			if !token.IsPunctuation() {
				if !token.IsProperNoun() {
					sentence.Tokens[i] = theModel.LowercaseToken(token)
				}
				break
			}
		}

		// Insert a card for each word.  Some words are two morphemes.
		allMorphemes := []model.Morpheme{}
		for _, token := range sentence.Tokens {
			morphemes, err := theModel.TokenToMorphemes(token)
			if err != nil {
				errors[sentenceNum] = err
			}
			allMorphemes = append(allMorphemes, morphemes...)

			theModel.InsertCardIfNotExists(model.Card{
				L2:        token.Form,
				Morphemes: morphemes,
			})
		}

		// Insert the entire sentence
		firstTokenBegin := mustAtoi(sentence.Tokens[0].Begin)
		lastTokenEnd := mustAtoi(sentence.Tokens[len(sentence.Tokens)-1].End)
		l2 := string(([]rune(phrase))[firstTokenBegin:lastTokenEnd])
		theModel.InsertCardIfNotExists(model.Card{
			L2:        l2,
			Morphemes: allMorphemes,
		})
	}
	return errors
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
