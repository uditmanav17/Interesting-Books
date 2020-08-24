package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getWebpage(url string) []uint8 {
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
	return body
}

func main() {
	pages := []string{
		"https://www.washingtonpost.com/sitemaps/index.xml",
		"https://httpbin.org/get",
	}

	page := pages[1]
	body := getWebpage(page)
	fmt.Println("URL -", page)
	fmt.Println("Size of page (bytes) -", len(body))
	fmt.Println("Content -\n", string(body))
}
