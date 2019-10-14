package main

import (
	"bufio"
	"github.com/antchfx/htmlquery"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data" // embed character set data in binary
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 1+1 {
		log.Fatalf("1st should be html file to parse")
	}
	htmlPath := os.Args[1]

	file, err := os.Open(htmlPath)
	if err != nil {
		panic(err)
	}

	reader, err := charset.NewReader("latin1", bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	doc, err := htmlquery.Parse(reader)
	if err != nil {
		panic(err)
	}

	nodes, err := htmlquery.QueryAll(doc, "/html/body/form/table/tbody/tr/td/a")
	if err != nil {
		panic(err)
	}

	texts := []string{}
	for _, node := range nodes {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				if strings.HasPrefix(attr.Val, "x3.asp?") {
					text := strings.TrimSpace(node.FirstChild.Data)
					if text != "" {
						texts = append(texts, text)
					}
				}
			}
		}
	}
	for _, text := range texts {
		println(text)
	}
}
