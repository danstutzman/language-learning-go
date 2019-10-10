package model

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"strings"
)

type Output struct {
	Phrase       string
	AnalysisJson string
	Analysis     Analysis
}

type Analysis struct {
	Sentences []Sentence `json:"sentences"`
}

type Sentence struct {
	Id           string        `json:"id"`
	Tokens       []Token       `json:"tokens"`
	Constituents []Constituent `json:"constituents"`
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

func (token *Token) IsPronoun() bool {
	return strings.HasPrefix(token.Tag, "P")
}

func (token *Token) IsInterjection() bool {
	return strings.HasPrefix(token.Tag, "I")
}

func (token *Token) IsNoun() bool {
	return strings.HasPrefix(token.Tag, "N")
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

func parseAnalysisJson(analysisJson string) Analysis {
	var analysis Analysis
	err := json.Unmarshal([]byte(analysisJson), &analysis)
	if err != nil {
		panic(err)
	}
	return analysis
}

/* Example analysis JSON is [ { "sentences" : [
      { "id":"1",
        "tokens" : [
           { "id" : "t1.1", "begin" : "0", "end" : "5", "form" : "Estoy", "lemma" : "estar", "tag" : "VMIP1S0", "ctag" : "VMI", "pos" : "verb", "type" : "main", "mood" : "indicative", "tense" : "present", "person" : "1", "num" : "singular"},
           { "id" : "t1.2", "begin" : "6", "end" : "11", "form" : "feliz", "lemma" : "feliz", "tag" : "AQ0CS00", "ctag" : "AQ", "pos" : "adjective", "type" : "qualificative", "gen" : "common", "num" : "singular"},
           { "id" : "t1.3", "begin" : "11", "end" : "12", "form" : ".", "lemma" : ".", "tag" : "Fp", "ctag" : "Fp", "pos" : "punctuation", "type" : "period"}],
        "constituents" : [
          {"label" : "S", "children" : [
            {"label" : "grup-verb", "children" : [
              {"label" : "verb", "head" : "1", "children" : [
                {"leaf" : "1", "head" : "1", "token" : "t1.1", "word" : "Estoy"}
              ]}
            ]},
            {"label" : "s-adj", "children" : [
              {"label" : "s-a-ms", "head" : "1", "children" : [
                {"label" : "a-ms", "head" : "1", "children" : [
                  {"leaf" : "1", "head" : "1", "token" : "t1.2", "word" : "feliz"}
                ]}
              ]}
            ]},
            {"label" : "F-term", "children" : [
              {"leaf" : "1", "head" : "1", "token" : "t1.3", "word" : "."}
            ]}
          ]}]}]}
]*/

func AnalyzePhrasesWithFreeling(phrases []string,
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
			analysisJson := strings.TrimSuffix(output, "\x00")
			analysis := parseAnalysisJson(analysisJson)
			output := Output{
				Phrase:       phrase,
				AnalysisJson: analysisJson,
				Analysis:     analysis,
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
			analysisJson := strings.TrimSuffix(output, "\x00")
			analysis := parseAnalysisJson(analysisJson)
			output := Output{
				Phrase:       phrase,
				AnalysisJson: analysisJson,
				Analysis:     analysis,
			}
			outputs = append(outputs, output)
		}
	}

	if len(outputs) != len(phrases) {
		log.Panicf("len(outputs)=%d but len(phrases)=%d",
			len(outputs), len(phrases))
	}

	return outputs
}
