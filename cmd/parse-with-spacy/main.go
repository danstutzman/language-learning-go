package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Token struct {
	Text  string `json:"text"`
	Lemma string `json:"lemma"`
	Pos   string `json:"pos"`
	Tag   string `json:"tag"`
	Dep   string `json:"dep"`
	Head  int    `json:"head"`
}

func main() {
	if len(os.Args) != 1+1 {
		fmt.Fprintf(os.Stderr, `Usage:
			1st arg is path to python3 binary
		`)
		os.Exit(1)
	}
	python3Path := os.Args[1]

	cmd := exec.Command(python3Path, "-c", `import json, spacy, sys
nlp = spacy.load('es_core_news_sm')
for line in sys.stdin:
	print(json.dumps([
		{'text':  token.text,
 		 'lemma': token.lemma_,
		 'pos':   token.pos_,
		 'tag':   token.tag_,
		 'dep':   token.dep_,
		 'head':  token.head.i,
		 'idx':   token.idx,
	  } for token in nlp(line.rstrip())]))`)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "me llamo\nte llamas\n")
	}()

	jsonLines, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			fmt.Fprintf(os.Stderr, "Stderr: %s\n", ee.Stderr)
		}
		panic(err)
	}

	for _, jsonLine := range bytes.Split(jsonLines, []byte{'\n'}) {
		if len(jsonLine) == 0 {
			continue
		}

		var tokens []Token
		err = json.Unmarshal(jsonLine, &tokens)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", tokens)
	}
}
