package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func printWebpageStats(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(reflect.TypeOf(body))
	fmt.Println("URL -", url)
	fmt.Println("Size of page (bytes) -", len(body))
	// fmt.Println("Content -\n", string(body))
}

func main() {
	pages := []string{
		"https://www.washingtonpost.com/sitemaps/index.xml",
		"https://httpbin.org/get",
	}

	for _, page := range pages {
		fmt.Println("started Routine with URL -", page)
		go printWebpageStats(page)

	}
	sleepTime := 10 // try increasing or decreasing
	// fmt.Println(reflect.TypeOf(time.Second))
	time.Sleep(time.Duration(sleepTime) * time.Second)

}
