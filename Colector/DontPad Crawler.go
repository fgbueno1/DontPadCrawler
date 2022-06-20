package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	URL_BASE   = "http://dontpad.com/"
	URL_SEARCH = URL_BASE + "{name}"
)

func crawler(keyword string) (response string, err error) {
	cli := resty.New()
	resp, err := cli.R().SetPathParam("name", keyword).Get(URL_SEARCH)
	if err != nil {
		return
	}
	response = string(resp.Body())
	return response, nil
}

func main() {
	keywords := []string{
		"test",
		"mydontpad",
	}
	for _, keyword := range keywords {
		result, err := crawler(keyword)
		if err != nil {
			log.Println(err)
			continue
		}
		currentTime := time.Now()
		fileName := fmt.Sprintf("%v-%v.html", keyword, currentTime.Format("2006-01-02-15-04-05"))
		f, err := os.OpenFile(fileName, os.O_CREATE, 0766)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		_, err = f.WriteString(string(result))
		if err != nil {
			log.Println(err)
		}
	}
}
