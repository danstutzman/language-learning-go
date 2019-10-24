package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"fmt"
	"strings"
)

type ParallelVerb struct {
	l2     string
	l1     string
	l1Past string
	l1Pres string
}

var parallelVerbs = []ParallelVerb{
	{"beber", "drink", "drank", "drinks"},
	{"cepillar", "brush", "brushed", "brushes"},
	{"comer", "eat", "ate", "eats"},
	{"comprar", "buy", "bought", "buys"},
	{"costar", "cost", "cost", "costs"},
	{"cumplir", "comply", "complied", "complies"},
	{"despertar", "wake", "woke", "wakes"},
	{"dormir", "sleep", "slept", "sleeps"},
	{"enseñar", "teach", "taught", "teaches"},
	{"escribir", "write", "wrote", "writes"},
	{"estar", "be", "was", "is"},
	{"estudiar", "study", "studied", "studies"},
	{"gustar", "please", "pleased", "pleases"},
	{"haber", "have", "had", "has"},
	{"hablar", "talk", "talked", "talks"},
	{"jugar", "play", "played", "plays"},
	{"lavar", "wash", "washed", "washes"},
	{"leer", "read", "read", "reads"},
	{"llamar", "call", "called", "calls"},
	{"llevar", "wear", "wore", "wears"},
	{"necesitar", "need", "needed", "needs"},
	{"oler", "smell", "smelled", "smells"},
	{"pagar", "pay", "paid", "pays"},
	{"probar", "try", "tred", "tries"},
	{"quedar", "stay", "stayed", "stays"},
	{"querer", "want", "wanted", "wants"},
	{"tener", "have", "had", "has"},
	{"trabajar", "work", "worked", "works"},
	{"vender", "sell", "sold", "sells"},
	{"venir", "come", "came", "comes"},
	{"ver", "watch", "watched", "watches"},
	{"vivir", "live", "lived", "lives"},

	{"cansar", "tire", "tired", "tires"},
	{"sentar", "sit", "sat", "sits"},
}

var parallelVerbByL2 = buildParallelVerbByL2()

func buildParallelVerbByL2() map[string]ParallelVerb {
	parallelVerbByL2 := map[string]ParallelVerb{}
	for _, parallelVerb := range parallelVerbs {
		parallelVerbByL2[parallelVerb.l2] = parallelVerb
	}
	return parallelVerbByL2
}

func translateVerbPastParticiple(dependency parsing.Dependency,
	tokenById map[string]parsing.Token) ([]string, error) {
	commands := []string{}

	token := tokenById[dependency.Token]
	parallelVerb := parallelVerbByL2[token.Lemma]
	if parallelVerb.l2 == "" {
		return nil, fmt.Errorf("Can't find parallelVerb for l2=%s", token.Lemma)
	}

	commands = append(commands,
		"ADD/ADJ/"+dependency.Word+"/"+parallelVerb.l1Past)

	return commands, nil
}

func translateVerbPhrase(dependency parsing.Dependency,
	tokenById map[string]parsing.Token) ([]string, error) {

	token := tokenById[dependency.Token]
	parallelVerb := parallelVerbByL2[token.Lemma]
	commands := []string{"ADD/VERB/" + token.Form + "/" + parallelVerb.l1Pres}

	for _, child := range dependency.Children {
		if child.Function == "f" { // punctuation
			// skip it
		} else if child.Function == "suj" {
			newCommands, err := translateNounPhrase(child, tokenById)
			if err != nil {
				return nil, err
			}
			commands = append(commands, newCommands...)
			commands = append(commands, "MAKE_AGENT")
		} else if child.Function == "cd" {
			newCommands, err := translateNounPhrase(child, tokenById)
			if err != nil {
				return nil, err
			}
			commands = append(commands, newCommands...)
			commands = append(commands, "MAKE_DOBJ")
		} else if child.Function == "atr" {
			childTag := tokenById[child.Token].Tag
			if strings.HasPrefix(childTag, "A") {
				newCommands, err := translateAdj(child, tokenById)
				if err != nil {
					return nil, err
				}
				commands = append(commands, newCommands...)
			} else if strings.HasPrefix(childTag, "VMP") {
				newCommands, err := translateVerbPastParticiple(child, tokenById)
				if err != nil {
					return nil, err
				}
				commands = append(commands, newCommands...)
			} else {
				return nil, fmt.Errorf("Can't handle atr for child %v with tag %s",
					child, childTag)
			}
			commands = append(commands, "ATTACH_ATR_TO_VP")
		}
	}
	return commands, nil
}
