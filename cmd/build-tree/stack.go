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
	case "MAKE_DET_NOUN_PHRASE":
		expectNumArgs(0, args)
		stack.makeDetNounPhrase()
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

var EXPECT_ADJ = map[string]bool{"ADJ": true}
var EXPECT_DET = map[string]bool{"DET": true}
var EXPECT_NOUN = map[string]bool{"NOUN": true}
var EXPECT_NOUN_OR_PHRASE = map[string]bool{"NOUN": true, "NOUN_PHRASE": true}
var EXPECT_PREP = map[string]bool{"PREP": true}
var EXPECT_PREP_PHRASE = map[string]bool{"PREP_PHRASE": true}
var EXPECT_VERB = map[string]bool{"VERB": true}
var EXPECT_VERB_OR_PHRASE = map[string]bool{"VERB": true, "VERB_PHRASE": true}
var EXPECT_VERB_STEM = map[string]bool{"VERB_STEM": true}
var EXPECT_VERB_UNIQUE = map[string]bool{"VERB_UNIQUE": true}

func (stack *Stack) makeAgent() {
	agent := stack.pop(EXPECT_NOUN)
	verbPhrase := stack.peek(EXPECT_VERB_OR_PHRASE)
	verbPhrase.setLeftChild(agent)
}

func (stack *Stack) makeCompoundVerb() {
	verbPhraseToAdd := stack.pop(EXPECT_VERB_OR_PHRASE)
	verbToGrow := stack.peek(EXPECT_VERB)
	verbToGrow.type_ = "VERB_PHRASE"
	verbToGrow.appendRightChild(verbPhraseToAdd)
}

func (stack *Stack) makeDetNounPhrase() {
	det := stack.pop(EXPECT_DET)
	noun := stack.peek(EXPECT_NOUN_OR_PHRASE)
	noun.type_ = "NOUN_PHRASE"
	noun.prependL2Prefix(det.l2)
	noun.prependL1Prefix(det.l1)
}

func (stack *Stack) makeInfinitive(l2, l1 string) {
	stem := stack.peek(EXPECT_VERB_STEM)
	stem.type_ = "VERB"
	stem.appendL2Suffix("-ar")
	stem.prependL1Prefix("to")
}

func (stack *Stack) makeNounAdj() {
	adj := stack.pop(EXPECT_ADJ)
	noun := stack.peek(EXPECT_NOUN)
	noun.type_ = "NOUN_PHRASE"
	noun.appendRightChild(adj)
}

func (stack *Stack) makeDirObj() {
	nounPhrase := stack.pop(EXPECT_NOUN_OR_PHRASE)
	verbPhrase := stack.peek(EXPECT_VERB_OR_PHRASE)
	verbPhrase.appendRightChild(nounPhrase)
}

func (stack *Stack) makeNounPhraseAddingPrepPhrase() {
	prepPhrase := stack.pop(EXPECT_PREP_PHRASE)
	nounPhrase := stack.peek(EXPECT_NOUN_OR_PHRASE)
	nounPhrase.appendRightChild(prepPhrase)
}

func (stack *Stack) makeVerbPhraseAddingPrepPhrase() {
	prepPhrase := stack.pop(EXPECT_PREP_PHRASE)
	verbPhrase := stack.peek(EXPECT_VERB_OR_PHRASE)
	verbPhrase.appendRightChild(prepPhrase)
}

func (stack *Stack) makePlural(l2, l1 string) {
	noun := stack.peek(EXPECT_NOUN)
	noun.appendL2Suffix(l2)
	noun.appendL1Suffix(l1)
}

func (stack *Stack) makePrepNoun() {
	prep := stack.pop(EXPECT_PREP)
	noun := stack.peek(EXPECT_NOUN_OR_PHRASE)
	noun.type_ = "PREP_PHRASE"
	noun.prependL2Prefix(prep.l2)
	noun.prependL1Prefix(prep.l1)
}

func (stack *Stack) makePresProg(l2Suffix, l1Suffix string) {
	estar := stack.pop(EXPECT_VERB_UNIQUE)
	stem := stack.peek(EXPECT_VERB_STEM)
	stem.type_ = "VERB"
	stem.prependL2Prefix(estar.l2)
	stem.prependL1Prefix(estar.l1)
	stem.appendL2Suffix(l2Suffix)
	stem.appendL1Suffix(l1Suffix)
}

func (stack *Stack) peek(expectedTypes map[string]bool) *Constituent {
	constituent := stack.constituents[len(stack.constituents)-1]

	if expectedTypes[constituent.type_] == false {
		log.Panicf("Expected types in %v but got type=%s",
			expectedTypes, constituent.type_)
	}

	return &stack.constituents[len(stack.constituents)-1]
}

func (stack *Stack) pop(expectedTypes map[string]bool) Constituent {
	lastConstituent := stack.peek(expectedTypes)
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
