package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func analyzeWithFreeling(input string, freelingHostAndPort string) string {
	allOutput := ""

	log.Printf("Conecting to %s\n", freelingHostAndPort)
	conn, err := net.Dial("tcp", freelingHostAndPort)
	if err != nil {
		panic(err)
	}

	log.Printf("Writing RESET_STATS...\n")
	_, err = conn.Write([]byte("RESET_STATS\x00"))
	if err != nil {
		panic(err)
	}

	log.Printf("Reading...\n")
	serverReady, err := bufio.NewReader(conn).ReadString('\x00')
	if err != nil {
		panic(err)
	}
	if serverReady != "FL-SERVER-READY\x00" {
		panic("Server not ready?")
	}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		log.Printf("Writing...\n")
		_, err := conn.Write([]byte(line + "\x00"))
		if err != nil {
			panic(err)
		}

		log.Printf("Reading...\n")
		output, err := bufio.NewReader(conn).ReadString('\x00')
		if err != nil {
			panic(err)
		}
		if output != "FL-SERVER-READY\x00" {
			allOutput += output
		}
	}

	log.Printf("Writing FLUSH_BUFFER...\n")
	_, err = conn.Write([]byte("FLUSH_BUFFER\x00"))
	if err != nil {
		panic(err)
	}

	log.Printf("Reading...\n")
	output, err := bufio.NewReader(conn).ReadString('\x00')
	if err != nil {
		panic(err)
	}
	if output != "FL-SERVER-READY\x00" {
		allOutput += output
	}

	return allOutput
}

func main() {
	freelingHostAndPort := os.Getenv("FREELING_HOST_AND_PORT")
	log.Println(analyzeWithFreeling("estoy feliz.", freelingHostAndPort))
}
