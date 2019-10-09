package model

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"strconv"
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
	Id            string `json:"id"`
	Begin         string `json:"begin"`
	End           string `json:"end"`
	Form          string `json:"form"`
	Lemma         string `json:"lemma"`
	Tag           string `json:"tag"`
	Ctag          string `json:"ctag"`
	Pos           string `json:"pos"`
	Type          string `json:"type"`
	Mood          string `json:"mood"`
	Person        string `json:"person"`
	Num           string `json:"num"`
	Gen           string `json:"gen"`
	BeginInPhrase int    // relative to phrase instead of all phrases
	EndInPhrase   int    // relative to phrase instead of all phrases
}

func (token *Token) IsPunctuation() bool {
	return strings.HasPrefix(token.Tag, "F")
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

func augmentAnalysis(analysis Analysis, phraseBeginIndex int) Analysis {
	for i := range analysis.Sentences {
		sentence := &analysis.Sentences[i]
		for j := range sentence.Tokens {
			token := &sentence.Tokens[j]

			begin, err := strconv.Atoi(token.Begin)
			if err != nil {
				panic(err)
			}
			token.BeginInPhrase = begin - phraseBeginIndex

			end, err := strconv.Atoi(token.End)
			if err != nil {
				panic(err)
			}
			token.EndInPhrase = end - phraseBeginIndex
		}
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
	phraseBeginIndex := 0
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
			analysis :=
				augmentAnalysis(parseAnalysisJson(analysisJson), phraseBeginIndex)
			output := Output{
				Phrase:       phrase,
				AnalysisJson: analysisJson,
				Analysis:     analysis,
			}
			outputs = append(outputs, output)
		}

		// Freeling measures offsets by Unicode codepoints (runes), not by bytes
		// Plus one for the newline between phrases?
		phraseBeginIndex += len([]rune(phrase)) + 1
	}

	log.Printf("Writing FLUSH_BUFFER...\n")
	_, err = conn.Write([]byte("FLUSH_BUFFER\x00"))
	if err != nil {
		panic(err)
	}

	log.Printf("Reading...\n")
	output, err := reader.ReadString('\x00')
	if err != nil {
		panic(err)
	}
	if output != "FL-SERVER-READY\x00" {
		panic("Unexpected output " + output)
	}

	return outputs
}
