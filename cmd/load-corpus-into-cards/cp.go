package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/spacy"
)

type CP struct {
	conj spacy.Token
	vp   VP
}

func (pp CP) GetType() string { return "CP" }

func (pp CP) GetChildren() []Constituent {
	children := []Constituent{}
	children = append(children, pp.vp)
	return children
}

func (pp CP) GetAllTokens() []spacy.Token {
	tokens := []spacy.Token{}
	tokens = append(tokens, pp.conj)
	tokens = append(tokens, pp.vp.GetAllTokens()...)
	return tokens
}

func (pp CP) Translate(dictionary english.Dictionary) ([]string,
	*CantTranslate) {
	l1 := []string{}

	conjL1, err := dictionary.Lookup(pp.conj.Text, "conj")
	if err != nil {
		return nil, &CantTranslate{Message: err.Error(), Token: pp.conj}
	}
	l1 = append(l1, conjL1)

	vpL1, err2 := pp.vp.Translate(dictionary)
	if err2 != nil {
		return nil, err2
	}
	l1 = append(l1, vpL1...)

	return l1, nil
}

func depToCP(dep spacy.Dep) (CP, *CantConvertDep) {

	// Make a copy of VP with everything except the conjunction
	var conj *spacy.Token
	nonConjs := []spacy.Dep{}
	for _, child := range dep.Children {
		if child.Function == "conj" {
			conj = &child.Token
		} else {
			nonConjs = append(nonConjs, child)
		}
	}
	newDep := spacy.Dep{
		Token:    dep.Token,
		Function: dep.Function,
		Children: nonConjs,
	}
	vp, err := depToVP(newDep)
	if err != nil {
		return CP{}, err
	}

	if conj == nil {
		return CP{}, &CantConvertDep{
			Parent:  dep,
			Message: "Can't find conj child",
		}
	}

	return CP{conj: *conj, vp: vp}, nil
}
