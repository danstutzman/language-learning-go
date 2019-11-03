package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/spacy"
	"fmt"
	"os"
	"sort"
)

func printTokensInOrder(where *os.File, tokens []spacy.Token) {
	sort.SliceStable(tokens, func(i, j int) bool {
		return tokens[i].Idx < tokens[j].Idx
	})
	for i, token := range tokens {
		thisBegin := token.Idx
		if i > 0 {
			prevEnd := tokens[i-1].Idx + len([]rune(tokens[i-1].Text))
			if thisBegin == prevEnd+1 {
				fmt.Fprintf(where, " ")
			} else if thisBegin > prevEnd+1 {
				fmt.Fprintf(where, " ... ")
			}
		}
		fmt.Fprintf(where, "%s", token.Text)
	}
	fmt.Fprintf(where, "\n")
}

type CantTranslate struct {
	Token   spacy.Token
	Message string
}

type CantConvertDep struct {
	Parent  spacy.Dep
	Child   spacy.Dep
	Message string
}

type Constituent interface {
	GetAllTokens() []spacy.Token
	GetType() string
	GetChildren() []Constituent
	Translate(dictionary english.Dictionary) ([]string, *CantTranslate)
}
