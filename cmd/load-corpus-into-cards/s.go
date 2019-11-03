package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/spacy"
	"fmt"
)

type S struct {
	punctuations []spacy.Token
	mods         []spacy.Token
	np           []NP
	vp           []VP
	number       []spacy.Token
	date         []spacy.Token
}

func (s S) GetType() string { return "S" }

func (s S) GetChildren() []Constituent {
	constituents := []Constituent{}
	for _, np := range s.np {
		constituents = append(constituents, np)
	}
	for _, vp := range s.vp {
		constituents = append(constituents, vp)
	}
	return constituents
}

func (s S) GetAllTokens() []spacy.Token {
	tokens := []spacy.Token{}
	tokens = append(tokens, s.punctuations...)
	tokens = append(tokens, s.mods...)
	for _, np := range s.np {
		tokens = append(tokens, np.GetAllTokens()...)
	}
	for _, vp := range s.vp {
		tokens = append(tokens, vp.GetAllTokens()...)
	}
	tokens = append(tokens, s.number...)
	tokens = append(tokens, s.date...)
	return tokens
}

func (s S) Translate(dictionary english.Dictionary) ([]string, *CantTranslate) {
	l1 := []string{}

	for _, np := range s.np {
		npL1, err := np.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, npL1...)
	}

	for _, vp := range s.vp {
		vpL1, err := vp.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, vpL1...)
	}

	return l1, nil
}

func depToS(dep spacy.Dep) (S, *CantConvertDep) {
	if dep.Token.Pos == "NOUN" {
		punctuations := []spacy.Token{}
		mods := []spacy.Token{}
		nonPunctuationsOrMods := []spacy.Dep{}
		for _, child := range dep.Children {
			if child.Token.Pos == "PUNCT" {
				punctuations = append(punctuations, child.Token)
			} else if child.Function == "mod" {
				mods = append(mods, child.Token)
			} else {
				nonPunctuationsOrMods = append(nonPunctuationsOrMods, child)
			}
		}
		newDep := spacy.Dep{
			Token:    dep.Token,
			Function: dep.Function,
			Children: nonPunctuationsOrMods,
		}
		np, err := depToNP(newDep)
		return S{punctuations: punctuations, mods: mods, np: []NP{np}}, err
	} else if dep.Token.Pos == "VERB" {
		punctuations := []spacy.Token{}
		mods := []spacy.Token{}
		nonPunctuationsOrMods := []spacy.Dep{}
		for _, child := range dep.Children {
			if child.Token.Pos == "PUNCT" {
				punctuations = append(punctuations, child.Token)
			} else if child.Function == "mod" {
				mods = append(mods, child.Token)
			} else {
				nonPunctuationsOrMods = append(nonPunctuationsOrMods, child)
			}
		}
		newDep := spacy.Dep{
			Token:    dep.Token,
			Function: dep.Function,
			Children: nonPunctuationsOrMods,
		}
		vp, err := depToVP(newDep)
		return S{punctuations: punctuations, mods: mods, vp: []VP{vp}}, err
	} else {
		return S{}, &CantConvertDep{
			Parent:  dep,
			Message: fmt.Sprintf("S's token has unexpected pos %s", dep.Token.Pos),
		}
	}
}
