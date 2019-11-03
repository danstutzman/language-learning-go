package spacy

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/freeling"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

type Token struct {
	Id       int    `json:"id"`
	Text     string `json:"text"`
	Lemma    string `json:"lemma"`
	Pos      string `json:"pos"`
	SpacyTag string `json:"spacy_tag"`
	Dep      string `json:"dep"`
	Head     int    `json:"head"`
	Idx      int    `json:"idx"`

	VerbTag  string            // Freeling tag, only for VERB or AUX
	Features map[string]string // Unpacked copy of SpacyTag's key-value pairs
}

var A1VerbTags = map[string]bool{
	"VMN0000": true, // infinitive
	"VMIP1S0": true, // 1st sing
	"VMIP2S0": true, // 2nd sing
	"VMIP3S0": true, // 3rd sing
	"VMIP1P0": true, // 1st plur
	"VMIP3P0": true, // 3rd plur
	"VMG0000": true, // gerund
}

func ParseWithSpacy(phrases []string, python3Path string) []string {
	fmt.Fprintf(os.Stderr, "Parsing with Spacy...")
	cmd := exec.Command(python3Path, "-c", `import json, spacy, sys
nlp = spacy.load('es_core_news_sm')
for line in sys.stdin:
	print(json.dumps([
  	{'id':        token.i,
		 'text':      token.text,
 		 'lemma':     token.lemma_,
		 'pos':       token.pos_,
		 'spacy_tag': token.tag_,
		 'dep':       token.dep_,
		 'head':      token.head.i,
		 'idx':       token.idx,
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
	fmt.Fprintf(os.Stderr, "completed.\n")

	parses := []string{}
	for _, jsonLine := range bytes.Split(jsonLines, []byte{'\n'}) {
		if len(jsonLine) > 0 {
			parses = append(parses, string(jsonLine))
		}
	}
	return parses
}

/*
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

	tokensByPhraseNum := ParseWithSpacy(phrases, python3Path)

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

			if !strings.HasSuffix(token.SpacyTag, "___") {
				part2 := strings.Split(token.SpacyTag, "__")[1]
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

		for _, phrase := range generateVerbPhrases(queryName, query) {
			fmt.Printf("+ %s\n", phrase)
		}
	}
}
*/

func findHeadOfQuery(queryName string, query [][]string) string {
	possibleHeads := map[string]bool{}
	for _, fact := range query {
		if fact[0] == "head" {
			possibleHeads[fact[2]] = true
		}
	}

	if len(possibleHeads) > 0 {
		// Exclude possible heads that are a child
		for _, fact := range query {
			if fact[0] == "head" {
				delete(possibleHeads, fact[1])
			}
		}
	} else {
		for _, fact := range query {
			switch fact[0] {
			case "pos", "lemma", "tag":
				possibleHeads[fact[1]] = true
			}
		}
	}

	if len(possibleHeads) > 1 {
		panic(fmt.Errorf("Too many possibleHeads of queryName=%s", queryName))
	}
	for possibleHead := range possibleHeads {
		return possibleHead
	}
	panic(fmt.Errorf("Can't find head of queryName=%s", queryName))
}

func assertHasFact(queryName string, haystack [][]string, needle []string) {
	for _, fact := range haystack {
		matches := true
		for i := 0; i < len(needle); i++ {
			if fact[i] != needle[i] {
				matches = false
				break
			}
		}
		if matches {
			return
		}
	}
	panic(fmt.Errorf("Can't find %v in %v", needle, haystack))
}

func generateVerbPhrases(queryName string, query [][]string) []string {
	head := findHeadOfQuery(queryName, query)
	assertHasFact(queryName, query, []string{"pos", head, "VERB"})

	verbTags := []VerbTag{}
	for _, verbTag := range allVerbTags {
		if A1VerbTags[verbTag.TagPair.FreelingTag] {
			verbTags = append(verbTags, verbTag)
		}
	}

	verbs := []string{}
	for _, verbTag := range verbTags {
		conjugations := freeling.AnalyzeVerb("llamar", verbTag.TagPair.FreelingTag)
		for _, conjugation := range conjugations {
			verb := conjugation.Stem + conjugation.Suffix
			verbs = append(verbs, verb)
		}
	}
	return verbs
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

func Uncapitalize1stLetter(s string) string {
	runes := []rune(s)
	if runes[0] == '¿' {
		runes[1] = unicode.To(unicode.LowerCase, runes[1])
	} else {
		runes[0] = unicode.To(unicode.LowerCase, runes[0])
	}
	return string(runes)
}
