package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	"regexp"
)

type ParallelAdj struct {
	l2 string
	l1 string
}

var ENDS_WITH_A = regexp.MustCompile("a$")

var parallelAdjs = []ParallelAdj{
	{"alto", "tall"},
	{"amarillo", "yellow"},
	{"antiguo", "old"},
	{"azul", "blue"},
	{"bajo", "low"},
	{"barato", "cheap"},
	{"blanco", "White"},
	{"bueno", "good"},
	{"canoso", "gray-haired"},
	{"caro", "expensive"},
	{"chino", "Chinese"},
	{"despierto", "awake"},
	{"diferente", "different"},
	{"efectivo", "effective"},
	{"enfermo", "sick"},
	{"español", "Spanish"},
	{"feo", "ugly"},
	{"grande", "big"},
	{"gris", "gray"},
	{"igual", "equal"},
	{"inglés", "English"},
	{"joven", "young"},
	{"lento", "slow"},
	{"limpio", "clean"},
	{"liviano", "lightweight"},
	{"maestro", "teacher"},
	{"marrón", "brown"},
	{"mayor", "higher"},
	{"mismo", "same"},
	{"médico", "doctor"},
	{"naranja", "orange"},
	{"negro", "black"},
	{"nuevo", "new"},
	{"pelirrojo", "redheaded"},
	{"pequeño", "small"},
	{"pesado", "heavy"},
	{"rico", "rich"},
	{"rojo", "red"},
	{"rosado", "pink"},
	{"rubio", "blond"},
	{"rápido", "quick"},
	{"seco", "dry"},
	{"sucio", "dirty"},
	{"verde", "green"},
	{"viejo", "old"},
	{"árabe", "Arab"},
}

var parallelAdjByL2 = buildParallelAdjByL2()

func buildParallelAdjByL2() map[string]ParallelAdj {
	parallelAdjByL2 := map[string]ParallelAdj{}
	for _, parallelAdj := range parallelAdjs {
		parallelAdjByL2[parallelAdj.l2] = parallelAdj
	}
	return parallelAdjByL2
}

func translateAdj(dependency parsing.Dependency,
	tokenById map[string]parsing.Token) ([]string, error) {
	masculine := ENDS_WITH_A.ReplaceAllString(dependency.Word, "o")
	parallelAdj := parallelAdjByL2[masculine]
	if parallelAdj.l2 == "" {
		return nil, fmt.Errorf("Unknown adjective %s", dependency.Word)
	}

	return []string{"ADD/ADJ/" + dependency.Word + "/" + parallelAdj.l1}, nil
}
