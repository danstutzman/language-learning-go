package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"log"
	"os"
)

func main() {
	freelingHostAndPort := os.Getenv("FREELING_HOST_AND_PORT")
	if freelingHostAndPort == "" {
		log.Fatalf("Specify FREELING_HOST_AND_PORT env var, for example: 1.2.3.4:5678")
	}

	phrases := []string{"Estoy feliz."}
	analyses := model.AnalyzePhrasesWithFreeling(phrases, freelingHostAndPort)
	log.Println(analyses)
}
