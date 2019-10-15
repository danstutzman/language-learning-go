package parsing

import (
	"bufio"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func ImportStoriesYaml(path, parseDir string) []string {
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
