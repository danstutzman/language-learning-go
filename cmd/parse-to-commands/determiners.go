package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"log"
	"strings"
)

type ParallelDet struct {
	l2 string
	l1 string
}

var parallelDets = []ParallelDet{
	{"el", "the"},
	{"la", "the"},
	{"un", "a"},
	{"una", "a"},
	{"los", "the"},
	{"las", "the"},
	{"unos", "some"},
	{"unas", "some"},
	{"alguno", "any"},
	{"cuánto", "how much"},
	{"ese", "that"},
	{"este", "this"},
	{"mi", "my"},
	{"mucho", "much"},
	{"nuestro", "our"},
	{"poco", "little"},
	{"qué", "than"},
	{"su", "his/her/its"},
	{"todo", "all"},
	{"tu", "your"},
	{"uno", "one"},
}

var parallelDetByL2 = buildParallelDetByL2()

func buildParallelDetByL2() map[string]ParallelDet {
	parallelDetByL2 := map[string]ParallelDet{}
	for _, parallelDet := range parallelDets {
		parallelDetByL2[parallelDet.l2] = parallelDet
	}
	return parallelDetByL2
}

func buildCommandsForDet(dependency parsing.Dependency,
	tokenById map[string]parsing.Token) []string {
	parallelDet := parallelDetByL2[strings.ToLower(dependency.Word)]
	if parallelDet.l2 == "" {
		log.Panicf("Unknown determiner %s", dependency.Word)
	}
	return []string{"ADD/DET/" + dependency.Word + "/" + parallelDet.l1}
}
