package spacy

import ()

type Dep struct {
	Token    Token
	Function string
	Children []Dep
}

func TokensToDeps(tokens []Token) []Dep {
	rootTokenNums := []int{}
	for tokenNum, token := range tokens {
		if token.Head == tokenNum {
			rootTokenNums = append(rootTokenNums, tokenNum)
		}
	}
	if len(rootTokenNums) == 0 {
		panic("Couldn't find root!")
	}

	deps := []Dep{}
	for _, rootTokenNum := range rootTokenNums {
		dep := tokenToDep(rootTokenNum, tokens)
		deps = append(deps, dep)
	}
	return deps
}

func tokenToDep(tokenNum int, tokens []Token) Dep {
	children := []Dep{}
	for childTokenNum, childToken := range tokens {
		if childToken.Head == tokenNum && childTokenNum != tokenNum {
			children = append(children, tokenToDep(childTokenNum, tokens))
		}
	}

	return Dep{
		Token:    tokens[tokenNum],
		Function: tokens[tokenNum].Dep,
		Children: children,
	}
}
