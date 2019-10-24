package main

import (
	commandsPkg "bitbucket.org/danstutzman/language-learning-go/internal/commands"
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
		parse := parsing.LoadSavedParse(phrase, PARSE_DIR).Parse
		for _, sentence := range parse.Sentences {
			tokenById := map[string]parsing.Token{}
			for _, token := range sentence.Tokens {
				tokenById[token.Id] = token
			}

			for _, dependency := range sentence.Dependencies {
				commands := []string{}
				verbToken := tokenById[dependency.Token]
				if verbToken.Tag == "VMIP3S0" { // indic present 3rd person singular
					parallelVerb := parallelVerbByL2[verbToken.Lemma]
					commands = append(commands,
						"ADD/VERB/"+verbToken.Form+"/"+parallelVerb.l1Pres)

					for _, child := range dependency.Children {
						if child.Function == "f" { // punctuation
							// skip it
						} else if child.Function == "suj" {
							commands = append(commands,
								buildCommandsForNounPhrase(child, tokenById)...)
							commands = append(commands, "MAKE_AGENT")
						} else if child.Function == "cd" {
							commands = append(commands,
								buildCommandsForNounPhrase(child, tokenById)...)
							commands = append(commands, "MAKE_DOBJ")
						} else if child.Function == "atr" {
							commands = append(commands,
								buildCommandsForAdj(child, tokenById)...)
							commands = append(commands, "ATTACH_ATR_TO_VP")
						}
					}

					stack := commandsPkg.NewStack()
					for _, command := range commands {
						stack.ExecCommand(command)
					}
					l1 := strings.Join(stack.GetL1Words(), " ")
					l2 := strings.Join(stack.GetL2Words(), " ")
					fmt.Printf("%-40s %-39s\n", l2, l1)

					if phraseNum >= 20 {
						os.Exit(1)
					}
				} // end if
			} // next top-level dependency
		} // next sentence
	} // next phrase
}

func buildCommandsForAdj(dependency parsing.Dependency,
	tokenById map[string]parsing.Token) []string {
	command := map[string]string{
		"sentado": "ADD/ADJ/sentado/seated",
		"cansada": "ADD/ADJ/cansada/tired",
		"alta":    "ADD/ADJ/alta/tall",
		"alto":    "ADD/ADJ/alto/tall",
		"gris":    "ADD/ADJ/gris/gray",
		"baja":    "ADD/ADJ/baja/short",
		"negro":   "ADD/ADJ/negro/black",
	}[strings.ToLower(dependency.Word)]
	if command == "" {
		log.Panicf("Unknown adjective: " + dependency.Word)
	}
	return []string{command}
}

func buildCommandsForDet(dependency parsing.Dependency,
	tokenById map[string]parsing.Token) []string {
	command := map[string]string{
		"el":   "ADD/DET/el/the",
		"la":   "ADD/DET/la/the",
		"un":   "ADD/DET/un/a",
		"una":  "ADD/DET/una/a",
		"los":  "ADD/DET/los/the",
		"las":  "ADD/DET/las/the",
		"unos": "ADD/DET/unos/some",
		"unas": "ADD/DET/unas/some",
	}[strings.ToLower(dependency.Word)]
	if command == "" {
		log.Panicf("Unknown determiner: " + dependency.Word)
	}
	return []string{command}
}

func buildCommandsForNounPhrase(dependency parsing.Dependency,
	tokenById map[string]parsing.Token) []string {
	commands := []string{}

	token := tokenById[dependency.Token]
	parallelNoun := parallelNounByL2[token.Lemma]
	if parallelNoun.l2 == "" {
		log.Panicf("Can't find parallelNoun for l2=%s", token.Lemma)
	}

	commands = append(commands,
		"ADD/NOUN/"+parallelNoun.l2+"/"+parallelNoun.l1)

	for _, child := range dependency.Children {
		if child.Function == "spec" {
			commands = append(commands, buildCommandsForDet(child, tokenById)...)
			commands = append(commands, "MAKE_DET_NOUN_PHRASE")
		} else if child.Function == "s.a" {
			commands = append(commands, buildCommandsForAdj(child, tokenById)...)
			commands = append(commands, "MAKE_NOUN_PHRASE_ADJ")
		}
	}

	return commands
}
