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

func translatePronoun(form string) string {
	return map[string]string{
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
		"estos":    "these",
		"estas":    "these",
		"ese":      "that",
		"esa":      "that",
		"esos":     "those",
		"esas":     "those",
		"poco":     "little",
		"poca":     "little",
		"pocos":    "little",
		"pocas":    "little",
	}[strings.ToLower(form)]
}

func (np NP) Translate(dictionary english.Dictionary) []string {
	l1 := []string{}

	for _, spec := range np.spec {
		en := translatePronoun(spec.Form)
		if en != "" {
			l1 = append(l1, en)
		}
	}

	nounEn := translateNoun(np.noun, dictionary)
	if np.noun.Num == "plural" {
		nounEn = english.PluralizeNoun(nounEn)
	}
	l1 = append(l1, nounEn)

	return l1
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

func translateNoun(noun parsing.Token, dictionary english.Dictionary) string {
	if noun.IsNoun() {
		return dictionary.Lookup(noun.Lemma, "n")
	} else if noun.IsPronoun() {
		l1 := strings.ToLower(translatePronoun(noun.Form))
		if l1 == "" {
			l1 = dictionary.Lookup(noun.Lemma, "pron")
		}
		return l1
	} else {
		return ""
	}
}
