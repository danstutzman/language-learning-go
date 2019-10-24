package commands

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/commands/constituent"
	"log"
	"strings"
)

type Stack struct {
	constituents []constituent.Constituent
}

func expectNumArgs(expectedNum int, args []string) {
	if len(args) != expectedNum {
		log.Fatalf("Expected %d args but got %v", expectedNum, args)
	}
}

func NewStack() Stack {
	return Stack{}
}

func (stack *Stack) ExecCommand(commandWithArgs string) {
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
	case "MAKE_DET_NOUN_PHRASE":
		expectNumArgs(0, args)
		stack.makeDetNounPhrase()
	case "MAKE_INFINITIVE":
		expectNumArgs(2, args)
		stack.makeInfinitive(args[0], args[1])
	case "MAKE_NOUN_PHRASE_ADJ":
		expectNumArgs(0, args)
		stack.makeNounPhraseAdj()
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
	case "MAKE_VOBJ":
		expectNumArgs(0, args)
		stack.makeVerbObj()
	case "MAKE_VERB_PHRASE_ADDING_PREP_PHRASE":
		expectNumArgs(0, args)
		stack.makeVerbPhraseAddingPrepPhrase()
	case "ATTACH_ATR_TO_VP":
		expectNumArgs(0, args)
		stack.makeVerbPhraseAddingAdj()
	default:
		panic("Unknown command " + command)
	}
}

func (stack *Stack) push(type_, l2, l1 string) {
	stack.constituents = append(stack.constituents,
		constituent.New(type_, l2, l1))
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
	agent := stack.pop(EXPECT_NOUN_OR_PHRASE)
	verbPhrase := stack.peek(EXPECT_VERB_OR_PHRASE)
	verbPhrase.SetLeftChild("VERB_PHRASE", agent)
}

func (stack *Stack) makeVerbObj() {
	verbPhraseToAdd := stack.pop(EXPECT_VERB_OR_PHRASE)
	verbToGrow := stack.peek(EXPECT_VERB_OR_PHRASE)
	verbToGrow.MakePhraseAppendingChild("VERB_PHRASE", verbPhraseToAdd)
}

func (stack *Stack) makeDetNounPhrase() {
	det := stack.pop(EXPECT_DET)
	noun := stack.peek(EXPECT_NOUN_OR_PHRASE)
	noun.ChangeInto("NOUN_PHRASE", det.GetL2(), det.GetL1(), "", "")
}

func (stack *Stack) makeInfinitive(l2, l1 string) {
	stem := stack.peek(EXPECT_VERB_STEM)
	stem.ChangeInto("VERB", "", "to", "-ar", "")
}

func (stack *Stack) makeNounPhraseAdj() {
	adj := stack.pop(EXPECT_ADJ)
	noun := stack.peek(EXPECT_NOUN_OR_PHRASE)
	noun.MakePhraseAppendingL2PrependingL1("NOUN_PHRASE", adj)
}

func (stack *Stack) makeDirObj() {
	nounPhrase := stack.pop(EXPECT_NOUN_OR_PHRASE)
	verbPhrase := stack.peek(EXPECT_VERB_OR_PHRASE)
	verbPhrase.MakePhraseAppendingChild("VERB_PHRASE", nounPhrase)
}

func (stack *Stack) makeNounPhraseAddingPrepPhrase() {
	prepPhrase := stack.pop(EXPECT_PREP_PHRASE)
	nounPhrase := stack.peek(EXPECT_NOUN_OR_PHRASE)
	nounPhrase.MakePhraseAppendingChild("NOUN_PHRASE", prepPhrase)
}

func (stack *Stack) makeVerbPhraseAddingPrepPhrase() {
	prepPhrase := stack.pop(EXPECT_PREP_PHRASE)
	verbPhrase := stack.peek(EXPECT_VERB_OR_PHRASE)
	verbPhrase.MakePhraseAppendingChild("VERB_PHRASE", prepPhrase)
}

func (stack *Stack) makeVerbPhraseAddingAdj() {
	adj := stack.pop(EXPECT_ADJ)
	verbPhrase := stack.peek(EXPECT_VERB_OR_PHRASE)
	verbPhrase.MakePhraseAppendingChild("VERB_PHRASE", adj)
}

func (stack *Stack) makePlural(l2, l1 string) {
	noun := stack.peek(EXPECT_NOUN)
	noun.ChangeInto("NOUN", "", "", l2, l1)
}

func (stack *Stack) makePrepNoun() {
	prep := stack.pop(EXPECT_PREP)
	noun := stack.peek(EXPECT_NOUN_OR_PHRASE)
	noun.ChangeInto("PREP_PHRASE", prep.GetL2(), prep.GetL1(), "", "")
}

func (stack *Stack) makePresProg(l2Suffix, l1Suffix string) {
	estar := stack.pop(EXPECT_VERB_UNIQUE)
	stem := stack.peek(EXPECT_VERB_STEM)
	stem.ChangeInto("VERB", estar.GetL2(), estar.GetL1(), l2Suffix, l1Suffix)
}

func (stack *Stack) peek(
	expectedTypes map[string]bool) *constituent.Constituent {
	constituent := stack.constituents[len(stack.constituents)-1]

	if expectedTypes[constituent.GetType()] == false {
		log.Panicf("Expected types in %v but got type=%s",
			expectedTypes, constituent.GetType())
	}

	return &stack.constituents[len(stack.constituents)-1]
}

func (stack *Stack) pop(
	expectedTypes map[string]bool) constituent.Constituent {
	lastConstituent := stack.peek(expectedTypes)
	stack.constituents = stack.constituents[0 : len(stack.constituents)-1]
	return *lastConstituent
}

func (stack *Stack) GetL1Words() []string {
	l1Words := []string{}
	for _, constituent := range stack.constituents {
		if len(l1Words) > 0 {
			l1Words = append(l1Words, "/")
		}
		l1Words = append(l1Words, constituent.GetL1Words()...)
	}
	return l1Words
}

func (stack *Stack) GetL2Words() []string {
	l2Words := []string{}
	for _, constituent := range stack.constituents {
		if len(l2Words) > 0 {
			l2Words = append(l2Words, "/")
		}
		l2Words = append(l2Words, constituent.GetL2Words()...)
	}
	return l2Words
}
