package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Import struct {
	phrase         string
	analysisJson   string
	analysis       model.Analysis
	sentenceErrors []error
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

	if len(os.Args) != 1+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: path to stories.yaml`)
	}
	storiesYamlPath := os.Args[1]

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

	importStoriesYaml(storiesYamlPath, theModel, freelingHostAndPort)
}

func removeCurlyBraces(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "{", ""), "}", "")
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func verbToMorphemes(token model.Token,
	theModel *model.Model) ([]model.Morpheme, error) {

	lemma := token.Lemma
	form := strings.ToLower(token.Form)
	tag := token.Tag

	unique := theModel.FindVerbUnique(form, lemma, tag)
	if unique != nil {
		return []model.Morpheme{*unique}, nil
	}

	var category string
	if strings.HasSuffix(lemma, "ar") {
		category = "ar"
	} else if strings.HasSuffix(lemma, "er") {
		category = "er"
	} else if strings.HasSuffix(lemma, "ir") {
		category = "ir"
	} else {
		log.Fatalf("Unknown category for lemma '%s'", lemma)
	}

	stemChangeMorpheme := theModel.FindVerbStemChange(lemma, token.Tense)
	if stemChangeMorpheme != nil {
		suffix := "-" + form[len(stemChangeMorpheme.L2)-1:len(form)]

		category = "stempret"
		suffixMorpheme := theModel.FindVerbSuffix(suffix, category, tag)
		if suffixMorpheme == nil {
			return []model.Morpheme{}, fmt.Errorf(
				"Can't find verb suffix '%s' with category=%s tag=%s",
				suffix, category, tag)
		}

		return []model.Morpheme{*stemChangeMorpheme, *suffixMorpheme}, nil
	} else { // If there is no stem change
		stem := lemma[0 : len(lemma)-len(category)]

		if !strings.HasPrefix(form, stem) {
			return []model.Morpheme{}, fmt.Errorf(
				"No stem change to explain why '%s' doesn't match lemma '%s'",
				form, lemma)
		}

		stemMorpheme := theModel.UpsertMorpheme(model.Morpheme{
			Type: "VERB_STEM",
			L2:   stem,
		})

		// Warning: for verbs like 'tengo' the suffix could be weird like 'go'.
		// This should be caught by the unique verb look up earlier, but otherwise
		// it will just fail on the suffix look up.
		suffix := "-" + form[len(stem):len(form)]

		suffixMorpheme := theModel.FindVerbSuffix(suffix, category, tag)
		if suffixMorpheme == nil {
			return []model.Morpheme{}, fmt.Errorf(
				"Can't find verb suffix '%s' with category=%s tag=%s",
				suffix, category, tag)
		}

		return []model.Morpheme{stemMorpheme, *suffixMorpheme}, nil
	}

}

func importImport(import_ *Import, theModel *model.Model) {
	import_.sentenceErrors = make([]error, len(import_.analysis.Sentences))
	for sentenceNum, sentence := range import_.analysis.Sentences {
		firstIndex := mustAtoi(sentence.Tokens[0].Begin)
		lastIndex := mustAtoi(sentence.Tokens[len(sentence.Tokens)-1].End)
		// Offset by number of Unicode points (runes), not number of bytes
		excerpt := string(([]rune(import_.phrase))[firstIndex:lastIndex])
		_ = excerpt

		expectedWords := []string{}
		for _, token := range sentence.Tokens {
			if !token.IsPunctuation() {
				expectedWords = append(expectedWords, strings.ToLower(token.Form))

				if strings.HasPrefix(token.Tag, "V") {
					morphemes, err := verbToMorphemes(token, theModel)
					if err != nil {
						import_.sentenceErrors[sentenceNum] = err
					}

					log.Printf("morphemes:%v %v", morphemes, err)
				}
			}
		}

		/*
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
				import_.sentenceErrors[sentenceNum] = fmt.Errorf(
					"Expected [%s] but got [%s]",
					expectedWordsJoined, actualWordsJoined)
			} else {
					theModel.InsertCard(model.Card{
						L1:        "",
						L2:        excerpt,
						Morphemes: morphemes,
					})
			}
		*/
	}
}

type Story struct {
	Url   string        `yaml:"url"`
	Lines []interface{} `yaml:"lines"`
}

func importStoriesYaml(path string, theModel *model.Model,
	freelingHostAndPort string) {

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

	outputs := model.AnalyzePhrasesWithFreeling(phrases, freelingHostAndPort)

	for i, _ := range imports {
		imports[i].analysis = outputs[i].Analysis
		imports[i].analysisJson = outputs[i].AnalysisJson
	}

	for i, _ := range imports {
		importImport(&imports[i], theModel)
	}

	//for _, import_ := range imports {
	//		for _, sentenceErr := range import_.sentenceErrors {
	//			if sentenceErr != nil {
	//				fmt.Fprintf(os.Stderr, "%s\n", sentenceErr)
	//fmt.Fprintf(os.Stderr, "%s\n", import_.analysisJson)
	//				os.Exit(1)
	//			}
	//		}
	//}
}
