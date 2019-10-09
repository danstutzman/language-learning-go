package model

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func AnalyzePhrasesWithFreeling(phrases []string,
	freelingHostAndPort string) []string {

	analysisJsons := []string{}

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
			analysisJsons = append(analysisJsons, strings.TrimSuffix(output, "\x00"))
		}
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
		analysisJsons = append(analysisJsons, strings.TrimSuffix(output, "\x00"))
	}

	return analysisJsons
}
