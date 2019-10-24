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

	if false {
		for _, phrase := range phrases {
			parse := parsing.LoadSavedParse(phrase, PARSE_DIR).Parse
			for _, sentence := range parse.Sentences {
				for _, token := range sentence.Tokens {
					if strings.HasPrefix(token.Tag, "D") {
						fmt.Printf("%s\n", token.Lemma)
					}
				}
			}
		}
		os.Exit(1)
	}

	for phraseNum, phrase := range phrases {
		log.Printf("Phrase: %s", phrase)
		switch phrase {
		case "La mujer está parada.":
			continue
		}

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
					commands = append(commands,
						translateVerbPhrase(dependency, tokenById)...)

					stack := commandsPkg.NewStack()
					for _, command := range commands {
						stack.ExecCommand(command)
					}
					l1 := strings.Join(stack.GetL1Words(), " ")
					l2 := strings.Join(stack.GetL2Words(), " ")
					fmt.Printf("%-40s %-39s\n", l2, l1)

					if phraseNum >= 30 {
						os.Exit(1)
					}
				} // end if
			} // next top-level dependency
		} // next sentence
	} // next phrase
}
