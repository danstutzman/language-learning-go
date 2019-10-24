package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2+1 { // Args[0] is name of program
		log.Fatalf(`Usage:
		Argument 1: hostname:port for freeling server
		Argument 2: phrase to parse`)
	}
	freelingHostAndPort := os.Args[1]
	phrase := os.Args[2]

	outputs := parsing.ParsePhrasesWithFreeling(
		[]string{phrase}, freelingHostAndPort)
	fmt.Println(outputs[0].ParseJson)
}
