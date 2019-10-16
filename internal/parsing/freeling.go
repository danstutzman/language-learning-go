package parsing

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

const BATCH_SIZE = 1

type Output struct {
	Phrase    string
	ParseJson string
	Parse     Parse
}

type Parse struct {
	Sentences []Sentence `json:"sentences"`
}

type Sentence struct {
	Id           string        `json:"id"`
	Tokens       []Token       `json:"tokens"`
	Constituents []Constituent `json:"constituents"`
	Dependencies []Dependency  `json:"dependencies"`
}

type Token struct {
	Id     string `json:"id"`
	Begin  string `json:"begin"`
	End    string `json:"end"`
	Form   string `json:"form"`
	Lemma  string `json:"lemma"`
	Tag    string `json:"tag"`
	Ctag   string `json:"ctag"`
	Pos    string `json:"pos"`
	Type   string `json:"type"`
	Mood   string `json:"mood"`
	Tense  string `json:"tense"`
	Person string `json:"person"`
	Num    string `json:"num"`
	Gen    string `json:"gen"`
}

func (token *Token) IsAdjective() bool {
	return strings.HasPrefix(token.Tag, "A")
}

func (token *Token) IsAdverb() bool {
	return strings.HasPrefix(token.Tag, "R")
}

func (token *Token) IsConjunction() bool {
	return strings.HasPrefix(token.Tag, "C")
}

func (token *Token) IsDate() bool {
	return strings.HasPrefix(token.Tag, "W")
}

func (token *Token) IsDeterminer() bool {
	return strings.HasPrefix(token.Tag, "D")
}

func (token *Token) IsPronoun() bool {
	return strings.HasPrefix(token.Tag, "P")
}

func (token *Token) IsInterjection() bool {
	return strings.HasPrefix(token.Tag, "I")
}

func (token *Token) IsPreposition() bool {
	return strings.HasPrefix(token.Tag, "SP")
}

func (token *Token) IsProperNoun() bool {
	return strings.HasPrefix(token.Tag, "NP")
}

func (token *Token) IsNoun() bool {
	return strings.HasPrefix(token.Tag, "N")
}

func (token *Token) IsNumber() bool {
	return strings.HasPrefix(token.Tag, "Z")
}

func (token *Token) IsPunctuation() bool {
	return strings.HasPrefix(token.Tag, "F")
}

func (token *Token) IsVerb() bool {
	return strings.HasPrefix(token.Tag, "V")
}

type Constituent struct {
	Label    string        `json:"label"`
	Children []Constituent `json:"children"`
	Leaf     string        `json:"leaf"`
	Head     string        `json:"head"`
	Token    string        `json:"token"`
	Word     string        `json:"word"`
}

type Dependency struct {
	Token    string       `json:"token"`
	Function string       `json:"function"`
	Word     string       `json:"word"`
	Children []Dependency `json:"children"`
}

func unmarshalParseJson(parseJson string) Parse {
	var parse Parse
	err := json.Unmarshal([]byte(parseJson), &parse)
	if err != nil {
		panic(err)
	}
	return parse
}

/* Example parse JSON of "Cómo está?" follows:
{ "sentences" : [
  { "id":"1",
    "tokens" : [
       { "id" : "t1.1", "begin" : "0", "end" : "1", "form" : "¿",
         "lemma" : "¿", "tag" : "Fia", "ctag" : "Fia", "pos" : "punctuation",
         "type" : "questionmark", "punctenclose" : "open"},
       { "id" : "t1.2", "begin" : "1", "end" : "5", "form" : "Cómo",
         "lemma" : "cómo", "tag" : "PT00000", "ctag" : "PT", "pos" : "pronoun",
          "type" : "interrogative"},
       { "id" : "t1.3", "begin" : "6", "end" : "10", "form" : "está",
         "lemma" : "estar", "tag" : "VMIP3S0", "ctag" : "VMI", "pos" : "verb",
         "type" : "main", "mood" : "indicative", "tense" : "present",
         "person" : "3", "num" : "singular"},
       { "id" : "t1.4", "begin" : "10", "end" : "11", "form" : "?",
         "lemma" : "?", "tag" : "Fit", "ctag" : "Fit", "pos" : "punctuation",
         "type" : "questionmark", "punctenclose" : "close"}],
    "constituents" : [
      {"label" : "grup-verb", "head" : "1", "children" : [
        {"label" : "F-no-c", "children" : [
          {"leaf" : "1", "head" : "1", "token" : "t1.1", "word" : "¿"}
        ]},
        {"label" : "sadv", "children" : [
          {"label" : "adv-interrog", "head" : "1", "children" : [
            {"leaf" : "1", "head" : "1", "token" : "t1.2", "word" : "Cómo"}
          ]}
        ]},
        {"label" : "verb", "head" : "1", "children" : [
          {"leaf" : "1", "head" : "1", "token" : "t1.3", "word" : "está"}
        ]},
        {"label" : "F-term", "children" : [
          {"leaf" : "1", "head" : "1", "token" : "t1.4", "word" : "?"}
        ]}
      ]}],
    "dependencies" : [
      {"token" : "t1.3", "function" : "top", "word" : "está", "children" : [
        {"token" : "t1.1", "function" : "punc", "word" : "¿"},
        {"token" : "t1.2", "function" : "adjt", "word" : "Cómo"},
        {"token" : "t1.4", "function" : "punc", "word" : "?"}
      ]}]}]}
]*/

