package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/freeling"
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	"strings"
)

type VP struct {
	verb            parsing.Token
	verbConjugation freeling.Conjugation
	suj             []NP
	cdNP            []NP // direct object
	cdVP            []VP // direct object
	cdPP            []PP // direct object, probably starting with "a"
	ci              []NP // indirect objec
	atrAdj          []parsing.Token
	atrAdv          []parsing.Token
	atrPP           []PP
	atrNP           []NP
	se              []parsing.Token
	adverbs         []parsing.Token
	auxVs           []parsing.Token
}

func (vp VP) GetType() string { return "VP" }

func (vp VP) GetChildren() []Constituent {
	children := []Constituent{}
	for _, suj := range vp.suj {
		children = append(children, suj)
	}
	for _, cdNP := range vp.cdNP {
		children = append(children, cdNP)
	}
	for _, cdVP := range vp.cdVP {
		children = append(children, cdVP)
	}
	for _, cdPP := range vp.cdPP {
		children = append(children, cdPP)
	}
	for _, ci := range vp.ci {
		children = append(children, ci)
	}
	for _, atrPP := range vp.atrPP {
		children = append(children, atrPP)
	}
	for _, atrNP := range vp.atrNP {
		children = append(children, atrNP)
	}
	return children
}

func (vp VP) GetAllTokens() []parsing.Token {
	tokens := []parsing.Token{}
	tokens = append(tokens, vp.verb)
	for _, suj := range vp.suj {
		tokens = append(tokens, suj.GetAllTokens()...)
	}
	for _, cdNP := range vp.cdNP {
		tokens = append(tokens, cdNP.GetAllTokens()...)
	}
	for _, cdVP := range vp.cdVP {
		tokens = append(tokens, cdVP.GetAllTokens()...)
	}
	for _, cdPP := range vp.cdPP {
		tokens = append(tokens, cdPP.GetAllTokens()...)
	}
	for _, ci := range vp.ci {
		tokens = append(tokens, ci.GetAllTokens()...)
	}
	tokens = append(tokens, vp.atrAdj...)
	tokens = append(tokens, vp.atrAdv...)
	for _, atrPP := range vp.atrPP {
		tokens = append(tokens, atrPP.GetAllTokens()...)
	}
	for _, atrNP := range vp.atrNP {
		tokens = append(tokens, atrNP.GetAllTokens()...)
	}
	tokens = append(tokens, vp.se...)
	tokens = append(tokens, vp.adverbs...)
	tokens = append(tokens, vp.auxVs...)
	return tokens
}

func translateVerb(verb parsing.Token,
	dictionary english.Dictionary) (string, error) {
	if verb.Lemma == "estar" || verb.Lemma == "ser" {
		l1, found := map[string]string{
			"VMIP1S0": "am",
			"VMIP1P0": "are",
			"VMIP2S0": "are",
			"VMIP2P0": "are",
			"VMIP3S0": "is",
			"VMIP3P0": "are",
			"VSIP1S0": "am",
			"VSIP1P0": "are",
			"VSIP2S0": "are",
			"VSIP2P0": "are",
			"VSIP3S0": "is",
			"VSIP3P0": "are",
			"VAII1S0": "used to be",
			"VAII1P0": "used to be",
			"VAII2S0": "used to be",
			"VAII2P0": "used to be",
			"VAII3S0": "used to be",
			"VAII3P0": "used to be",
			"VMII1S0": "used to be",
			"VMII1P0": "used to be",
			"VMII2S0": "used to be",
			"VMII2P0": "used to be",
			"VMII3S0": "used to be",
			"VMII3P0": "used to be",
		}[verb.Tag]
		if !found {
			return "", fmt.Errorf("Can't find verb for tag %s", verb.Tag)
		}
		return l1, nil
	} else {
		en, err := dictionary.Lookup(strings.ToLower(verb.Lemma), "v")
		if err != nil {
			return "", err
		}

		if verb.Tense == "present" &&
			verb.Num == "singular" &&
			verb.Person == "3" {
			en = english.ConjugateVerb(en, english.PRES_S)
		} else if verb.Tense == "past" {
			en = english.ConjugateVerb(en, english.PAST)
		} else if verb.Mood == "infinitive" {
			en = "to " + en
		} else if verb.Mood == "participle" {
			en = english.ConjugateVerb(en, english.PAST_PART)
		} else if verb.Tense == "conditional" {
			en = "would " + en
		} else if verb.Mood == "gerund" {
			en = english.ConjugateVerb(en, english.GERUND)
		}
		return en, nil
	}
}

