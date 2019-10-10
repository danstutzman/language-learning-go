package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: path to stories.yaml
		Argument 2: hostname:port for freeling server`)
	}
	storiesYamlPath := os.Args[1]
	freelingHostAndPort := os.Args[2]

	parseDir := "db/1_parses"

	phrases := parsing.ImportStoriesYaml(storiesYamlPath, parseDir)

	outputs := parsing.ParsePhrasesWithFreeling(phrases, freelingHostAndPort)

	for _, output := range outputs {
		parsing.SaveParse(output.Phrase, output.ParseJson, parseDir)
		fmt.Fprintf(os.Stderr, "%s\n", parseDir+"/"+output.Phrase+".json")
	}
}
