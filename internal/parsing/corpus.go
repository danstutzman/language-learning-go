package parsing

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"strings"
)

func ListPhrasesInCorpusYaml(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	phrases := []string{}
	decoder := yaml.NewDecoder(bufio.NewReader(file))
	for {
		var lines []interface{}
		err = decoder.Decode(&lines)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		for _, line := range lines {
			var l2BySpeaker = line.(map[string]interface{})
			for _, l2 := range l2BySpeaker {
				phrase := l2.(string)
				phrases = append(phrases, phrase)
			}
		}
	}
	return phrases
}

func ListPhrasesInCorpusCsv(path string) []string {
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
	l2Index := indexOf("l2", columnNames)

	phrases := []string{}
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

		l2 := values[l2Index]
		phrases = append(phrases, l2)
	}
	return phrases
}

func indexOf(needle string, haystack []string) int {
	for index, element := range haystack {
		if element == needle {
			return index
		}
	}
	panic(fmt.Sprintf("Needle '%s' not found in %v", needle, haystack))
}

func ListPhrasesInCorpusTxt(path string) []string {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	phrases := []string{}
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		phrase := strings.TrimSpace(line)
		if phrase != "" {
			phrases = append(phrases, phrase)
		}
	}
	return phrases
}
