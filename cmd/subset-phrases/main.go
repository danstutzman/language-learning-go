package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	"log"
	"os"
	"strings"
)

const PARSE_DIR = "db/1_parses"

func main() {
	if len(os.Args) != 1+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
			Argument 1: .txt corpus to load`)
	}
	corpusPath := os.Args[1]

	phrases := parsing.ListPhrasesInCorpusTxt(corpusPath)

	for phraseNum, phrase := range phrases {
		if !strings.Contains(phrase, "Ese es su hermano") {
			//continue
		}

		parse := parsing.LoadSavedParse(phrase, PARSE_DIR).Parse
		for _, sentence := range parse.Sentences {
			combos := combosOfDeps(sentence.Dependencies)
			for _, combo := range combos {
				words := []string{}
				for _, token := range sentence.Tokens {
					if combo[token.Id] {
						words = append(words, token.Form)
					}
				}
				fmt.Printf("%v\n", words)
			}
		}

		if false && phraseNum > 20 {
			break
		}
	}
}

type Dep struct {
	Token    string
	Children []Dep
}

func mapWithJust(token string, value bool) []map[string]bool {
	combo := map[string]bool{}
	combo[token] = false
	return []map[string]bool{combo}
}

func combosOfDep(dep parsing.Dependency) []map[string]bool {
	if len(dep.Children) == 0 {
		if dep.Function == "f" { // punctuation
			return []map[string]bool{
				map[string]bool{dep.Token: false},
			}
		}

		if dep.Function == "spec" || dep.Function == "conj" {
			return []map[string]bool{
				map[string]bool{dep.Token: true},
			}
		}

		return []map[string]bool{
			map[string]bool{dep.Token: false},
			map[string]bool{dep.Token: true},
		}
	}

	combos := []map[string]bool{}

	combos = append(combos, map[string]bool{dep.Token: false})

	for _, combo := range combosOfDeps(dep.Children) {
		copy := map[string]bool{}
		for key, value := range combo {
			copy[key] = value
		}
		copy[dep.Token] = true

		hasAtLeastOneTrue := false
		for _, value := range combo {
			if value {
				hasAtLeastOneTrue = true
			}
		}

		if (dep.Function == "sp" || dep.Function == "cc" || dep.Function ==
			"atr") && !hasAtLeastOneTrue {
			// skip
		} else {
			combos = append(combos, copy)
		}
	}

	return combos
}

func combosOfDeps(deps []parsing.Dependency) []map[string]bool {
	if len(deps) == 1 {
		return combosOfDep(deps[0])
	}

	// cartesian join
	combos := []map[string]bool{}
	for _, combo1 := range combosOfDep(deps[0]) {
		for _, combo2 := range combosOfDeps(deps[1:len(deps)]) {
			merged := map[string]bool{}
			for key, value := range combo1 {
				merged[key] = value
			}
			for key, value := range combo2 {
				merged[key] = value
			}
			combos = append(combos, merged)
		}
	}
	return combos
}
