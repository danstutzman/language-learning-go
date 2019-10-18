package main

import (
	"log"
	"strings"
)

type Stack struct {
	constituents []Constituent
}

func expectNumArgs(expectedNum int, args []string) {
	if len(args) != expectedNum {
		log.Fatalf("Expected %d args but got %v", expectedNum, args)
	}
}

func (stack *Stack) execCommand(commandWithArgs string) {
	args := strings.Split(commandWithArgs, "/")
	command := args[0]
	args = args[1:len(args)]

	switch command {
	case "ADD":
		expectNumArgs(3, args)
		stack.push(args[0], args[1], args[2])
	case "MAKE_AGENT":
		expectNumArgs(0, args)
		stack.makeAgent()
	case "MAKE_COMPOUND_VERB":
		expectNumArgs(0, args)
		stack.makeCompoundVerb()
	case "MAKE_DET_NOUN":
		expectNumArgs(0, args)
		stack.makeDetNoun()
	case "MAKE_INFINITIVE":
		expectNumArgs(2, args)
		stack.makeInfinitive(args[0], args[1])
	case "MAKE_NOUN_ADJ":
		expectNumArgs(0, args)
		stack.makeNounAdj()
	case "MAKE_PLURAL":
		expectNumArgs(2, args)
		stack.makePlural(args[0], args[1])
	case "MAKE_PREP_NOUN":
		expectNumArgs(0, args)
		stack.makePrepNoun()
	case "MAKE_PRES_PROG":
		expectNumArgs(2, args)
		stack.makePresProg(args[0], args[1])
	case "MAKE_DOBJ":
		expectNumArgs(0, args)
		stack.makeDirObj()
	case "MAKE_NOUN_PHRASE_ADDING_PREP_PHRASE":
		expectNumArgs(0, args)
		stack.makeNounPhraseAddingPrepPhrase()
	case "MAKE_VERB_PHRASE_ADDING_PREP_PHRASE":
		expectNumArgs(0, args)
		stack.makeVerbPhraseAddingPrepPhrase()
	default:
		panic("Unknown command " + command)
	}
}

func (stack *Stack) push(type_, l2, l1 string) {
	stack.constituents = append(stack.constituents,
		Constituent{type_: type_, l2: l2, l1: l1})
}

func (stack *Stack) makeAgent() {
	agent := stack.pop("NOUN")
	verbPhrase := stack.peek("VERB")
	verbPhrase.setLeftChild(agent)
}

func (stack *Stack) makeCompoundVerb() {
	verbPhraseToAdd := stack.pop("VERB")
	verbPhraseToGrow := stack.peek("VERB")
	verbPhraseToGrow.appendRightChild(verbPhraseToAdd)
}

func (stack *Stack) makeDetNoun() {
	det := stack.pop("DET")
	noun := stack.peek("NOUN")
	noun.prependL2Prefix(det.l2)
	noun.prependL1Prefix(det.l1)
}

func (stack *Stack) makeInfinitive(l2, l1 string) {
	stem := stack.peek("VERB_STEM")
	stem.type_ = "VERB"
	stem.appendL2Suffix("-ar")
	stem.prependL1Prefix("to")
}

func (stack *Stack) makeNounAdj() {
	adj := stack.pop("ADJ")
	noun := stack.peek("NOUN")
	noun.appendRightChild(adj)
}

func (stack *Stack) makeDirObj() {
	nounPhrase := stack.pop("NOUN")
	verbPhrase := stack.peek("VERB")
	verbPhrase.appendRightChild(nounPhrase)
}

func (stack *Stack) makeNounPhraseAddingPrepPhrase() {
	prepPhrase := stack.pop("PREP")
	nounPhrase := stack.peek("NOUN")
	nounPhrase.appendRightChild(prepPhrase)
}

func (stack *Stack) makeVerbPhraseAddingPrepPhrase() {
	prepPhrase := stack.pop("PREP")
	verbPhrase := stack.peek("VERB")
	verbPhrase.appendRightChild(prepPhrase)
}

func (stack *Stack) makePlural(l2, l1 string) {
	noun := stack.peek("NOUN")
	noun.appendL2Suffix(l2)
	noun.appendL1Suffix(l1)
}

func (stack *Stack) makePrepNoun() {
	prep := stack.pop("PREP")
	noun := stack.peek("NOUN")
	noun.type_ = "PREP"
	noun.prependL2Prefix(prep.l2)
	noun.prependL1Prefix(prep.l1)
}

func (stack *Stack) makePresProg(l2Suffix, l1Suffix string) {
	estar := stack.pop("VERB_UNIQUE")
	stem := stack.peek("VERB_STEM")
	stem.type_ = "VERB"
	stem.prependL2Prefix(estar.l2)
	stem.prependL1Prefix(estar.l1)
	stem.appendL2Suffix(l2Suffix)
	stem.appendL1Suffix(l1Suffix)
}

func (stack *Stack) peek(expectedType string) *Constituent {
	constituent := stack.constituents[len(stack.constituents)-1]

	if constituent.type_ != expectedType {
		panic("Expected type=" + expectedType +
			" but got type=" + constituent.type_)
	}

	return &stack.constituents[len(stack.constituents)-1]
}

func (stack *Stack) pop(expectedType string) Constituent {
	lastConstituent := stack.peek(expectedType)
	stack.constituents = stack.constituents[0 : len(stack.constituents)-1]
	return *lastConstituent
}

func (stack *Stack) getL1Words() []string {
	l1Words := []string{}
	for _, constituent := range stack.constituents {
		if len(l1Words) > 0 {
			l1Words = append(l1Words, "/")
		}
		l1Words = append(l1Words, constituent.getL1Words()...)
	}
	return l1Words
}

func (stack *Stack) getL2Words() []string {
	l2Words := []string{}
	for _, constituent := range stack.constituents {
		if len(l2Words) > 0 {
			l2Words = append(l2Words, "/")
		}
		l2Words = append(l2Words, constituent.getL2Words()...)
	}
	return l2Words
}

func (stack *Stack) pushConstituent(l2, l1 string) {
	stack.constituents =
		append(stack.constituents, Constituent{l2: l2, l1: l1})
}
