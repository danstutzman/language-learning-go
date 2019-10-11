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
	"sort"
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
		tokenById := map[string]parsing.Token{}
		for _, token := range sentence.Tokens {
			tokenById[token.Id] = token
		}

		// Uncapitalize first token
		for i, token := range sentence.Tokens {
			if !token.IsPunctuation() {
				if !token.IsProperNoun() {
					sentence.Tokens[i] = theModel.LowercaseToken(token)
				}
				break
			}
		}

		cardByTokenId := map[string]model.Card{}
		for _, token := range sentence.Tokens {
			unsavedCard, err := theModel.TokenToCard(token)
			if err != nil {
				errors[sentenceNum] = err
			}

			card := theModel.InsertCardIfNotExists(*unsavedCard)
			cardByTokenId[token.Id] = card
		}

		for _, constituent := range sentence.Constituents {
			importConstituent(constituent, cardByTokenId, tokenById, theModel)
		}

	}
	return errors
}

func importConstituent(constituent parsing.Constituent,
	cardByTokenId map[string]model.Card, tokenById map[string]parsing.Token,
	theModel *model.Model) model.Card {

	tokens := getTokensForConstituent(constituent, tokenById)

	l2 := ""
	for i, token := range tokens {
		if i > 0 && mustAtoi(token.Begin) > mustAtoi(tokens[i-1].End) {
			l2 += " "
		}
		l2 += token.Form
	}

	morphemes := []model.Morpheme{}
	for _, token := range tokens {
		morphemes = append(morphemes, cardByTokenId[token.Id].Morphemes...)
	}

	for _, child := range constituent.Children {
		importConstituent(child, cardByTokenId, tokenById, theModel)
	}

	if constituent.Leaf == "1" {
		return cardByTokenId[constituent.Token]
	} else {
		card := theModel.InsertCardIfNotExists(model.Card{
			Type:      constituent.Label,
			L2:        l2,
			Morphemes: morphemes,
		})
		return card
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
