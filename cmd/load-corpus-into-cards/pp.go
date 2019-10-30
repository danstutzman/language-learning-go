package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
)

type PP struct {
	prep parsing.Token
	np   []NP
	vp   []VP
}

func (pp PP) GetType() string { return "PP" }

func (pp PP) GetChildren() []Constituent {
	children := []Constituent{}
	for _, np := range pp.np {
		children = append(children, np)
	}
	for _, vp := range pp.vp {
		children = append(children, vp)
	}
	return children
}

func (pp PP) GetAllTokens() []parsing.Token {
	tokens := []parsing.Token{}
	tokens = append(tokens, pp.prep)
	for _, np := range pp.np {
		tokens = append(tokens, np.GetAllTokens()...)
	}
	for _, vp := range pp.vp {
		tokens = append(tokens, vp.GetAllTokens()...)
	}
	return tokens
}

func (pp PP) Translate(dictionary english.Dictionary) ([]string, error) {
	l1 := []string{}

	prepL1, err := dictionary.Lookup(pp.prep.Form, "prep")
	if err != nil {
		return nil, err
	}
	l1 = append(l1, prepL1)

	for _, np := range pp.np {
		npL1, err := np.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, npL1...)
	}

	for _, vp := range pp.vp {
		vpL1, err := vp.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, vpL1...)
	}

	return l1, nil
}

func depToPP(dep parsing.Dependency,
	tokenById map[string]parsing.Token) (PP, error) {
	var np []NP
	var vp []VP
	for _, child := range dep.Children {
		childToken := tokenById[child.Token]

		if child.Function == "sn" {
			newNp, err := depToNP(child, tokenById)
			if err != nil {
				return PP{}, err
			}
			np = append(np, newNp)
		} else if child.Function == "S" && childToken.IsVerb() {
			newVp, err := depToVP(child, tokenById)
			if err != nil {
				return PP{}, err
			}
			vp = append(vp, newVp)
		} else {
			return PP{}, fmt.Errorf("PP child of %s: %v", child.Function, dep)
		}
	}
	return PP{prep: tokenById[dep.Token], np: np, vp: vp}, nil
}
