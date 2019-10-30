package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
)

type CP struct {
	conj parsing.Token
	vp   VP
}

func (pp CP) GetType() string { return "CP" }

func (pp CP) GetChildren() []Constituent {
	children := []Constituent{}
	children = append(children, pp.vp)
	return children
}

func (pp CP) GetAllTokens() []parsing.Token {
	tokens := []parsing.Token{}
	tokens = append(tokens, pp.conj)
	tokens = append(tokens, pp.vp.GetAllTokens()...)
	return tokens
}

func (pp CP) Translate(dictionary english.Dictionary) ([]string, error) {
	l1 := []string{}

	conjL1, err := dictionary.Lookup(pp.conj.Form, "conj")
	if err != nil {
		return nil, err
	}
	l1 = append(l1, conjL1)

	vpL1, err := pp.vp.Translate(dictionary)
	if err != nil {
		return nil, err
	}
	l1 = append(l1, vpL1...)

	return l1, nil
}

func depToCP(dep parsing.Dependency,
	tokenById map[string]parsing.Token) (CP, error) {

	// Make a copy of VP with everything except the conjunction
	var conj parsing.Token
	nonConjs := []parsing.Dependency{}
	for _, child := range dep.Children {
		if child.Function == "conj" {
			conj = tokenById[child.Token]
		} else {
			nonConjs = append(nonConjs, child)
		}
	}
	newDep := parsing.Dependency{
		Token:    dep.Token,
		Function: dep.Function,
		Word:     dep.Word,
		Children: nonConjs,
	}
	vp, err := depToVP(newDep, tokenById)
	if err != nil {
		return CP{}, err
	}

	if conj.Id == "" {
		return CP{}, fmt.Errorf("Can't find conj child of %v", dep)
	}

	return CP{conj: conj, vp: vp}, nil
}
