package english

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var inParensRegex = regexp.MustCompile(`\([^)]*\)`)

var inBracketsRegex = regexp.MustCompile(`\[[^]]*\]`)

var commaAndAfterRegex = regexp.MustCompile(`,.*`)

type Dictionary struct {
	englishByEsAndPartOfSpeech map[string]string
}

func (dictionary Dictionary) Lookup(es, partOfSpeech string) string {
	return dictionary.englishByEsAndPartOfSpeech[es+"/"+partOfSpeech]
}

func LoadDictionary(path string) Dictionary {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	columnNames, err := reader.Read()
	if err != nil {
		panic(err)
	}
	partOfSpeechIndex := indexOf("part_of_speech", columnNames)
	esIndex := indexOf("es", columnNames)
	englishIndex := indexOf("english", columnNames)

	englishByEsAndPartOfSpeech := map[string]string{}
	for {
		values, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		partOfSpeech := values[partOfSpeechIndex]
		es := values[esIndex]
		english := values[englishIndex]

		english = inParensRegex.ReplaceAllString(english, "")
		english = inBracketsRegex.ReplaceAllString(english, "")
		english = commaAndAfterRegex.ReplaceAllString(english, "")
		english = strings.TrimSpace(english)

		if partOfSpeech == "v" && strings.HasPrefix(english, "to ") {
			english = english[3:len(english)]
		}

		englishByEsAndPartOfSpeech[es+"/"+partOfSpeech] = english
	}

	englishByEsAndPartOfSpeech["llevar"+"/"+"v"] = "wear"
	englishByEsAndPartOfSpeech["gustar"+"/"+"v"] = "please"

	return Dictionary{englishByEsAndPartOfSpeech}
}

func indexOf(needle string, haystack []string) int {
	for index, element := range haystack {
		if element == needle {
			return index
		}
	}
	panic(fmt.Sprintf("Needle '%s' not found in %v", needle, haystack))
}
