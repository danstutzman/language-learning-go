package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/spacy"
	"fmt"
)

type PP struct {
	prep spacy.Token
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

func (pp PP) GetAllTokens() []spacy.Token {
	tokens := []spacy.Token{}
	tokens = append(tokens, pp.prep)
	for _, np := range pp.np {
		tokens = append(tokens, np.GetAllTokens()...)
	}
	for _, vp := range pp.vp {
		tokens = append(tokens, vp.GetAllTokens()...)
	}
	return tokens
}

func (pp PP) Translate(dictionary english.Dictionary) ([]string,
	*CantTranslate) {
	l1 := []string{}

	prepL1, err := dictionary.Lookup(pp.prep.Text, "prep")
	if err != nil {
		return nil, &CantTranslate{Message: err.Error(), Token: pp.prep}
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

func depToPP(dep spacy.Dep) (PP, *CantConvertDep) {
	var np []NP
	var vp []VP
	for _, child := range dep.Children {
		if child.Function == "sn" {
			newNp, err := depToNP(child)
			if err != nil {
				return PP{}, err
			}
			np = append(np, newNp)
		} else if child.Function == "S" && child.Token.Pos == "VERB" {
			newVp, err := depToVP(child)
			if err != nil {
				return PP{}, err
			}
			vp = append(vp, newVp)
		} else {
			return PP{}, &CantConvertDep{
				Parent:  dep,
				Child:   child,
				Message: fmt.Sprintf("Unexpected function %s", child.Function),
			}
		}
	}
	return PP{prep: dep.Token, np: np, vp: vp}, nil
}
