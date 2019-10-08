package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
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

	if len(os.Args) != 4 {
		log.Fatalf(`Usage:
		Argument 1: path to morphemes.csv
		Argument 2: path to cards.csv
		Argument 3: path to stories.yaml`)
	}
	morphemesCsvPath := os.Args[1]
	cardsCsvPath := os.Args[2]
	storiesYamlPath := os.Args[3]

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

	importMorphemesCsv(morphemesCsvPath, theModel)
	importCardsCsv(cardsCsvPath, theModel)
	importStoriesYaml(storiesYamlPath, theModel)
}

func importMorphemesCsv(path string, theModel *model.Model) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(bufio.NewReader(file))

	columnNames, err := reader.Read()
	if err != nil {
		panic(err)
	}
	l2Index := indexOf("l2", columnNames)
	glossIndex := indexOf("gloss", columnNames)

	for {
		values, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		l2 := values[l2Index]
		gloss := values[glossIndex]

		theModel.InsertMorpheme(model.Morpheme{L2: l2, Gloss: gloss})
	}
}

func importCardsCsv(path string, theModel *model.Model) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(bufio.NewReader(file))

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

		insertCardForPhrase(l1, l2, theModel)
	}
}

func removeCurlyBraces(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "{", ""), "}", "")
}

// Convert "{a {b}} c" to ["a b c", "a b", "b"]
func extractBraceSurroundedPhrases(input string) []string {
	indexStack := []int{}
	out := []string{removeCurlyBraces(input)}
	for i, c := range input {
		if c == '{' {
			indexStack = append(indexStack, i)
		} else if c == '}' {
			beginIndex := indexStack[len(indexStack)-1] + 1
			indexStack = indexStack[0 : len(indexStack)-1]
			extracted := removeCurlyBraces(input[beginIndex:i])
			out = append(out, extracted)
		}
	}
	return out
}

func insertCardForPhrase(l1, l2 string, theModel *model.Model) {
	cardPhrases := extractBraceSurroundedPhrases(l2)

	for _, cardPhrase := range cardPhrases {
		expectedWords := theModel.SplitL2PhraseIntoWords(cardPhrase)

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
		actualWordsJoined = strings.ReplaceAll(actualWordsJoined, "- -", "")
		actualWordsJoined = strings.ReplaceAll(actualWordsJoined, " -", "")
		actualWordsJoined = strings.ReplaceAll(actualWordsJoined, "- ", "")
		if actualWordsJoined != expectedWordsJoined {
			log.Fatalf("Expected [%s] but got [%s]", expectedWordsJoined, actualWordsJoined)
		}

		theModel.InsertCard(model.Card{
			L1:        l1,
			L2:        cardPhrase,
			Morphemes: morphemes,
		})
	}
}

type Story struct {
	Url   string        `yaml:"url"`
	Lines []interface{} `yaml:"lines"`
}

func importStoriesYaml(path string, theModel *model.Model) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	decoder := yaml.NewDecoder(bufio.NewReader(file))
	for {
		var story Story
		err = decoder.Decode(&story)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		for _, line := range story.Lines {
			var l2BySpeaker = line.(map[string]interface{})
			for _, l2 := range l2BySpeaker {
				insertCardForPhrase("", l2.(string), theModel)
			}
		}
	}
}
