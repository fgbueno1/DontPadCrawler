package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v3"
)

const (
	URL_BASE   = "http://dontpad.com/"
	URL_SEARCH = URL_BASE + "{name}"
	FILE_PATH  = "CrawlerResults/"
)

type configuration struct {
	Keywords []string `yaml:"keywords"`
}

func main() {
	var config configuration
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Println(err)
	}
	for _, keyword := range config.Keywords {
		result, err := crawler(keyword)
		if err != nil {
			log.Println(err)
			continue
		}
		currentTime := time.Now()
		fileName := fmt.Sprintf("%v%v-%v.html", FILE_PATH, keyword, currentTime.Format("2006-01-02-15-04-05"))
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

func crawler(keyword string) (response string, err error) {
	cli := resty.New()
	resp, err := cli.R().SetPathParam("name", keyword).Get(URL_SEARCH)
	if err != nil {
		return
	}
	response = string(resp.Body())
	return response, nil
}
