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
	cd              []NP // direct object
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
	for _, cd := range vp.cd {
		children = append(children, cd)
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
	for _, cd := range vp.cd {
		tokens = append(tokens, cd.GetAllTokens()...)
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

func translateVerb(verb parsing.Token, dictionary english.Dictionary) string {
	if verb.Lemma == "estar" || verb.Lemma == "ser" {
		return map[string]string{
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
		}[verb.Tag]
	} else {
		en := dictionary.Lookup(strings.ToLower(verb.Lemma), "v")
		if verb.Tense == "present" &&
			verb.Num == "singular" &&
			verb.Person == "3" {
			en = english.ConjugateVerb(en, english.PRES_S)
		}
		return en
	}
}

func (vp VP) Translate(dictionary english.Dictionary) []string {
	var l1 []string

	for _, suj := range vp.suj {
		l1 = append(l1, suj.Translate(dictionary)...)
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
		l1 = append(l1, pronoun)
	}

	l1 = append(l1, translateVerb(vp.verb, dictionary))

	for _, cd := range vp.cd {
		l1 = append(l1, cd.Translate(dictionary)...)
	}

	for _, atrAdj := range vp.atrAdj {
		if atrAdj.IsAdjective() {
			l1 = append(l1, dictionary.Lookup(atrAdj.Lemma, "adj"))
		} else if atrAdj.IsVerb() {
			verb := dictionary.Lookup(atrAdj.Lemma, "v")
			l1 = append(l1, english.ConjugateVerb(verb, english.PAST_PART))
		}
	}

	for _, atrPP := range vp.atrPP {
		l1 = append(l1, atrPP.Translate(dictionary)...)
	}

	for _, atrNP := range vp.atrNP {
		l1 = append(l1, atrNP.Translate(dictionary)...)
	}

	return l1
}

func depToVP(dep parsing.Dependency,
	tokenById map[string]parsing.Token) (VP, error) {
	var suj []NP
	var cd []NP
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
			np, err := depToNP(child, tokenById)
			if err != nil {
				return VP{}, err
			}
			cd = append(cd, np)
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
			cd = append(cd, np)
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
		suj: suj, cd: cd, ci: ci, se: se,
		atrAdj: atrAdj, atrNP: atrNP, atrPP: atrPP, adverbs: adverbs,
		auxVs: auxVs}, nil
}
