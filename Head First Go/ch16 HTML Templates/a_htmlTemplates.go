package main

import (
	"log"
	"net/http"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	// placeholder := []byte("Signature list goes here!")
	html, err := template.ParseFiles("test.html") // use content of view to create new template
	check(err)
	err = html.Execute(writer, nil) // write template content to ResponseWriter
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
