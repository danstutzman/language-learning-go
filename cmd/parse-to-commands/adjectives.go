package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"log"
	"regexp"
)

type ParallelAdj struct {
	l2 string
	l1 string
}

var ENDS_WITH_A = regexp.MustCompile("a$")

var parallelAdjs = []ParallelAdj{
	{"sentado", "seated"},
	{"cansado", "tired"},
	{"alto", "tall"},
	{"gris", "gray"},
	{"bajo", "short"},
	{"negro", "black"},
}

var parallelAdjByL2 = buildParallelAdjByL2()

func buildParallelAdjByL2() map[string]ParallelAdj {
	parallelAdjByL2 := map[string]ParallelAdj{}
	for _, parallelAdj := range parallelAdjs {
		parallelAdjByL2[parallelAdj.l2] = parallelAdj
	}
	return parallelAdjByL2
}

func buildCommandsForAdj(dependency parsing.Dependency,
	tokenById map[string]parsing.Token) []string {
	masculine := ENDS_WITH_A.ReplaceAllString(dependency.Word, "o")
	parallelAdj := parallelAdjByL2[masculine]
	if parallelAdj.l2 == "" {
		log.Panicf("Unknown adjective %s", dependency.Word)
	}

	return []string{"ADD/ADJ/" + dependency.Word + "/" + parallelAdj.l1}
}
