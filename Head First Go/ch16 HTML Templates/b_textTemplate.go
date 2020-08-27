package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	text := "Here's my template!\n"               // template text
	tmpl, err := template.New("test").Parse(text) // create new template based on text
	check(err)
	err = tmpl.Execute(os.Stdout, nil) // writeout template text to terminal instead of httpresponse
	check(err)
}
