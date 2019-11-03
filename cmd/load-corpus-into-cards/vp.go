package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/english"
	"bitbucket.org/danstutzman/language-learning-go/internal/freeling"
	"bitbucket.org/danstutzman/language-learning-go/internal/spacy"
	"fmt"
	"strings"
)

type VP struct {
	verb            spacy.Token
	verbConjugation freeling.Conjugation
	nsubj           []NP
	objNP           []NP // direct object
	objVP           []VP // direct object
	objPP           []PP // direct object, probably starting with "a"
	objCP           []CP // direct object, probably starting with "que"
	ci              []NP // indirect objec
	atrAdj          []spacy.Token
	atrAdv          []spacy.Token
	atrPP           []PP
	atrNP           []NP
	atrVP           []VP // usually a participle?
	ccPP            []PP
	ccCP            []CP
	creg            []PP // "prepositional complement"
	se              []spacy.Token
	adverbs         []spacy.Token
	auxVs           []spacy.Token
}

func (vp VP) GetType() string { return "VP" }

func (vp VP) GetChildren() []Constituent {
	children := []Constituent{}
	for _, nsubj := range vp.nsubj {
		children = append(children, nsubj)
	}
	for _, objNP := range vp.objNP {
		children = append(children, objNP)
	}
	for _, objVP := range vp.objVP {
		children = append(children, objVP)
	}
	for _, objPP := range vp.objPP {
		children = append(children, objPP)
	}
	for _, objCP := range vp.objCP {
		children = append(children, objCP)
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
	for _, atrVP := range vp.atrVP {
		children = append(children, atrVP)
	}
	for _, ccPP := range vp.ccPP {
		children = append(children, ccPP)
	}
	for _, ccCP := range vp.ccCP {
		children = append(children, ccCP)
	}
	for _, creg := range vp.creg {
		children = append(children, creg)
	}
	return children
}

func (vp VP) GetAllTokens() []spacy.Token {
	tokens := []spacy.Token{}
	tokens = append(tokens, vp.verb)
	for _, nsubj := range vp.nsubj {
		tokens = append(tokens, nsubj.GetAllTokens()...)
	}
	for _, objNP := range vp.objNP {
		tokens = append(tokens, objNP.GetAllTokens()...)
	}
	for _, objVP := range vp.objVP {
		tokens = append(tokens, objVP.GetAllTokens()...)
	}
	for _, objPP := range vp.objPP {
		tokens = append(tokens, objPP.GetAllTokens()...)
	}
	for _, objCP := range vp.objCP {
		tokens = append(tokens, objCP.GetAllTokens()...)
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
	for _, atrVP := range vp.atrVP {
		tokens = append(tokens, atrVP.GetAllTokens()...)
	}
	for _, ccPP := range vp.ccPP {
		tokens = append(tokens, ccPP.GetAllTokens()...)
	}
	for _, ccCP := range vp.ccCP {
		tokens = append(tokens, ccCP.GetAllTokens()...)
	}
	for _, creg := range vp.creg {
		tokens = append(tokens, creg.GetAllTokens()...)
	}
	tokens = append(tokens, vp.se...)
	tokens = append(tokens, vp.adverbs...)
	tokens = append(tokens, vp.auxVs...)
	return tokens
}

func translateVerb(verb spacy.Token,
	dictionary english.Dictionary) (string, *CantTranslate) {
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
		}[verb.VerbTag]
		if !found {
			return "", &CantTranslate{
				Message: fmt.Sprintf("Can't find verb for tag %s", verb.VerbTag),
				Token:   verb,
			}
		}
		return l1, nil
	} else {
		en, err := dictionary.Lookup(strings.ToLower(verb.Lemma), "v")
		if err != nil {
			return "", &CantTranslate{Message: err.Error(), Token: verb}
		}

		if verb.Features["Tense"] == "Pres" &&
			verb.Features["Number"] == "Sing" &&
			verb.Features["Person"] == "3" {
			en = english.ConjugateVerbPhrase(en, english.PRES_S)
		} else if verb.Features["Tense"] == "Past" {
			en = english.ConjugateVerbPhrase(en, english.PAST)
		} else if verb.Features["Mood"] == "Inf" {
			en = "to " + en
			//} else if verb.Mood == "participle" {
			//	en = english.ConjugateVerbPhrase(en, english.PAST_PART)
		} else if verb.Features["Mood"] == "Cnd" {
			en = "would " + en
		} else if verb.Features["VerbForm"] == "Ger" {
			en = english.ConjugateVerbPhrase(en, english.GERUND)
		} else if verb.Features["Tense"] == "Imp" {
			en = "used to " + en
		}
		return en, nil
	}
}

