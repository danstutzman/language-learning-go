package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	"strings"
)

type NP struct {
	noun parsing.Token
	spec []parsing.Token
	sa   []parsing.Token // function = "s.a"
	sp   []PP
}

func (np NP) GetType() string { return "NP" }

func (np NP) GetChildren() []Constituent {
	constituents := []Constituent{}
	for _, sp := range np.sp {
		constituents = append(constituents, sp)
	}
	return constituents
}

func (np NP) GetAllTokens() []parsing.Token {
	tokens := []parsing.Token{}
	tokens = append(tokens, np.noun)
	for _, spec := range np.spec {
		tokens = append(tokens, spec)
	}
	for _, sa := range np.sa {
		tokens = append(tokens, sa)
	}
	for _, sp := range np.sp {
		tokens = append(tokens, sp.GetAllTokens()...)
	}
	return tokens
}

var L2_PRONOUN_TO_L1 = map[string]string{
	"un":       "a",
	"una":      "a",
	"unos":     "some",
	"unas":     "some",
	"el":       "the",
	"la":       "the",
	"los":      "the",
	"las":      "the",
	"mi":       "my",
	"mis":      "my",
	"tu":       "your",
	"tus":      "your",
	"su":       "his/her/its",
	"sus":      "his/her/its",
	"nuestro":  "our",
	"nuestros": "our",
	"nuestra":  "our",
	"nuestras": "our",
	"este":     "this",
	"esta":     "this",
	"esto":     "this",
	"estos":    "these",
	"estas":    "these",
	"ese":      "that",
	"esa":      "that",
	"eso":      "that",
	"esos":     "those",
	"esas":     "those",
	"poco":     "little",
	"poca":     "little",
	"pocos":    "little",
	"pocas":    "little",

	"algo":     "something",
	"algunas":  "some",
	"algunos":  "some",
	"cuál":     "which",
	"cuáles":   "which",
	"cuánto":   "how much",
	"él":       "he",
	"ella":     "she",
	"ellos":    "they",
	"le":       "him/her",
	"me":       "me",
	"ninguna":  "none",
	"ninguno":  "none",
	"ningunas": "none",
	"ningunos": "none",
	"qué":      "what",
	"te":       "you",
	"toda":     "all",
	"todo":     "all",
	"todas":    "all",
	"todos":    "all",
	"usted":    "your grace",
	"yo":       "me",
}

func translatePronoun(form string) (string, error) {
	l1, ok := L2_PRONOUN_TO_L1[strings.ToLower(form)]
	if !ok {
		return "", fmt.Errorf("Can't find pronoun %s", strings.ToLower(form))
	}
	return l1, nil
}

func (np NP) Translate(dictionary english.Dictionary) ([]string, error) {
	l1 := []string{}

	for _, spec := range np.spec {
		en, err := translatePronoun(spec.Form)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, en)
	}

	nounEn, err := translateNoun(np.noun, dictionary)
	if err != nil {
		return nil, err
	}
	if np.noun.Num == "plural" {
		nounEn = english.PluralizeNoun(nounEn)
	}
	l1 = append(l1, nounEn)

	return l1, nil
}

func depToNP(dep parsing.Dependency,
	tokenById map[string]parsing.Token) (NP, error) {
	var spec []parsing.Token
	var sa []parsing.Token
	var sp []PP
	for _, child := range dep.Children {
		childToken := tokenById[child.Token]
		if child.Function == "spec" && len(child.Children) == 0 {
			spec = append(spec, childToken)
		} else if child.Function == "spec" && len(child.Children) == 1 &&
			child.Children[0].Function == "d" {
			spec = append(spec, childToken)
			spec = append(spec, tokenById[child.Children[0].Token])
		} else if child.Function == "s.a" {
			if len(child.Children) == 0 {
				sa = append(sa, childToken)
			} else {
				return NP{}, fmt.Errorf("NP child of s.a not len 0: %v", dep)
			}
		} else if child.Function == "sp" {
			pp, err := depToPP(child, tokenById)
			if err != nil {
				return NP{}, err
			}
			sp = append(sp, pp)
		} else {
			return NP{}, fmt.Errorf("NP child of %s: %v", child.Function, dep)
		}
	}
	return NP{noun: tokenById[dep.Token], spec: spec, sa: sa}, nil
}

func translateNoun(noun parsing.Token,
	dictionary english.Dictionary) (string, error) {
	if noun.IsNoun() {
		return dictionary.Lookup(noun.Lemma, "n")
	} else if noun.IsPronoun() {
		return translatePronoun(strings.ToLower(noun.Form))
	} else {
		return "",
			fmt.Errorf("Don't know how to translateNoun with tag %s", noun.Tag)
	}
}
