package main

import (
	"bufio"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data" // embed character set data in binary
	"log"
	"os"
	"strings"
)

type Row struct {
	form string
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

	trs, err := htmlquery.QueryAll(doc, "/html/body/form/table/tbody/tr")
	if err != nil {
		panic(err)
	}

	for _, tr := range trs {
		tds := htmlquery.Find(tr, "td")
		if len(tds) == 5 {
			a := htmlquery.FindOne(tds[2], "a")
			if a != nil {
				href := ""
				for _, attr := range a.Attr {
					if attr.Key == "href" {
						if strings.HasPrefix(attr.Val, "x3.asp?") {
							href = attr.Val
						}
					}
				}

				freq := htmlquery.InnerText(tds[3])
				fmt.Printf("%v,%s,%s\n",
					strings.TrimSpace(htmlquery.InnerText(a)),
					href,
					strings.TrimSpace(freq))
			}
		}
	}
}
