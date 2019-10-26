package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	"strings"
)

func oldSubsettingDemo(phrases []string) {
	for phraseNum, phrase := range phrases {
		if !strings.Contains(phrase, "prueba la manzana") {
			//continue
		}

		parse := parsing.LoadSavedParse(phrase, PARSE_DIR).Parse
		for _, sentence := range parse.Sentences {
			tokenById := map[string]parsing.Token{}
			for _, token := range sentence.Tokens {
				tokenById[token.Id] = token
			}

			for _, dep := range listAllDepsOfDeps(sentence.Dependencies) {
				combos := combosOfDep(dep, tokenById)

				for _, combo := range combos {
					words := []string{}
					for _, token := range sentence.Tokens {
						if combo[token.Id] {
							words = append(words, token.Form)
						}
					}

					if len(words) > 0 {
						fmt.Printf("%v\n", words)
					}
				}
			}
		}

		if false && phraseNum > 20 {
			break
		}
	}
}

func printDeps(deps []parsing.Dependency, tokenById map[string]parsing.Token,
	indentation int) {
	for _, dep := range deps {

		for i := 0; i < indentation; i += 1 {
			fmt.Printf("  ")
		}
		fmt.Printf("%s: %s (%s)\n",
			dep.Function, dep.Word, tokenById[dep.Token].Tag)

		printDeps(dep.Children, tokenById, indentation+1)
	}
}

func listAllDepsOfDeps(deps []parsing.Dependency) []parsing.Dependency {
	allDeps := []parsing.Dependency{}
	for _, dep := range deps {
		allDeps = append(allDeps, listAllDepsOfDep(dep)...)
	}
	return allDeps
}

func listAllDepsOfDep(dep parsing.Dependency) []parsing.Dependency {
	deps := []parsing.Dependency{dep}
	for _, child := range dep.Children {
		deps = append(deps, listAllDepsOfDep(child)...)
	}
	return deps
}

func mapWithJust(token string, value bool) []map[string]bool {
	combo := map[string]bool{}
	combo[token] = false
	return []map[string]bool{combo}
}

func isSerOrEstar(lemma string) bool {
	return lemma == "ser" || lemma == "estar"
}

func combosOfDep(dep parsing.Dependency,
	tokenById map[string]parsing.Token) []map[string]bool {

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

	for _, combo := range combosOfDeps(dep.Children, tokenById) {
		hasAtLeastOneTrue := false
		for _, value := range combo {
			if value {
				hasAtLeastOneTrue = true
			}
		}

		hasAtr := false
		for _, child := range dep.Children {
			if child.Function == "atr" && combo[child.Token] {
				hasAtr = true
			}
		}
		isLinking := isSerOrEstar(tokenById[dep.Token].Lemma)

		if (dep.Function == "sp" || dep.Function == "cc" || dep.Function ==
			"atr") && !hasAtLeastOneTrue {
			// skip
		} else if isLinking && !hasAtr {
			// skip
		} else {
			copy := map[string]bool{}
			for key, value := range combo {
				copy[key] = value
			}
			copy[dep.Token] = true

			combos = append(combos, copy)
		}
	}

	return combos
}

func combosOfDeps(deps []parsing.Dependency,
	tokenById map[string]parsing.Token) []map[string]bool {

	if len(deps) == 1 {
		return combosOfDep(deps[0], tokenById)
	}

	// cartesian join
	combos := []map[string]bool{}
	for _, combo1 := range combosOfDep(deps[0], tokenById) {
		for _, combo2 := range combosOfDeps(deps[1:len(deps)], tokenById) {
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
