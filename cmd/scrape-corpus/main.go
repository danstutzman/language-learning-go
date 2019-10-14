package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2+1 || strings.TrimSpace(os.Args[1]) == "" {
		log.Fatalf("1st arg: x2 or x3")
		log.Fatalf("2nd arg: query (for example: 'cara J')")
	}
	page := os.Args[1]
	query := strings.TrimSpace(os.Args[2])

	var req *http.Request
	var err error
	if page == "x2" {
		params := url.Values{}
		params.Add("chooser", "seq")
		params.Add("wl", "4")
		params.Add("wr", "4")
		params.Add("p", query)
		params.Add("sortBy", "freq")
		params.Add("numhits", "100")
		params.Add("groupBy", "words")

		req, err = http.NewRequest("POST",
			"https://www.corpusdelespanol.org/now/x2.asp",
			strings.NewReader(params.Encode()))
		req.Header.Set("Referer",
			"https://www.corpusdelespanol.org/now/x1.asp?w=1440&h=900&c=nowsp")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			panic(err)
		}
	} else if page == "x3" {
		params := url.Values{}
		params.Add("xx", "1")
		params.Add("w11", "cara")
		//params.Add("w12", "nueva.[n*]")
		params.Add("w12", "es")
		params.Add("r", "")

		req, err = http.NewRequest("GET",
			"https://www.corpusdelespanol.org/now/x3.asp?"+params.Encode(), nil)
		if err != nil {
			panic(err)
		}
		req.Header.Set("Referer", "https://www.corpusdelespanol.org/now/x2.asp")
		log.Fatalf("Need to set this Cookie value")
		req.Header.Set("Cookie",
			"ASPSESSIONIDCGBSRCDA=AAAAAAAAAAAAAAAAAAAAAAAA; ii=1")
	} else {
		log.Fatalf("Unknown page")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	path := "db/4_corpus_scrapes/" + page + "/" + query + ".html"
	file, err := os.Create(path)
	fmt.Fprintln(os.Stderr, path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	io.Copy(writer, resp.Body)
	writer.Flush()

	// Need to save cookie from x2 for x3
	for key, values := range resp.Header {
		log.Printf("%s -> %v", key, values)
	}
}