func (vp VP) Translate(dictionary english.Dictionary) ([]string, error) {
	var l1 []string

	for _, suj := range vp.suj {
		sujL1, err := suj.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, sujL1...)
	}

	if len(vp.suj) == 0 {
		pronoun := map[string]string{
			"1singular": "I",
			"1plural":   "we",
			"2singular": "you",
			"2plural":   "you",
			"3singular": "he/she/it",
			"3plural":   "they",
		}[vp.verb.Person+vp.verb.Num]
		if pronoun != "" {
			l1 = append(l1, pronoun)
		}
	}

	for _, auxV := range vp.auxVs {
		verbL1, err := translateVerb(auxV, dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, verbL1)
	}

	verbL1, err := translateVerb(vp.verb, dictionary)
	if err != nil {
		return nil, err
	}
	l1 = append(l1, verbL1)

	for _, cdNP := range vp.cdNP {
		cdL1, err := cdNP.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, cdL1...)
	}

	for _, cdVP := range vp.cdVP {
		cdL1, err := cdVP.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, cdL1...)
	}

	for _, cdPP := range vp.cdPP {
		cdL1, err := cdPP.Translate(dictionary)
		log.Printf("TRANSLATION OF PP %+v", cdL1)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, cdL1...)
	}

	for _, atrAdj := range vp.atrAdj {
		if atrAdj.IsAdjective() {
			adjL1, err := dictionary.Lookup(atrAdj.Lemma, "adj")
			if err != nil {
				return nil, err
			}
			l1 = append(l1, adjL1)
		} else if atrAdj.IsVerb() {
			verbL1, err := dictionary.Lookup(atrAdj.Lemma, "v")
			if err != nil {
				return nil, err
			}
			l1 = append(l1, english.ConjugateVerb(verbL1, english.PAST_PART))
		} else {
			return nil, fmt.Errorf("Don't know how to translate atrAdj")
		}
	}

	for _, atrPP := range vp.atrPP {
		atrPPL1, err := atrPP.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, atrPPL1...)
	}

	for _, atrNP := range vp.atrNP {
		atrNPL1, err := atrNP.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, atrNPL1...)
	}

	return l1, nil
}

func depToVP(dep parsing.Dependency,
	tokenById map[string]parsing.Token) (VP, error) {
	var suj []NP
	var cdNP []NP
	var cdVP []VP
	var cdPP []PP
	var ci []NP
	var atrAdj []parsing.Token
	var atrAdv []parsing.Token
	var atrPP []PP
	var atrNP []NP
	var se []parsing.Token
	var adverbs []parsing.Token
	var auxVs []parsing.Token
	for _, child := range dep.Children {
		childToken := tokenById[child.Token]
		if child.Function == "suj" {
			np, err := depToNP(child, tokenById)
			if err != nil {
				return VP{}, err
			}
			suj = append(suj, np)
		} else if child.Function == "cd" {

			if childToken.IsNoun() {
				np, err := depToNP(child, tokenById)
				if err != nil {
					return VP{}, err
				}
				cdNP = append(cdNP, np)
			} else if childToken.IsVerb() {
				vp, err := depToVP(child, tokenById)
				if err != nil {
					return VP{}, err
				}
				cdVP = append(cdVP, vp)
			} else if childToken.IsPreposition() {
				pp, err := depToPP(child, tokenById)
				if err != nil {
					return VP{}, err
				}
				cdPP = append(cdPP, pp)
			} else {
				return VP{}, fmt.Errorf(
					"VP's child of cd has unexpected tag : %v/%s", dep, childToken.Tag)
			}
		} else if child.Function == "ci" {
			np, err := depToNP(child, tokenById)
			if err != nil {
				return VP{}, err
			}
			ci = append(ci, np)
		} else if (child.Function == "atr" || child.Function == "cc") &&
			childToken.Tag == "SP" &&
			len(child.Children) == 1 &&
			(child.Children[0].Function == "sn" ||
				child.Children[0].Function == "sadv") {
			pp, err := depToPP(child, tokenById)
			if err != nil {
				return VP{}, err
			}
			atrPP = append(atrPP, pp)
		} else if child.Function == "atr" && childToken.IsNoun() {
			np, err := depToNP(child, tokenById)
			if err != nil {
				return VP{}, err
			}
			atrNP = append(atrNP, np)
		} else if child.Function == "atr" && len(child.Children) == 0 &&
			childToken.IsAdjective() {
			atrAdj = append(atrAdj, childToken)
		} else if child.Function == "atr" && len(child.Children) == 0 &&
			strings.HasPrefix(childToken.Tag, "VMP") {
			atrAdj = append(atrAdj, childToken)
		} else if child.Function == "atr" && len(child.Children) == 0 &&
			childToken.IsAdverb() {
			atrAdj = append(atrAdv, childToken)
		} else if child.Function == "cpred" &&
			(childToken.IsAdjective() || childToken.IsNoun()) {
			// Treat an adjective (chino) as a noun
			np, err := depToNP(child, tokenById)
			if err != nil {
				return VP{}, err
			}
			cdNP = append(cdNP, np)
		} else if child.Function == "pass" && strings.ToLower(child.Word) == "se" {
			se = append(se, childToken)
		} else if child.Function == "cc" && childToken.IsAdverb() {
			adverbs = append(adverbs, childToken)
		} else if child.Function == "v" && len(child.Children) == 0 {
			auxVs = append(auxVs, childToken)
		} else {
			return VP{}, fmt.Errorf(
				"VP child of %s: %v/%s", child.Function, dep, childToken.Tag)
		}
	}

	token := tokenById[dep.Token]
	conjugations := freeling.AnalyzeVerb(token.Lemma, token.Tag)
	if len(conjugations) == 0 {
		return VP{}, fmt.Errorf("No conjugations of %v", dep)
	}
	conjugation := conjugations[0]

	return VP{verb: tokenById[dep.Token], verbConjugation: conjugation,
		suj: suj, cdNP: cdNP, cdVP: cdVP, cdPP: cdPP, ci: ci, se: se,
		atrAdj: atrAdj, atrNP: atrNP, atrPP: atrPP, adverbs: adverbs,
		auxVs: auxVs}, nil
}
