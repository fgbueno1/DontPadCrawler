package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parser(fileName string) {
	f, e := os.Open(fileName)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#text").Each(func(i int, s *goquery.Selection) {
		inside_html, _ := s.Html()
		decodedHtml := html.UnescapeString(inside_html)
		fmt.Printf(decodedHtml)

	})
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.Contains(file.Name(), ".html") {
			continue
		}
		parser(file.Name())
	}
}
