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
	URL_BASE  = "https://api.dontpad.com/{query}.body.json?lastModified=0"
	FILE_PATH = "CrawlerResults/"
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
	dataToWrite := ""
	for _, keyword := range config.Keywords {
		result, err := crawler(keyword)
		if err != nil {
			log.Println(err)
			continue
		}
		dataToWrite += string(result) + "\n"
	}
	currentTime := time.Now()
	fileName := fmt.Sprintf("%v%v.json", FILE_PATH, currentTime.Format("2006-01-02-15-04-05"))
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	_, err = f.WriteString(dataToWrite)
	if err != nil {
		log.Println(err)
	}
}

func crawler(keyword string) (response string, err error) {
	cli := resty.New()
	resp, err := cli.R().SetPathParam("query", keyword).Get(URL_BASE)
	if err != nil {
		return
	}
	response = string(resp.Body())
	return response, nil
}
