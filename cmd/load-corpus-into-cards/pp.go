package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
)

type PP struct {
	prep parsing.Token
	np   NP
}

func (pp PP) GetType() string { return "PP" }

func (pp PP) GetChildren() []Constituent {
	return []Constituent{pp.np}
}

func (pp PP) GetAllTokens() []parsing.Token {
	tokens := []parsing.Token{}
	tokens = append(tokens, pp.prep)
	tokens = append(tokens, pp.np.GetAllTokens()...)
	return tokens
}

func depToPP(dep parsing.Dependency,
	tokenById map[string]parsing.Token) (PP, error) {
	np := NP{}
	var err error
	for _, child := range dep.Children {
		if child.Function == "sn" {
			np, err = depToNP(child, tokenById)
			if err != nil {
				return PP{}, err
			}
		} else {
			return PP{}, fmt.Errorf("PP child of %s: %v", child.Function, dep)
		}
	}
	return PP{prep: tokenById[dep.Token], np: np}, nil
}
