package parsing

import (
	"bufio"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Story struct {
	Url   string        `yaml:"url"`
	Lines []interface{} `yaml:"lines"`
}

func ImportStoriesYaml(path, parseDir string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	phrases := []string{}
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
				phrase := l2.(string)
				phrases = append(phrases, phrase)
			}
		}
	}
	return phrases
}
