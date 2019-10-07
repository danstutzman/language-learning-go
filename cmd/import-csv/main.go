package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"os"
	"strings"
)

func indexOf(needle string, haystack []string) int {
	for index, element := range haystack {
		if element == needle {
			return index
		}
	}
	panic(fmt.Sprintf("Needle '%s' not found in %v", needle, haystack))
}

func main() {
	dbPath := os.Getenv("DB_PATH")

	if len(os.Args) != 2 {
		log.Fatalf("As first argument, specify path to CSV to import")
	}
	csvPath := os.Args[1]

	// Set mode=rw so it doesn't create database if file doesn't exist
	connString := fmt.Sprintf("file:%s?mode=rw", dbPath)
	dbConn, err := sql.Open("sqlite3", connString)
	if err != nil {
		panic(err)
	}
	db.AssertCardsHasCorrectSchema(dbConn)
	db.AssertCardsMorphemesHasCorrectSchema(dbConn)
	db.AssertMorphemesHasCorrectSchema(dbConn)
	theModel := model.NewModel(dbConn)

	csvFile, err := os.Open(csvPath)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	columnNames, err := reader.Read()
	if err != nil {
		panic(err)
	}
	l1Index := indexOf("l1", columnNames)
	l2Index := indexOf("l2", columnNames)

	for {
		values, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		l1 := values[l1Index]
		l2 := values[l2Index]

		expectedWords := theModel.SplitL2PhraseIntoWords(l2)

		morphemes := []model.Morpheme{}
		for _, word := range expectedWords {
			morphemes = append(morphemes, theModel.ParseL2WordIntoMorphemes(word)...)
		}

		actualWords := []string{}
		for _, morpheme := range morphemes {
			actualWords = append(actualWords, morpheme.L2)
		}

		expectedWordsJoined := strings.Join(expectedWords, " ")
		actualWordsJoined := strings.Join(actualWords, " ")
		if actualWordsJoined != expectedWordsJoined {
			log.Fatalf("Expected [%s] but got [%s]", expectedWordsJoined, actualWordsJoined)
		}

		theModel.InsertCard(model.Card{L1: l1, L2: l2, Morphemes: morphemes})
	}
}
