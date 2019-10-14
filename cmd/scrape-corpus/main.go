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
	if len(os.Args) != 1+1 || strings.TrimSpace(os.Args[1]) == "" {
		log.Fatalf("1st arg: query (for example: 'cara J')")
	}
	query := strings.TrimSpace(os.Args[1])

	params := url.Values{}
	params.Add("chooser", "seq")
	params.Add("wl", "4")
	params.Add("wr", "4")
	params.Add("p", query)
	params.Add("sortBy", "freq")
	params.Add("numhits", "100")
	params.Add("groupBy", "words")
	params.Encode()

	req, err := http.NewRequest("POST",
		"https://www.corpusdelespanol.org/now/x2.asp",
		strings.NewReader(params.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Referer", "https://www.corpusdelespanol.org/now/x1.asp?w=1440&h=900&c=nowsp")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	path := "db/4_corpus_scrapes/" + query + ".html"
	file, err := os.Create(path)
	fmt.Fprintln(os.Stderr, path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	io.Copy(writer, resp.Body)
	writer.Flush()
}
