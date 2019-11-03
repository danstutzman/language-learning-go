package spacy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const BATCH_SIZE = 100

func unmarshalParseTxt(jsonLine []byte) []Token {
	var tokens []Token
	err := json.Unmarshal(jsonLine, &tokens)
	if err != nil {
		panic(err)
	}

	for i, token := range tokens {
		if token.Pos == "VERB" || token.Pos == "AUX" {
			freelingTag := FreelingTagBySpacyTag[token.SpacyTag]
			if freelingTag == "" {
				fmt.Fprintf(os.Stderr,
					"Can't find freelingTag for %v\n", token.SpacyTag)
			}
			tokens[i].VerbTag = freelingTag
		}

		if !strings.HasSuffix(token.SpacyTag, "___") {
			part2 := strings.Split(token.SpacyTag, "__")[1]

			features := map[string]string{}
			for _, pair := range strings.Split(part2, "|") {
				parts := strings.Split(pair, "=")
				features[parts[0]] = parts[1]
			}
			tokens[i].Features = features
		}
	}

	return tokens
}

func SaveParse(phrase string, parseTxt, parseDir string) {
	path := parseDir + "/" + phrase
	err := ioutil.WriteFile(path, []byte(parseTxt), 0644)
	if err != nil {
		panic(err)
	}

	tokens := unmarshalParseTxt([]byte(parseTxt))
	deps := TokensToDeps(tokens)
	lines := stringifyDeps(deps, 0)
	yamlPath := parseDir + "/" + phrase + ".yaml"
	err = ioutil.WriteFile(yamlPath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		panic(err)
	}
}

func stringifyDeps(deps []Dep, indentation int) []string {
	lines := []string{}
	for _, dep := range deps {
		line := ""
		for i := 0; i < indentation; i++ {
			line += "  "
		}
		line += dep.Function + ": " + dep.Token.Text + "/" + dep.Token.SpacyTag
		lines = append(lines, line)

		lines = append(lines, stringifyDeps(dep.Children, indentation+1)...)
	}
	return lines
}

func LoadSavedParse(phrase string, parseDir string) []Token {
	path := parseDir + "/" + phrase
	parseTxt, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return unmarshalParseTxt(parseTxt)
}

func ParsePhrasesWithSpacyCached(phrases []string,
	python3Path, parseDir string) {
	newPhrases := []string{}
	for i, phrase := range phrases {
		if !fileExists(parseDir + "/" + phrase) {
			newPhrases = append(newPhrases, phrase)
		}

		if len(newPhrases) >= BATCH_SIZE || i == len(phrases)-1 {
			parseTxts := ParseWithSpacy(newPhrases, python3Path)

			for j, parseTxt := range parseTxts {
				SaveParse(newPhrases[j], parseTxt, parseDir)
				fmt.Fprintf(os.Stderr, "%s\n", parseDir+"/"+newPhrases[j])
			}

			newPhrases = []string{}
		}
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