func (vp VP) Translate(dictionary english.Dictionary) ([]string,
	*CantTranslate) {
	var l1 []string

	for _, nsubj := range vp.nsubj {
		sujL1, err := nsubj.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, sujL1...)
	}

	if len(vp.nsubj) == 0 {
		pronoun := map[string]string{
			"1sing": "I",
			"1plur": "we",
			"2sing": "you",
			"2plur": "you",
			"3sing": "he/she/it",
			"3plur": "they",
		}[vp.verb.Features["Person"]+vp.verb.Features["Number"]]
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

	for _, objNP := range vp.objNP {
		objL1, err := objNP.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, objL1...)
	}

	for _, objVP := range vp.objVP {
		objL1, err := objVP.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, objL1...)
	}

	for _, objPP := range vp.objPP {
		objL1, err := objPP.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, objL1...)
	}

	for _, objCP := range vp.objCP {
		objL1, err := objCP.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, objL1...)
	}

	for _, atrAdj := range vp.atrAdj {
		if atrAdj.Pos == "ADJ" {
			adjL1, err := dictionary.Lookup(atrAdj.Lemma, "adj")
			if err != nil {
				return nil, &CantTranslate{Message: err.Error(), Token: atrAdj}
			}
			l1 = append(l1, adjL1)
		} else {
			return nil, &CantTranslate{
				Message: fmt.Sprintf("Don't know how to translate atrAdj with pos=%s",
					atrAdj.Pos),
				Token: atrAdj,
			}
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

	for _, atrVP := range vp.atrVP {
		atrVPL1, err := atrVP.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, atrVPL1...)
	}

	for _, ccPP := range vp.ccPP {
		ccPPL1, err := ccPP.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, ccPPL1...)
	}

	for _, ccCP := range vp.ccCP {
		ccCPL1, err := ccCP.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, ccCPL1...)
	}

	for _, creg := range vp.creg {
		cregL1, err := creg.Translate(dictionary)
		if err != nil {
			return nil, err
		}
		l1 = append(l1, cregL1...)
	}

	return l1, nil
}

