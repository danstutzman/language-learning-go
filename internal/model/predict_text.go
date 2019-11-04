package model

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ProbWord struct {
	prob float64
	word string
}

func readProbWords(languageModelPath string) []ProbWord {
	file, err := os.Open(languageModelPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	probWords := []ProbWord{}
	scanner := bufio.NewScanner(file)
	found1Grams := false
	for scanner.Scan() {
		line := scanner.Text()
		if !found1Grams {
			if line == `\1-grams:` {
				found1Grams = true
			}
		} else {
			if line == "" {
				break
			}

			values := strings.Split(line, "\t")
			prob, err := strconv.ParseFloat(values[0], 64)
			if err != nil {
				panic(err)
			}
			word := values[1]
			probWords = append(probWords, ProbWord{prob: prob, word: word})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return probWords
}

func (model *Model) PredictText(wordSoFar string) []string {
	if model.probWords == nil {
		model.probWords = readProbWords(model.languageModelPath)
	}

	if wordSoFar == "" {
		return []string{}
	}
	wordSoFarLower := strings.ToLower(wordSoFar)

	probWords := []ProbWord{}
	for _, probWord := range model.probWords {
		if strings.HasPrefix(probWord.word, wordSoFarLower) {
			probWords = append(probWords, probWord)
		}
	}

	sort.SliceStable(probWords, func(i, j int) bool {
		return probWords[i].prob > probWords[j].prob
	})

	predictions := []string{}
	for i, probWord := range probWords {
		predictions = append(predictions, probWord.word)

		if i >= 6 {
			break
		}
	}
	return predictions
}
