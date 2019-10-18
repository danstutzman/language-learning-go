package main

type Stack struct {
	constituents []Constituent
}

func (stack *Stack) exec(op Op) {
	if op.type_ == "ADD_VERB_INFINITIVE" ||
		op.type_ == "ADD_VERB_UNIQUE" ||
		op.type_ == "ADD_NOUN" ||
		op.type_ == "ADD_DET" ||
		op.type_ == "ADD_ADJ" ||
		op.type_ == "ADD_PREP" {
		stack.push(Constituent{l1: op.l1, l2: op.l2})
		return
	} else if op.type_ == "MAKE_PRES_PROG" {
		estar := stack.pop()
		infinitive := stack.peek()
		infinitive.prependL2Prefix(estar.l2)
		infinitive.prependL1Prefix(estar.l1)
		infinitive.appendL2Suffix(op.l2)
		infinitive.appendL1Suffix(op.l1)
		return
	} else if op.type_ == "MAKE_AGENT" {
		agent := stack.pop()
		lastConstituent := stack.peek()
		lastConstituent.leftChildren =
			append([]Constituent{agent}, lastConstituent.leftChildren...)
		return
	} else if op.type_ == "MAKE_DET_NOUN" {
		det := stack.pop()
		noun := stack.peek()
		noun.prependL2Prefix(det.l2)
		noun.prependL1Prefix(det.l1)
		return
	} else if op.type_ == "MAKE_NOUN_ADJ" {
		adj := stack.pop()
		noun := stack.peek()
		noun.appendRightChild(adj)
		return
	} else if op.type_ == "MAKE_PREP_NOUN" {
		prep := stack.pop()
		noun := stack.peek()
		noun.prependL2Prefix(prep.l2)
		noun.prependL1Prefix(prep.l1)
		return
	} else if op.type_ == "MAKE_NOUN_PHRASE_ADDING_PREP_PHRASE" ||
		op.type_ == "MAKE_VERB_PHRASE_ADDING_PREP_PHRASE" {
		prepPhrase := stack.pop()
		growingPhrase := stack.peek()
		growingPhrase.appendRightChild(prepPhrase)
		return
	} else if op.type_ == "MAKE_DOBJ" {
		nounPhrase := stack.pop()
		verbPhrase := stack.peek()
		verbPhrase.appendRightChild(nounPhrase)
		return
	} else if op.type_ == "MAKE_PLURAL" {
		noun := stack.peek()
		noun.appendL2Suffix("-es")
		noun.appendL1Suffix("-s")
		return
	} else if op.type_ == "MAKE_INFINITIVE" {
		noun := stack.peek()
		noun.appendL2Suffix("-ar")
		noun.prependL1Prefix("to")
		return
	} else if op.type_ == "MAKE_COMPOUND_VERB" {
		verbPhraseToAdd := stack.pop()
		verbPhraseToGrow := stack.peek()
		verbPhraseToGrow.appendRightChild(verbPhraseToAdd)
		return
	} else {
		panic("Unknown op.type " + op.type_)
	}
}

func (stack *Stack) push(newConstituent Constituent) {
	stack.constituents = append(stack.constituents, newConstituent)
}

func (stack *Stack) peek() *Constituent {
	return &stack.constituents[len(stack.constituents)-1]
}

func (stack *Stack) pop() Constituent {
	lastConstituent := stack.peek()
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
