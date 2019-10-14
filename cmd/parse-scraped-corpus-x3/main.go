package main

import (
	"bufio"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data" // embed character set data in binary
	"golang.org/x/net/html"
	"log"
	"os"
)

func gatherText(node *html.Node) string {
	if node == nil {
		return ""
	}

	text := ""

	if node.Type == html.TextNode {
		text += node.Data
	}

	node = node.FirstChild
	if node != nil {
		text += gatherText(node)
	}

	for node != nil {
		node = node.NextSibling
		text += gatherText(node)
	}

	return text
}

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

	nodes, err := htmlquery.QueryAll(doc, "/html/body/form/form/table/tbody/tr/td/span")
	if err != nil {
		panic(err)
	}

	texts := []string{}
	for _, node := range nodes {
		fmt.Println(gatherText(node))
	}
	for _, text := range texts {
		println(text)
	}
}
