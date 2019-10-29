package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
)

type S struct {
	punctuations []parsing.Token
	mods         []parsing.Token
	np           []NP
	vp           []VP
	number       []parsing.Token
	date         []parsing.Token
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

func (s S) GetAllTokens() []parsing.Token {
	tokens := []parsing.Token{}
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

func (s S) Translate(dictionary english.Dictionary) []string {
	l1 := []string{}
	for _, np := range s.np {
		l1 = append(l1, np.Translate(dictionary)...)
	}
	for _, vp := range s.vp {
		l1 = append(l1, vp.Translate(dictionary)...)
	}
	return l1
}

func depToS(dep parsing.Dependency,
	tokenById map[string]parsing.Token) (S, error) {
	headToken := tokenById[dep.Token]
	if headToken.IsNoun() {
		punctuations := []parsing.Token{}
		mods := []parsing.Token{}
		nonPunctuationsOrMods := []parsing.Dependency{}
		for _, child := range dep.Children {
			childToken := tokenById[child.Token]
			if childToken.IsPunctuation() {
				punctuations = append(punctuations, childToken)
			} else if child.Function == "mod" {
				mods = append(mods, childToken)
			} else {
				nonPunctuationsOrMods = append(nonPunctuationsOrMods, child)
			}
		}
		newDep := parsing.Dependency{
			Token:    dep.Token,
			Function: dep.Function,
			Word:     dep.Word,
			Children: nonPunctuationsOrMods,
		}
		np, err := depToNP(newDep, tokenById)
		return S{punctuations: punctuations, mods: mods, np: []NP{np}}, err
	} else if headToken.IsVerb() {
		punctuations := []parsing.Token{}
		mods := []parsing.Token{}
		nonPunctuationsOrMods := []parsing.Dependency{}
		for _, child := range dep.Children {
			childToken := tokenById[child.Token]
			if childToken.IsPunctuation() {
				punctuations = append(punctuations, childToken)
			} else if child.Function == "mod" {
				mods = append(mods, childToken)
			} else {
				nonPunctuationsOrMods = append(nonPunctuationsOrMods, child)
			}
		}
		newDep := parsing.Dependency{
			Token:    dep.Token,
			Function: dep.Function,
			Word:     dep.Word,
			Children: nonPunctuationsOrMods,
		}
		vp, err := depToVP(newDep, tokenById)
		return S{punctuations: punctuations, mods: mods, vp: []VP{vp}}, err
	} else if headToken.IsNumber() && len(dep.Children) == 0 {
		return S{number: []parsing.Token{headToken}}, nil
	} else if headToken.IsDate() && len(dep.Children) == 0 {
		return S{date: []parsing.Token{headToken}}, nil
	} else {
		return S{}, fmt.Errorf("S child of tag=%s: %v",
			tokenById[dep.Token].Tag, dep)
	}
}
