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

const (
	CRAWLER_FILE_PATH = "CrawlerResults/"
	PARSED_FILE_PATH  = "Parsed Files/"
)

func main() {
	files, err := ioutil.ReadDir(CRAWLER_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.Contains(file.Name(), ".html") {
			continue
		}
		err = parser(file.Name())
		if err != nil {
			log.Print(err)
			continue
		}
		err = os.Rename(CRAWLER_FILE_PATH+file.Name(), PARSED_FILE_PATH+file.Name())
		if err != nil {
			log.Print(err)
		}
	}
}

func parser(fileName string) error {
	f, e := os.Open(CRAWLER_FILE_PATH + fileName)
	if e != nil {
		return e
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return err
	}

	doc.Find("#text").Each(func(i int, s *goquery.Selection) {
		inside_html, _ := s.Html()
		decodedHtml := html.UnescapeString(inside_html)
		fmt.Printf(decodedHtml)

	})
	return nil
}
