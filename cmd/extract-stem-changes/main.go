package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/freeling"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1+1 {
		log.Fatalf("1st arg: path to dicc.src")
	}
	freelingDiccPath := os.Args[1]

	freeling.PrintVerbExceptions(freelingDiccPath)
}
