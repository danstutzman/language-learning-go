package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/mem_model"
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"bitbucket.org/danstutzman/language-learning-go/internal/spacy"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"sort"
	"strings"
)

const PARSE_DIR = "db/1_parses"

const FLAT = true

func main() {
	if len(os.Args) != 3+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: path to corpus (.txt or .csv file)
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
	} else if strings.HasSuffix(corpusPath, ".csv") {
		phrases = parsing.ListPhrasesInCorpusCsv(corpusPath)
	} else {
		log.Fatalf("Unrecognized extension for path '%s'", corpusPath)
	}

	phraseL2s := []string{}
	for _, phrase := range phrases {
		phraseL2s = append(phraseL2s, spacy.Uncapitalize1stLetter(phrase.L2))
	}

	dictionary := english.LoadDictionary(dictionaryDbPath)

	for phraseNum, phrase := range phrases {
		if phrase.L2 != "Él vende fruta." {
			//continue
		}

		parse := spacy.LoadSavedParse(phrase.L2, PARSE_DIR, "es")

		if FLAT {
			importPhraseFlat(phrase, parse, memModel)
			continue
		}

		output2s := importPhrase(phrase.L2, parse, dictionary, memModel)
		for _, output2 := range output2s {
			if output2.Error != nil {
				fmt.Printf("%d ", phraseNum+1)

				e := output2.Error
				if cantTranslate, ok := e.(CantTranslate); ok {
					fmt.Printf("text=%s %s\n",
						cantTranslate.Token.Text, cantTranslate.Message)
				} else if cantConvertDep, ok := e.(*CantConvertDep); ok {
					fmt.Printf("%s\n", cantConvertDep.Message)
					fmt.Printf("        %s\n", phrase.L2)
					fmt.Printf("        %s/%s\n",
						cantConvertDep.Parent.Token.Text,
						cantConvertDep.Parent.Token.Pos)
					for _, child := range cantConvertDep.Parent.Children {
						if child.Token.Id == cantConvertDep.Child.Token.Id {
							fmt.Printf("        ! ")
						} else {
							fmt.Printf("          ")
						}
						fmt.Printf("%s: %s/%s\n", child.Function, child.Token.Text,
							child.Token.Pos)

						if child.Token.Id == cantConvertDep.Child.Token.Id {
							for _, grandchild := range child.Children {
								fmt.Printf("            %s: %s/%s\n",
									grandchild.Function, grandchild.Token.Text,
									grandchild.Token.Pos)
							}
						}
					}
				} else {
					fmt.Printf("%v\n", e)
				}
			} else { // If no error
				//fmt.Printf("%d ", phraseNum+1)
				//fmt.Printf("%s\n", output2.Phrase)
			}
		}
	}

	memModel.SaveMorphemesToDb(dbConn)
	memModel.SaveCardsToDb(dbConn)
}

type TokenError struct {
	Token spacy.Token
	Error interface{}
}

type Output2 struct {
	Phrase string
	Error  interface{}
}

func importPhraseFlat(phrase parsing.Phrase, tokens []spacy.Token,
	memModel *mem_model.MemModel) {

	cardByTokenId := map[int]mem_model.Card{}
	for _, token := range tokens {
		card, err := memModel.TokenToCard(token)
		if err != nil {
			panic(err)
		}
		cardByTokenId[token.Id] = card
	}

	cardMorphemes := []mem_model.CardMorpheme{}
	for _, token := range tokens {
		cardMorphemes = append(cardMorphemes, cardByTokenId[token.Id].Morphemes...)
	}

	memModel.InsertCardIfNotExists(mem_model.Card{
		Type:       "Sentence",
		L1:         phrase.L1,
		L2:         phrase.L2,
		IsSentence: true,
		Morphemes:  cardMorphemes,
	})
}

func importPhrase(phrase string, tokens []spacy.Token,
	dictionary english.Dictionary, memModel *mem_model.MemModel) []Output2 {

	tokenErrors := []TokenError{}

	cardByTokenId := map[int]mem_model.Card{}
	for _, token := range tokens {
		card, err := memModel.TokenToCard(token)
		if err != nil {
			tokenErrors = append(tokenErrors, TokenError{Token: token, Error: err})
			continue
		}
		cardByTokenId[token.Id] = card
	}

	output2s := []Output2{}
	for _, dep := range spacy.TokensToDeps(tokens) {
		s, err := depToS(dep)
		if err != nil {
			if false {
				printTokensInOrder(os.Stderr, tokens)
				fmt.Fprintf(os.Stderr, "\\ %v\n", err)
			}

			output2s = append(output2s, Output2{
				Phrase: phrase,
				Error:  err,
			})
			continue
		}

		err2 := importConstituent(s, cardByTokenId, true, dictionary, memModel)
		if err2 != nil {
			output2s = append(output2s, Output2{
				Phrase: phrase,
				Error:  err2,
			})
			continue
		}

		output2s = append(output2s, Output2{
			Phrase: phrase,
			Error:  nil,
		})
	}
	return output2s
}

func importConstituent(constituent Constituent,
	cardByTokenId map[int]mem_model.Card,
	isSentence bool, dictionary english.Dictionary,
	memModel *mem_model.MemModel) *CantTranslate {

	tokens := constituent.GetAllTokens()
	sort.SliceStable(tokens, func(i, j int) bool {
		return tokens[i].Idx < tokens[j].Idx
	})

	hasBlankToken := false
	for _, token := range tokens {
		if token.Text == "" {
			hasBlankToken = true
		}
	}
	if hasBlankToken {
		log.Fatalf("Blank token from %+v", constituent)
	}

	l2 := ""
	for i, token := range tokens {
		if i > 0 && token.Idx > (tokens[i-1].Idx+len([]rune(tokens[i-1].Text))) {
			l2 += " "
		}
		l2 += token.Text
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

func minBeginForDep(dep spacy.Dep) int {
	minBegin := dep.Token.Idx

	for _, child := range dep.Children {
		childMinBegin := minBeginForDep(child)
		if childMinBegin < minBegin {
			minBegin = childMinBegin
		}
	}

	return minBegin
}
