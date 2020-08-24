package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func printWebpageStats(url string, channel chan []uint8) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("URL -", url)
	fmt.Println("Size of page (bytes) -", len(body))
	// fmt.Println("Content -\n", string(body))
	channel <- body
}

func main() {
	webResponses := make(chan []uint8)
	pages := []string{
		"https://www.washingtonpost.com/sitemaps/index.xml",
		"https://httpbin.org/get",
		"http://www.example.com/",
		"https://golang.org/",
	}

	for _, page := range pages {
		fmt.Println("started Routine with URL -", page)
		go printWebpageStats(page, webResponses)
	}
	for i := 0; i < len(pages); i++ {
		<-webResponses
	}
}
