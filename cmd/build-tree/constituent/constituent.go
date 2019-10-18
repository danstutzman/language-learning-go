package constituent

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

func New(type_, l2, l1 string) Constituent {
	return Constituent{
		type_: type_,
		l2:    l2,
		l1:    l1,
	}
}

func (constituent Constituent) GetType() string {
	return constituent.type_
}
func (constituent Constituent) GetL1() string {
	return constituent.l1
}
func (constituent Constituent) GetL2() string {
	return constituent.l2
}

func (constituent Constituent) GetL1Words() []string {
	words := []string{}
	for _, child := range constituent.leftChildren {
		words = append(words, child.GetL1Words()...)
	}
	for _, l1Prefix := range constituent.l1Prefixes {
		words = append(words, l1Prefix)
	}
	words = append(words, constituent.l1)
	for _, l1Suffix := range constituent.l1Suffixes {
		words = append(words, l1Suffix)
	}
	for _, child := range constituent.rightChildren {
		words = append(words, child.GetL1Words()...)
	}
	return words
}

func (constituent Constituent) GetL2Words() []string {
	words := []string{}
	for _, child := range constituent.leftChildren {
		words = append(words, child.GetL2Words()...)
	}
	for _, l2Prefix := range constituent.l2Prefixes {
		words = append(words, l2Prefix)
	}
	words = append(words, constituent.l2)
	for _, l2Suffix := range constituent.l2Suffixes {
		words = append(words, l2Suffix)
	}
	for _, child := range constituent.rightChildren {
		words = append(words, child.GetL2Words()...)
	}
	return words
}

func (constituent *Constituent) SetLeftChild(newType string,
	newChild Constituent) {

	constituent.type_ = newType

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

func (constituent *Constituent) ChangeInto(
	newType, l2Prefix, l1Prefix, l2Suffix, l1Suffix string) {

	constituent.type_ = newType
	if l2Prefix != "" {
		constituent.prependL2Prefix(l2Prefix)
	}
	if l1Prefix != "" {
		constituent.prependL1Prefix(l1Prefix)
	}
	if l2Suffix != "" {
		constituent.appendL2Suffix(l2Suffix)
	}
	if l1Suffix != "" {
		constituent.appendL1Suffix(l1Suffix)
	}
}

func (constituent *Constituent) MakePhrase(newType string,
	newChild Constituent) {

	constituent.type_ = newType
	constituent.appendRightChild(newChild)
}
