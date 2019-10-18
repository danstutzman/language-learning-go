package main

type Constituent struct {
	l1            string
	l2            string
	leftChildren  []Constituent
	rightChildren []Constituent
}

func (constituent Constituent) getL1Words() []string {
	words := []string{}
	for _, child := range constituent.leftChildren {
		words = append(words, child.getL1Words()...)
	}
	words = append(words, constituent.l1)
	for _, child := range constituent.rightChildren {
		words = append(words, child.getL1Words()...)
	}
	return words
}

func (constituent Constituent) getL2Words() []string {
	words := []string{}
	for _, child := range constituent.leftChildren {
		words = append(words, child.getL2Words()...)
	}
	words = append(words, constituent.l2)
	for _, child := range constituent.rightChildren {
		words = append(words, child.getL2Words()...)
	}
	return words
}

func (constituent *Constituent) setLeftChild(newChild Constituent) {
	if len(constituent.leftChildren) != 0 {
		panic("Left child already set")
	}
	constituent.leftChildren = []Constituent{newChild}
}

func (constituent *Constituent) appendRightChild(newChild Constituent) {
	constituent.rightChildren = append(constituent.rightChildren, newChild)
}
