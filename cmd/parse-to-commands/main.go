package main

import (
	"fmt"
)

type Word struct {
	index       int
	form        string
	tag         string
	dep         string
	parentIndex int
}

var words = []Word{
	Word{0, "Apple", "PROPN", "nsubj", 2},
	Word{1, "está", "AUX", "aux", 2},
	Word{2, "buscando", "VERB", "ROOT", 2},
	Word{3, "comprar", "VERB", "xcomp", 2},
	Word{4, "una", "DET", "det", 5},
	Word{5, "startup", "NOUN", "obj", 3},
	Word{6, "del", "ADP", "case", 7},
	Word{7, "Reino", "PROPN", "nmod", 5},
	Word{8, "Unido", "PROPN", "flat", 7},
	Word{9, "por", "ADP", "case", 11},
	Word{10, "mil", "NUM", "nummod", 11},
	Word{11, "millones", "NOUN", "obl", 3},
	Word{12, "de", "ADP", "case", 13},
	Word{13, "dólares", "NOUN", "nmod", 11},
}

func main() {
	fmt.Println(words)
}
