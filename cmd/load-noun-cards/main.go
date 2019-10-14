package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/guregu/null.v3"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: path to nouns.csv
		Argument 2: path to sqlite3 database file`)
	}
	csvPath := os.Args[1]
	dbPath := os.Args[2]

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

	csvFile, err := os.Open(csvPath)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	columnNames, err := reader.Read()
	if err != nil {
		panic(err)
	}
	l1Index := indexOf("l1", columnNames)
	l2Index := indexOf("l2", columnNames)
	mnemonic12Index := indexOf("mnemonic12", columnNames)
	mnemonic21Index := indexOf("mnemonic21", columnNames)
	nounGenderIndex := indexOf("noun_gender", columnNames)

	for {
		values, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if strings.HasPrefix(values[0], "#") {
			continue
		}

		l1 := values[l1Index]
		l2 := values[l2Index]
		mnemonic12 := values[mnemonic12Index]
		mnemonic21 := values[mnemonic21Index]
		nounGender := values[nounGenderIndex]

		card := theModel.InsertCard(model.Card{
			L1:         l1,
			L2:         l2,
			Mnemonic12: null.StringFrom(mnemonic12),
			Mnemonic21: null.StringFrom(mnemonic21),
			NounGender: null.StringFrom(nounGender),
			Type:       "NOUN",
		})

		theModel.InsertChallenge(model.Challenge{
			Type:   "Given1Type2",
			CardId: card.Id,

			Expectation: "NO_MNEMONIC",
		})

		theModel.InsertChallenge(model.Challenge{
			Type:   "Given2Type1",
			CardId: card.Id,

			Expectation: "NO_MNEMONIC",
		})
	}
}

func indexOf(needle string, haystack []string) int {
	for index, element := range haystack {
		if element == needle {
			return index
		}
	}
	panic(fmt.Sprintf("Needle '%s' not found in %v", needle, haystack))
}