func depToVP(dep spacy.Dep) (VP, *CantConvertDep) {
	var nsubj []NP
	var objNP []NP
	var objVP []VP
	var objPP []PP
	var objCP []CP
	var ci []NP
	var atrAdj []spacy.Token
	var atrAdv []spacy.Token
	var atrPP []PP
	var atrNP []NP
	var atrVP []VP
	var ccPP []PP
	var ccCP []CP
	var creg []PP
	var se []spacy.Token
	var adverbs []spacy.Token
	var auxVs []spacy.Token
	for _, child := range dep.Children {
		if child.Function == "nsubj" {
			np, err := depToNP(child)
			if err != nil {
				return VP{}, err
			}
			nsubj = append(nsubj, np)
		} else if child.Function == "obj" {
			if child.Token.Pos == "NOUN" || child.Token.Pos == "PRON" {
				np, err := depToNP(child)
				if err != nil {
					return VP{}, err
				}
				objNP = append(objNP, np)
			} else if child.Token.Pos == "VERB" {
				if hasConjChild(child) {
					cp, err := depToCP(child)
					if err != nil {
						return VP{}, err
					}
					objCP = append(objCP, cp)
				} else {
					vp, err := depToVP(child)
					if err != nil {
						return VP{}, err
					}
					objVP = append(objVP, vp)
				}
			} else if child.Token.Pos == "ADV" {
				pp, err := depToPP(child)
				if err != nil {
					return VP{}, err
				}
				objPP = append(objPP, pp)
			} else {
				return VP{}, &CantConvertDep{
					Parent: dep,
					Child:  child,
					Message: fmt.Sprintf("obj child has unexpected pos %s",
						child.Token.Pos),
				}
			}
		} else if child.Function == "ci" {
			np, err := depToNP(child)
			if err != nil {
				return VP{}, err
			}
			ci = append(ci, np)
			/*} else if child.Function == "atr" && childToken.Pos == "SP" {
			pp, err := depToPP(child)
			if err != nil {
				return VP{}, err
			}
			atrPP = append(atrPP, pp)*/
		} else if child.Function == "atr" && child.Token.Pos == "NOUN" {
			np, err := depToNP(child)
			if err != nil {
				return VP{}, err
			}
			atrNP = append(atrNP, np)
		} else if child.Function == "atr" && len(child.Children) == 0 &&
			child.Token.Pos == "ADJ" {
			atrAdj = append(atrAdj, child.Token)
		} else if child.Function == "atr" && child.Token.Pos == "VERB" {
			vp, err := depToVP(child)
			if err != nil {
				return VP{}, err
			}
			atrVP = append(atrVP, vp)
		} else if child.Function == "atr" && len(child.Children) == 0 &&
			child.Token.Pos == "ADV" {
			atrAdj = append(atrAdv, child.Token)
		} else if child.Function == "cpred" &&
			(child.Token.Pos == "ADV" || child.Token.Pos == "NOUN") {
			// Treat an adjective (chino) as a noun
			np, err := depToNP(child)
			if err != nil {
				return VP{}, err
			}
			objNP = append(objNP, np)
		} else if child.Function == "pass" &&
			strings.ToLower(child.Token.Text) == "se" {
			se = append(se, child.Token)
		} else if child.Function == "cc" {
			if child.Token.Pos == "ADV" {
				adverbs = append(adverbs, child.Token)
			} else if child.Token.Pos == "ADP" {
				pp, err := depToPP(child)
				if err != nil {
					return VP{}, err
				}
				ccPP = append(ccPP, pp)
			} else if child.Token.Pos == "VERB" {
				cp, err := depToCP(child)
				if err != nil {
					return VP{}, err
				}
				ccCP = append(ccCP, cp)
			} else {
				return VP{}, &CantConvertDep{
					Parent: dep,
					Child:  child,
					Message: fmt.Sprintf("Child cc has unexpected Pos %s",
						child.Token.Pos),
				}
			}
		} else if child.Function == "v" && len(child.Children) == 0 {
			auxVs = append(auxVs, child.Token)
		} else if child.Function == "creg" && child.Token.Pos == "ADP" {
			pp, err := depToPP(child)
			if err != nil {
				return VP{}, err
			}
			creg = append(creg, pp)
		} else {
			return VP{}, &CantConvertDep{
				Parent:  dep,
				Child:   child,
				Message: fmt.Sprintf("VP's child has function %s", child.Function),
			}
		}
	}

	token := dep.Token
	conjugations := freeling.AnalyzeVerb(token.Lemma, token.VerbTag)
	if len(conjugations) == 0 {
		return VP{}, &CantConvertDep{
			Parent: dep,
			Message: fmt.Sprintf("No conjugations of %s/%s",
				token.Lemma, token.VerbTag),
		}
	}
	conjugation := conjugations[0]

	return VP{verb: dep.Token, verbConjugation: conjugation,
		nsubj: nsubj,
		objNP: objNP, objVP: objVP, objPP: objPP, objCP: objCP, ci: ci, se: se,
		atrAdj: atrAdj, atrNP: atrNP, atrPP: atrPP, atrVP: atrVP, adverbs: adverbs,
		ccPP: ccPP, ccCP: ccCP, creg: creg, auxVs: auxVs}, nil
}

func hasConjChild(dep spacy.Dep) bool {
	for _, child := range dep.Children {
		if child.Function == "conj" {
			return true
		}
	}
	return false
}
