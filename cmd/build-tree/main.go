package main

import (
	"fmt"
	"strings"
)

const commands = `
ADD_VERB_INFINITIVE/busc-/seek
ADD_VERB_UNIQUE/está/is
MAKE_PRES_PROG/-ando/-ing
ADD_NOUN/Apple/Apple
MAKE_AGENT
ADD_VERB_INFINITIVE/compr-/buy
MAKE_INFINITIVE/-ar/to
ADD_NOUN/startup/startup
ADD_DET/una/a
MAKE_DET_NOUN
ADD_NOUN/Reino/Kingdom
ADD_ADJ/Unido/United
MAKE_NOUN_ADJ
ADD_DET/el/the
MAKE_DET_NOUN
ADD_PREP/de/from
MAKE_PREP_NOUN
MAKE_NOUN_PHRASE_ADDING_PREP_PHRASE
MAKE_DOBJ
ADD_NOUN/millón/million
MAKE_PLURAL/-es/-s
ADD_NOUN/dólar/dollar
MAKE_PLURAL/-es/-s
ADD_PREP/de/of
MAKE_PREP_NOUN
MAKE_NOUN_PHRASE_ADDING_PREP_PHRASE
ADD_PREP/por/for
MAKE_PREP_NOUN
MAKE_VERB_PHRASE_ADDING_PREP_PHRASE
MAKE_COMPOUND_VERB
`

func main() {
	stack := Stack{}
	for _, commandWhitespace := range strings.Split(commands, "\n") {
		command := strings.TrimSpace(commandWhitespace)
		if command != "" {
			stack.execCommand(command)
		}
	}

	fmt.Printf("%+v\n", stack.getL1Words())
	fmt.Printf("%+v\n", stack.getL2Words())
}
