package commands

import (
	"fmt"
	"log"
	"strings"
)

const commands = `
ADD/VERB_STEM/busc-/seek
ADD/VERB_UNIQUE/está/is
MAKE_PRES_PROG/-ando/-ing
ADD/NOUN/Apple/Apple
MAKE_AGENT
ADD/VERB_STEM/compr-/buy
MAKE_INFINITIVE/-ar/to
ADD/NOUN/startup/startup
ADD/DET/una/a
MAKE_DET_NOUN_PHRASE
ADD/NOUN/Reino/Kingdom
ADD/ADJ/Unido/United
MAKE_NOUN_PHRASE_ADJ
ADD/PREP/del/from-the
MAKE_PREP_NOUN
MAKE_NOUN_PHRASE_ADDING_PREP_PHRASE
MAKE_DOBJ
ADD/NOUN/millón/million
MAKE_PLURAL/-es/-s
ADD/DET/mil/thousand
MAKE_DET_NOUN_PHRASE
ADD/NOUN/dólar/dollar
MAKE_PLURAL/-es/-s
ADD/PREP/de/of
MAKE_PREP_NOUN
MAKE_NOUN_PHRASE_ADDING_PREP_PHRASE
ADD/PREP/por/for
MAKE_PREP_NOUN
MAKE_VERB_PHRASE_ADDING_PREP_PHRASE
MAKE_VOBJ
`

func Demo() {
	stack := NewStack()
	for _, commandWhitespace := range strings.Split(commands, "\n") {
		command := strings.TrimSpace(commandWhitespace)
		if command != "" {
			log.Printf("%s", command)
			stack.ExecCommand(command)
		}
	}

	fmt.Printf("%+v\n", stack.GetL1Words())
	fmt.Printf("%+v\n", stack.GetL2Words())
}
