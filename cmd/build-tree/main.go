package main

import (
	"fmt"
)

type Op struct {
	type_ string
	l2    string
	l1    string
}

func main() {
	stack := Stack{}
	stack.exec(Op{type_: "ADD_VERB_INFINITIVE", l2: "buscar", l1: "to seek"})
	stack.exec(Op{type_: "ADD_VERB_UNIQUE", l2: "está", l1: "is"})
	stack.exec(Op{type_: "MAKE_PRES_PROG", l2: "-ar +ando", l1: "-to +ing"})
	stack.exec(Op{type_: "ADD_NOUN", l2: "Apple", l1: "Apple"})
	stack.exec(Op{type_: "MAKE_AGENT", l2: "", l1: ""})
	stack.exec(Op{type_: "ADD_VERB_INFINITIVE", l2: "comprar", l1: "to buy"})
	stack.exec(Op{type_: "ADD_NOUN", l2: "startup", l1: "startup"})
	stack.exec(Op{type_: "ADD_DET", l2: "una", l1: "a"})
	stack.exec(Op{type_: "MAKE_DET_NOUN", l2: "", l1: ""})
	stack.exec(Op{type_: "ADD_NOUN", l2: "Reino", l1: "Kingdom"})
	stack.exec(Op{type_: "ADD_ADJ", l2: "Unido", l1: "United"})
	stack.exec(Op{type_: "MAKE_NOUN_ADJ", l2: "", l1: ""})
	stack.exec(Op{type_: "ADD_DET", l2: "el", l1: "the"})
	stack.exec(Op{type_: "MAKE_DET_NOUN", l2: "", l1: ""})
	stack.exec(Op{type_: "ADD_PREP", l2: "de", l1: "from"})
	stack.exec(Op{type_: "MAKE_PREP_NOUN", l2: "", l1: ""})
	stack.exec(Op{type_: "MAKE_NOUN_PHRASE_ADDING_PREP_PHRASE", l2: "", l1: ""})
	stack.exec(Op{type_: "MAKE_DOBJ", l2: "", l1: ""})
	stack.exec(Op{type_: "ADD_NOUN", l2: "millón", l1: "million"})
	stack.exec(Op{type_: "MAKE_PLURAL", l2: "-es", l1: "-s"})
	stack.exec(Op{type_: "ADD_NOUN", l2: "dólar", l1: "dollar"})
	stack.exec(Op{type_: "MAKE_PLURAL", l2: "-es", l1: "-s"})
	stack.exec(Op{type_: "ADD_PREP", l2: "de", l1: "of"})
	stack.exec(Op{type_: "MAKE_PREP_NOUN", l2: "", l1: ""})
	stack.exec(Op{type_: "MAKE_NOUN_PHRASE_ADDING_PREP_PHRASE", l2: "", l1: ""})
	stack.exec(Op{type_: "ADD_PREP", l2: "por", l1: "for"})
	stack.exec(Op{type_: "MAKE_PREP_NOUN", l2: "", l1: ""})
	stack.exec(Op{type_: "MAKE_VERB_PHRASE_ADDING_PREP_PHRASE", l2: "", l1: ""})
	stack.exec(Op{type_: "MAKE_COMPOUND_VERB", l2: "", l1: ""})

	fmt.Printf("%+v\n", stack)
	fmt.Printf("%+v\n", stack.getL1Words())
	fmt.Printf("%+v\n", stack.getL2Words())
}
