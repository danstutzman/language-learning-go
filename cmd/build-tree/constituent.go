package main

type Constituent struct {
	type_         string
	l2Prefixes    []string
	l1Prefixes    []string
	l2            string
	l1            string
	l2Suffixes    []string
	l1Suffixes    []string
	leftChildren  []Constituent
	rightChildren []Constituent
}

func (constituent Constituent) getL1Words() []string {
	words := []string{}
	for _, child := range constituent.leftChildren {
		words = append(words, child.getL1Words()...)
	}
	for _, l1Prefix := range constituent.l1Prefixes {
		words = append(words, l1Prefix)
	}
	words = append(words, constituent.l1)
	for _, l1Suffix := range constituent.l1Suffixes {
		words = append(words, l1Suffix)
	}
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
	for _, l2Prefix := range constituent.l2Prefixes {
		words = append(words, l2Prefix)
	}
	words = append(words, constituent.l2)
	for _, l2Suffix := range constituent.l2Suffixes {
		words = append(words, l2Suffix)
	}
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

func (constituent *Constituent) prependL2Prefix(l2Prefix string) {
	constituent.l2Prefixes = append([]string{l2Prefix}, constituent.l2Prefixes...)
}

func (constituent *Constituent) prependL1Prefix(l1Prefix string) {
	constituent.l1Prefixes = append([]string{l1Prefix}, constituent.l1Prefixes...)
}

func (constituent *Constituent) appendL2Suffix(l2Suffix string) {
	constituent.l2Suffixes = append([]string{l2Suffix}, constituent.l2Suffixes...)
}

func (constituent *Constituent) appendL1Suffix(l1Suffix string) {
	constituent.l1Suffixes = append([]string{l1Suffix}, constituent.l1Suffixes...)
}

func (constituent *Constituent) appendRightChild(newChild Constituent) {
	constituent.rightChildren = append(constituent.rightChildren, newChild)
}
