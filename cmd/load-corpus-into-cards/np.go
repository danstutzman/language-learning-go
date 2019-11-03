package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/spacy"
	"fmt"
	"strings"
)

type NP struct {
	noun  spacy.Token
	det   []spacy.Token
	amods []spacy.Token
	sp    []PP
}

func (np NP) GetType() string { return "NP" }

func (np NP) GetChildren() []Constituent {
	constituents := []Constituent{}
	for _, sp := range np.sp {
		constituents = append(constituents, sp)
	}
	return constituents
}

func (np NP) GetAllTokens() []spacy.Token {
	tokens := []spacy.Token{}
	tokens = append(tokens, np.noun)
	for _, det := range np.det {
		tokens = append(tokens, det)
	}
	for _, amod := range np.amods {
		tokens = append(tokens, amod)
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
			Message: fmt.Sprintf("Can't find det pronoun %s", formLower),
		}
	}
	return l1, nil
}

func (np NP) Translate(dictionary english.Dictionary) ([]string,
	*CantTranslate) {
	l1 := []string{}

	var nounEn string
	if np.noun.Pos == "PROPN" {
		nounEn = np.noun.Text // Leave proper nouns untranslated
	} else {
		var err error
		nounEn, err = translateNoun(np.noun, dictionary)
		if err != nil {
			return nil, &CantTranslate{Message: err.Error(), Token: np.noun}
		}

		if np.noun.Features["Number"] == "Plur" {
			nounEn = english.PluralizeNoun(nounEn)
		}
	}

	for _, det := range np.det {
		en, err := translateSpecPronoun(det.Text, nounEn)
		if err != nil {
			err.Token = det
			return nil, err
		}
		l1 = append(l1, en)
	}

	l1 = append(l1, nounEn)

	return l1, nil
}

func depToNP(dep spacy.Dep) (NP, *CantConvertDep) {
	var det []spacy.Token
	var amods []spacy.Token
	var sp []PP
	for _, child := range dep.Children {
		if child.Function == "det" && len(child.Children) == 0 {
			det = append(det, child.Token)
		} else if child.Function == "det" && len(child.Children) == 1 &&
			child.Children[0].Function == "det" {
			det = append(det, child.Token)
			det = append(det, child.Children[0].Token)
		} else if child.Function == "amod" {
			if len(child.Children) == 0 {
				amods = append(amods, child.Token)
			} else {
				return NP{}, &CantConvertDep{
					Parent:  dep,
					Child:   child,
					Message: "NP child of amod not len 0",
				}
			}
		} else if child.Function == "sp" || child.Function == "cc" {
			pp, err := depToPP(child)
			if err != nil {
				return NP{}, err
			}
			sp = append(sp, pp)
		} else {
			return NP{}, &CantConvertDep{
				Parent: dep,
				Child:  child,
				Message: fmt.Sprintf("NP child has unexpected function '%s'",
					child.Function),
			}
		}
	}
	return NP{noun: dep.Token, det: det, amods: amods}, nil
}

func translateNoun(noun spacy.Token,
	dictionary english.Dictionary) (string, error) {
	if noun.Pos == "NOUN" {
		return dictionary.Lookup(noun.Lemma, "n")
	} else if noun.Pos == "PRON" {
		return translateMainPronoun(noun.Text)
	} else {
		return "", fmt.Errorf("Don't know how to translateNoun '%s' with pos %s",
			noun.Text, noun.Pos)
	}
}