func ParsePhrasesWithFreeling(phrases []string,
	freelingHostAndPort string) []Output {

	log.Printf("Conecting to %s\n", freelingHostAndPort)
	conn, err := net.Dial("tcp", freelingHostAndPort)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	log.Printf("Writing RESET_STATS...\n")
	_, err = conn.Write([]byte("RESET_STATS\x00"))
	if err != nil {
		panic(err)
	}

	log.Printf("Reading...\n")
	serverReady, err := reader.ReadString('\x00')
	if err != nil {
		panic(err)
	}
	if serverReady != "FL-SERVER-READY\x00" {
		panic("Server not ready?")
	}

	outputs := []Output{}
	for _, phrase := range phrases {
		if strings.ContainsRune(phrase, '\x00') {
			log.Panicf("Phrase contains \\x00: '%s'", phrase)
		}
		if strings.ContainsRune(phrase, '\n') {
			log.Panicf("Phrase contains newline: '%s'", phrase)
		}

		if len(phrase) > 120 {
			outputs = append(outputs, Output{})
			continue
		}

		log.Printf("Writing...\n")
		_, err := conn.Write([]byte(phrase + "\x00"))
		if err != nil {
			panic(err)
		}

		log.Printf("Reading...\n")
		output, err := reader.ReadString('\x00')
		if err != nil {
			panic(err)
		}
		if output != "FL-SERVER-READY\x00" {
			parseJson := strings.TrimSuffix(output, "\x00")
			parse := unmarshalParseJson(parseJson)
			output := Output{
				Phrase:    phrase,
				ParseJson: parseJson,
				Parse:     parse,
			}
			outputs = append(outputs, output)
		}

		log.Printf("Writing FLUSH_BUFFER...\n")
		_, err = conn.Write([]byte("FLUSH_BUFFER\x00"))
		if err != nil {
			panic(err)
		}

		log.Printf("Reading...\n")
		output, err = reader.ReadString('\x00')
		if err != nil {
			panic(err)
		}
		if output != "FL-SERVER-READY\x00" {
			parseJson := strings.TrimSuffix(output, "\x00")
			parse := unmarshalParseJson(parseJson)
			output := Output{
				Phrase:    phrase,
				ParseJson: parseJson,
				Parse:     parse,
			}
			outputs = append(outputs, output)
		}
	}

	if len(outputs) != len(phrases) {
		if false {
			for i, phrase := range phrases {
				log.Printf("Phrase[%d]: %s", i, phrase)
			}
			for i, output := range outputs {
				log.Printf("Output[%d]: %s", i, output.Phrase)
			}
		}
		log.Printf("len(outputs)=%d but len(phrases)=%d for phrases=%v",
			len(outputs), len(phrases), phrases)
		return []Output{}
	}

	return outputs
}

func SaveParse(phrase, parseJson, phraseDir string) {
	path := phraseDir + "/" + phrase + ".json"
	err := ioutil.WriteFile(path, []byte(parseJson), 0644)
	if err != nil {
		panic(err)
	}
}

func LoadSavedParse(phrase string, phraseDir string) Output {
	path := phraseDir + "/" + phrase + ".json"
	parseJson, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	parse := unmarshalParseJson(string(parseJson))
	return Output{
		Phrase:    phrase,
		ParseJson: string(parseJson),
		Parse:     parse,
	}
}

func ParsePhrasesWithFreelingCached(phrases []string,
	freelingHostAndPort string, parseDir string) {

	newPhrases := []string{}
	for i, phrase := range phrases {
		if strings.HasSuffix(phrase, "...") {
			phrase = phrase + " ." // prevent freeling from responding in two parts
		}

		if !fileExists(parseDir + "/" + phrase + ".json") {
			newPhrases = append(newPhrases, phrase)
		}

		if len(newPhrases) >= BATCH_SIZE || i == len(phrases)-1 {
			outputs := ParsePhrasesWithFreeling(newPhrases, freelingHostAndPort)
			newPhrases = []string{}

			for _, output := range outputs {
				SaveParse(output.Phrase, output.ParseJson, parseDir)
				fmt.Fprintf(os.Stderr, "%s\n", parseDir+"/"+output.Phrase+".json")
			}
		}
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
