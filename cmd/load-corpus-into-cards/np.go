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

var L2_SPEC_PRONOUN_TO_L1 = map[string]string{
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

	"algunas":  "some",
	"algunos":  "some",
	"cuál":     "which",
	"cuáles":   "which",
	"cuánto":   "how much",
	"ninguna":  "none",
	"ninguno":  "none",
	"ningunas": "none",
	"ningunos": "none",
	"toda":     "all",
	"todo":     "all",
	"todas":    "all",
	"todos":    "all",
}

var L2_MAIN_PRONOUN_TO_L1 = map[string]string{
	"algo":    "something",
	"algunas": "somethings",
	"algunos": "somethings",
	"él":      "he",
	"ella":    "she",
	"ellos":   "they",
	"la":      "her/it",
	"le":      "him/her",
	"lo":      "him/it",
	"me":      "me",
	"qué":     "what",
	"te":      "you",
	"un":      "one",
	"una":     "one",
	"usted":   "your grace",
	"yo":      "I/me",
}

func translateMainPronoun(form string) (string, error) {
	formLower := strings.ToLower(form)

	l1, ok := L2_MAIN_PRONOUN_TO_L1[formLower]
	if !ok {
		return "", fmt.Errorf("Can't find main pronoun %s", formLower)
	}
	return l1, nil
}

func translateSpecPronoun(form, enNoun string) (string, *CantTranslate) {
	formLower := strings.ToLower(form)

	if formLower == "un" || formLower == "una" {
		return english.IndefiniteArticleFor(enNoun), nil
	}

	l1, ok := L2_SPEC_PRONOUN_TO_L1[formLower]
	if !ok {
		return "", &CantTranslate{
			Message: fmt.Sprintf("Can't find spec pronoun %s", formLower),
		}
	}
	return l1, nil
}

func (np NP) Translate(dictionary english.Dictionary) ([]string,
	*CantTranslate) {
	l1 := []string{}

	var nounEn string
	if np.noun.IsProperNoun() {
		nounEn = np.noun.Form // Leave proper nouns untranslated
	} else {
		var err error
		nounEn, err = translateNoun(np.noun, dictionary)
		if err != nil {
			return nil, &CantTranslate{Message: err.Error(), Token: np.noun}
		}

		if np.noun.Num == "plural" {
			nounEn = english.PluralizeNoun(nounEn)
		}
	}

	for _, spec := range np.spec {
		en, err := translateSpecPronoun(spec.Form, nounEn)
		if err != nil {
			err.Token = spec
			return nil, err
		}
		l1 = append(l1, en)
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
		} else if child.Function == "sp" || child.Function == "cc" {
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
		return translateMainPronoun(noun.Form)
	} else {
		return "", fmt.Errorf("Don't know how to translateNoun '%s' with tag %s",
			noun.Form, noun.Tag)
	}
}
