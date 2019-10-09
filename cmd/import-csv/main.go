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

type Import struct {
	phrase   string
	analysis model.Analysis
	err      error
}

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
	if dbPath == "" {
		log.Fatalf("Specify DB_PATH env var")
	}

	freelingHostAndPort := os.Getenv("FREELING_HOST_AND_PORT")
	if freelingHostAndPort == "" {
		log.Fatalf("Specify FREELING_HOST_AND_PORT env var, for example: 1.2.3.4:5678")
	}

	if len(os.Args) != 1+2 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: path to morphemes.csv
		Argument 2: path to stories.yaml`)
	}
	morphemesCsvPath := os.Args[1]
	storiesYamlPath := os.Args[2]

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
	importStoriesYaml(storiesYamlPath, theModel, freelingHostAndPort)
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

func importImport(import_ *Import, theModel *model.Model) {
	phrase := ""
	for _, sentence := range import_.analysis.Sentences {
		for _, token := range sentence.Tokens {
			phrase = phrase + token.Form + " "
		}
	}

	cardPhrases := extractBraceSurroundedPhrases(phrase)

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
			import_.err = fmt.Errorf("Expected [%s] but got [%s]",
				expectedWordsJoined, actualWordsJoined)
		} else {
			theModel.InsertCard(model.Card{
				L1:        "",
				L2:        cardPhrase,
				Morphemes: morphemes,
			})
		}
	}
}

type Story struct {
	Url   string        `yaml:"url"`
	Lines []interface{} `yaml:"lines"`
}

func importStoriesYaml(path string, theModel *model.Model, freelingHostAndPort string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	imports := []Import{}
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
				imports = append(imports, Import{phrase: l2.(string)})
			}
		}
	}

	phrases := []string{}
	for _, import_ := range imports {
		phrases = append(phrases, import_.phrase)
	}

	analyses := model.AnalyzePhrasesWithFreeling(phrases, freelingHostAndPort)

	for i, _ := range imports {
		imports[i].analysis = analyses[i]
	}

	for i, _ := range imports {
		importImport(&imports[i], theModel)
	}

	for _, import_ := range imports {
		if import_.err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", import_.err)
		}
	}
}
