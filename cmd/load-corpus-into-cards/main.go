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
		Argument 1: path to corpus (.yaml or .csv or .txt file)
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

	dictionary := english.LoadDictionary(dictionaryDbPath)

	for _, phrase := range phrases {
		output := parsing.LoadSavedParse(phrase, PARSE_DIR)

		errorLists := importPhrase(output.Parse, dictionary, memModel)
		for _, errorList := range errorLists {
			for _, err := range errorList {
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				}
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

func importPhrase(parse parsing.Parse, dictionary english.Dictionary,
	memModel *mem_model.MemModel) [][]error {

	errors := make([][]error, len(parse.Sentences))
	for sentenceNum, sentence := range parse.Sentences {
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
		errors[sentenceNum] = []error{}
		for _, token := range sentence.Tokens {
			card, err := memModel.TokenToCard(token)
			if err == nil {
				cardByTokenId[token.Id] = card
			} else {
				errors[sentenceNum] = append(errors[sentenceNum], err)
			}
		}

		if len(errors[sentenceNum]) == 0 {
			for _, dep := range sentence.Dependencies {
				s, err := depToS(dep, tokenById)
				if err != nil {
					if false {
						printTokensInOrder(os.Stderr, allTokens)
						fmt.Fprintf(os.Stderr, "\\ %s\n", err)
					}
				} else {
					if false {
						printTokensInOrder(os.Stdout, s.GetAllTokens())
						if len(s.vp) == 1 {
							fmt.Printf("Conjugations: %v\n", s.vp[0].verbConjugation)
						}
					}

					importConstituent(s, cardByTokenId, true, dictionary, memModel)
				}
			}
		}
	}
	return errors
}

func importConstituent(constituent Constituent,
	cardByTokenId map[string]mem_model.Card,
	isSentence bool, dictionary english.Dictionary,
	memModel *mem_model.MemModel) {

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
		importConstituent(child, cardByTokenId, false, dictionary, memModel)
	}

	memModel.InsertCardIfNotExists(mem_model.Card{
		Type:       constituent.GetType(),
		L1:         strings.Join(constituent.Translate(dictionary), " "),
		L2:         l2,
		IsSentence: isSentence,
		Morphemes:  cardMorphemes,
	})
}
