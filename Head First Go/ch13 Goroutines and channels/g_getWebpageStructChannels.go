package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type page struct {
	URL  string
	Size int
}

func getWebpageStats(url string, channel chan page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- page{URL: url, Size: len(body)}
}

func main() {
	webResponses := make(chan page)
	pages := []string{
		"https://www.washingtonpost.com/sitemaps/index.xml/",
		"https://httpbin.org/get/",
		"http://www.example.com/",
		"https://golang.org/",
	}

	for _, page := range pages {
		fmt.Println("started Routine with URL -", page)
		go getWebpageStats(page, webResponses)
	}
	var resp page
	for i := 0; i < len(pages); i++ {
		resp = <-webResponses
		fmt.Printf("URL - %s\t\tSize - %d\n", resp.URL, resp.Size)
	}
}
