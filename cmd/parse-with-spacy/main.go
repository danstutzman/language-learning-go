package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Token struct {
	Text  string `json:"text"`
	Lemma string `json:"lemma"`
	Pos   string `json:"pos"`
	Tag   string `json:"tag"`
	Dep   string `json:"dep"`
	Head  int    `json:"head"`
}

func parseWithSpacy(phrases []string, python3Path string) [][]Token {
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
		for _, phrase := range phrases {
			io.WriteString(stdin, phrase+"\n")
		}
	}()

	jsonLines, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			fmt.Fprintf(os.Stderr, "Stderr: %s\n", ee.Stderr)
		}
		panic(err)
	}

	parses := [][]Token{}
	for _, jsonLine := range bytes.Split(jsonLines, []byte{'\n'}) {
		if len(jsonLine) == 0 {
			continue
		}

		var tokenList []Token
		err = json.Unmarshal(jsonLine, &tokenList)
		if err != nil {
			panic(err)
		}

		parses = append(parses, tokenList)
	}
	return parses
}

func main() {
	if len(os.Args) != 1+1 {
		fmt.Fprintf(os.Stderr, `Usage:
			1st arg is path to python3 binary
		`)
		os.Exit(1)
	}
	python3Path := os.Args[1]

	// Gotta avoid 'llamas' since it's incorrectly tagged as a noun
	phrases := []string{
		"Me llamo Daniel.", // "I'm called Daniel."
		"Se llama Daniel.", // "He/she's called Daniel."
		"¿Cómo me llamo?",  // "What am I called?"
		"¿Cómo me llama?",  // "What does he/she call me?" (non-reflexive)
		"¿Cómo se llama?",  // "What is he/she called?"
	}

	tokensByPhraseNum := parseWithSpacy(phrases, python3Path)

	if false {
		for _, parse := range tokensByPhraseNum {
			for _, token := range parse {
				fmt.Printf("%+v\n", token)
			}
			fmt.Printf("\n")
		}
	}

	factsByPhraseNum := [][][]string{}
	for _, parse := range tokensByPhraseNum {
		facts := [][]string{}
		for i, token := range parse {
			iStr := strconv.Itoa(i)

			facts = append(facts, []string{"pos", iStr, token.Pos})
			facts = append(facts, []string{"lemma", iStr, token.Lemma})
			facts = append(facts,
				[]string{"head", iStr, strconv.Itoa(token.Head), token.Dep})

			if !strings.HasSuffix(token.Tag, "___") {
				part2 := strings.Split(token.Tag, "__")[1]
				for _, pair := range strings.Split(part2, "|") {
					parts := strings.Split(pair, "=")
					facts = append(facts, []string{"tag", iStr, parts[0], parts[1]})
				}
			}

		}
		factsByPhraseNum = append(factsByPhraseNum, facts)
	}

	queries := map[string][][]string{
		"SE LLAMAR PROPN": {
			{"pos", "?1", "VERB"},
			{"lemma", "?1", "llamar"},
			{"tag", "?1", "Person", "?4"},

			{"pos", "?2", "PRON"},
			{"head", "?2", "?1", "obj"},
			{"tag", "?2", "Person", "?4"},

			{"pos", "?3", "PROPN"},
			{"head", "?3", "?1", "nsubj"},
		},
		"Cómo SE LLAMAR": {
			{"pos", "?1", "VERB"},
			{"lemma", "?1", "llamar"},
			{"tag", "?1", "Person", "?4"},

			{"pos", "?2", "PRON"},
			{"head", "?2", "?1", "obj"},
			{"tag", "?2", "Person", "?4"},

			{"pos", "?3", "PRON"},
			{"lemma", "?3", "Cómo"},
			{"head", "?3", "?1", "obl"},
		},
		"LLAMAR": {
			{"pos", "?1", "VERB"},
			{"lemma", "?1", "llamar"},
		},
	}

	for queryName, query := range queries {
		fmt.Printf("%s:\n", queryName)
		for phraseNum, facts := range factsByPhraseNum {
			if variables := factsMatchQuery(facts, query); variables != nil {
				fmt.Printf("- %v\n", phrases[phraseNum])
			}
		}
	}
}

func factsMatchQuery(facts [][]string,
	query [][]string) map[string]string {
	possibleValues := gatherPossibleValues(facts, query)
	possibilities := cartesianJoin(possibleValues)
	for _, variables := range possibilities {
		if allQueryFactsMatch(query, facts, variables) {
			return variables
		}
	}
	return nil
}

func gatherPossibleValues(facts, query [][]string) map[string]map[string]bool {
	possibleValues := map[string]map[string]bool{}
	for _, queryFact := range query {
		for _, fact := range facts {
			//fmt.Printf("Considering fact %v\n", fact)
			if queryFact[0] == fact[0] {
				for i := 1; i < len(fact); i++ {
					queryFactArg := queryFact[i]
					if strings.HasPrefix(queryFactArg, "?") {
						set, ok := possibleValues[queryFactArg]
						if !ok {
							set = map[string]bool{}
							possibleValues[queryFactArg] = set
						}
						//fmt.Printf("Fact %v shows %s can be %s\n",
						//	fact, queryFactArg, fact[i])
						set[fact[i]] = true
					}
				}
			}
		}
	}
	return possibleValues
}

func cartesianJoin(
	possibleValues map[string]map[string]bool) []map[string]string {

	possibilities := []map[string]string{map[string]string{}}
	for variable, values := range possibleValues {
		newPossibilities := []map[string]string{}
		for newValue := range values {
			for _, oldPossibility := range possibilities {
				newPossibility := map[string]string{}
				for key, value := range oldPossibility {
					newPossibility[key] = value
				}
				newPossibility[variable] = newValue
				newPossibilities = append(newPossibilities, newPossibility)
			}
		}
		possibilities = newPossibilities
	}
	return possibilities
}

func allQueryFactsMatch(query, facts [][]string,
	variables map[string]string) bool {
	for _, queryFact := range query {
		//fmt.Printf("    QueryFact %v\n", queryFact)
		if !hasMatch(queryFact, facts, variables) {
			return false
		}
	}
	return true
}

func hasMatch(queryFact []string, facts [][]string,
	variables map[string]string) bool {
	for _, fact := range facts {
		if matches(queryFact, fact, variables) {
			//fmt.Printf("      QueryFact satisfied: %v\n", queryFact)
			return true
		}
	}
	return false
}

func matches(queryFact, fact []string, variables map[string]string) bool {
	if queryFact[0] == fact[0] {
		for i := 1; i < len(fact); i++ {
			queryFactArg := queryFact[i]

			var queryValue string
			if queryFactArg[0] == '?' {
				queryValue = variables[queryFactArg]
			} else {
				queryValue = queryFactArg
			}

			if queryValue != fact[i] {
				//fmt.Printf("      %v[%d] != %s\n", fact, i, queryValue)
				return false
			}
		}
		return true
	}
	return false
}
