package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/mem_model"
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"sort"
	"strings"
)

const PARSE_DIR = "db/1_parses"

func main() {
	if len(os.Args) != 3+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: path to corpus (.txt file)
		Argument 2: path to sqlite3 database file
		Argument 3: path to dictionary sqlite3 database file`)
	}
	corpusPath := os.Args[1]
	dbPath := os.Args[2]
	dictionaryDbPath := os.Args[3]

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
	memModel := mem_model.NewMemModel()

	var phrases []parsing.Phrase
	if strings.HasSuffix(corpusPath, ".txt") {
		phrases = parsing.ListPhrasesInCorpusTxt(corpusPath)
	} else {
		log.Fatalf("Unrecognized extension for path '%s'", corpusPath)
	}

	dictionary := english.LoadDictionary(dictionaryDbPath)

	for _, phrase := range phrases {
		if phrase.L2 != "Ella manejaba un auto verde cuando la vi." {
			//continue
		}

		output := parsing.LoadSavedParse(phrase, PARSE_DIR)

		output2s := importPhrase(output, dictionary, memModel)
		for _, output2 := range output2s {
			if output2.Error != nil {
				fmt.Fprintf(os.Stderr, "%d:%d %s\n",
					output2.Phrase.LineNum, output2.Phrase.CharNum, output2.Error)
			}
		}
	}

	memModel.SaveMorphemesToDb(dbConn)
	memModel.SaveCardsToDb(dbConn)
}

func lowercaseToken(token parsing.Token) parsing.Token {
	token.Form = strings.ToLower(token.Form)
	return token
}

type Output2 struct {
	Phrase parsing.Phrase
	Error  error
}

func importPhrase(output parsing.Output, dictionary english.Dictionary,
	memModel *mem_model.MemModel) []Output2 {

	output2s := []Output2{}
	for _, sentence := range output.Parse.Sentences {
		tokenById := map[string]parsing.Token{}
		allTokens := []parsing.Token{}
		for _, token := range sentence.Tokens {
			tokenById[token.Id] = token
			allTokens = append(allTokens, token)
		}

		// Uncapitalize first token
		for i, token := range sentence.Tokens {
			if !token.IsPunctuation() {
				if !token.IsProperNoun() {
					sentence.Tokens[i] = lowercaseToken(token)
				}
				break
			}
		}

		cardByTokenId := map[string]mem_model.Card{}
		for _, token := range sentence.Tokens {
			card, err := memModel.TokenToCard(token)
			if err != nil {
				output2s = append(output2s, Output2{
					Phrase: parsing.Phrase{
						L2:      token.Form,
						LineNum: output.Phrase.LineNum,
						CharNum: output.Phrase.CharNum + mustAtoi(token.Begin),
					},
					Error: err,
				})
				continue
			}
			cardByTokenId[token.Id] = card
		}

		for _, dep := range sentence.Dependencies {
			s, err := depToS(dep, tokenById)
			if err != nil {
				if false {
					printTokensInOrder(os.Stderr, allTokens)
					fmt.Fprintf(os.Stderr, "\\ %s\n", err)
				}

				output2s = append(output2s, Output2{
					Phrase: parsing.Phrase{
						L2:      "TODO",
						LineNum: output.Phrase.LineNum,
						CharNum: output.Phrase.CharNum + minBeginForDep(dep, tokenById),
					},
					Error: err,
				})
				continue
			}

			err = importConstituent(s, cardByTokenId, true, dictionary, memModel)
			output2s = append(output2s, Output2{
				Phrase: parsing.Phrase{
					L2:      "TODO",
					LineNum: output.Phrase.LineNum,
					CharNum: output.Phrase.CharNum + minBeginForDep(dep, tokenById),
				},
				Error: err,
			})
		}
	}
	return output2s
}

func importConstituent(constituent Constituent,
	cardByTokenId map[string]mem_model.Card,
	isSentence bool, dictionary english.Dictionary,
	memModel *mem_model.MemModel) error {

	tokens := constituent.GetAllTokens()
	sort.SliceStable(tokens, func(i, j int) bool {
		return mustAtoi(tokens[i].Begin) < mustAtoi(tokens[j].Begin)
	})

	hasBlankToken := false
	for _, token := range tokens {
		if token.Form == "" {
			hasBlankToken = true
		}
	}
	if hasBlankToken {
		log.Fatalf("Blank token from %+v", constituent)
	}

	l2 := ""
	for i, token := range tokens {
		if i > 0 && mustAtoi(token.Begin) > mustAtoi(tokens[i-1].End) {
			l2 += " "
		}
		l2 += token.Form
	}

	cardMorphemes := []mem_model.CardMorpheme{}
	for _, token := range tokens {
		cardMorphemes = append(cardMorphemes, cardByTokenId[token.Id].Morphemes...)
	}

	for _, child := range constituent.GetChildren() {
		err := importConstituent(child, cardByTokenId, false, dictionary, memModel)
		if err != nil {
			return err
		}
	}

	l1, err := constituent.Translate(dictionary)
	if err != nil {
		return err
	}

	memModel.InsertCardIfNotExists(mem_model.Card{
		Type:       constituent.GetType(),
		L1:         strings.Join(l1, " "),
		L2:         l2,
		IsSentence: isSentence,
		Morphemes:  cardMorphemes,
	})
	return nil
}

func minBeginForDep(dep parsing.Dependency,
	tokenById map[string]parsing.Token) int {

	minBegin := mustAtoi(tokenById[dep.Token].Begin)

	for _, child := range dep.Children {
		childMinBegin := minBeginForDep(child, tokenById)
		if childMinBegin < minBegin {
			minBegin = childMinBegin
		}
	}

	return minBegin
}
