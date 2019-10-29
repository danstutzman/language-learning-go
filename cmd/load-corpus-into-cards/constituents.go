package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func printTokensInOrder(where *os.File, tokens []parsing.Token) {
	sort.SliceStable(tokens, func(i, j int) bool {
		return mustAtoi(tokens[i].Begin) < mustAtoi(tokens[j].Begin)
	})
	for i, token := range tokens {
		thisBegin := mustAtoi(token.Begin)
		if i > 0 {
			prevEnd := mustAtoi(tokens[i-1].End)
			if thisBegin == prevEnd+1 {
				fmt.Fprintf(where, " ")
			} else if thisBegin > prevEnd+1 {
				fmt.Fprintf(where, " ... ")
			}
		}
		fmt.Fprintf(where, "%s", token.Form)
	}
	fmt.Fprintf(where, "\n")
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

type Constituent interface {
	GetAllTokens() []parsing.Token
	GetType() string
	GetChildren() []Constituent
	Translate(dictionary english.Dictionary) []string
}
