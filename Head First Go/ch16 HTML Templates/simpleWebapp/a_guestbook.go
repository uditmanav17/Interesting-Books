package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// GuestBook to hold signatures
type GuestBook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html") // use content of view to create new template
	check(err)
	err = html.Execute(writer, nil) // write template content to ResponseWriter
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature) //write sign to file
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	fmt.Printf("%#v", signatures)
	guestbook := GuestBook{
		Signatures:     signatures,
		SignatureCount: len(signatures),
	}
	html, err := template.ParseFiles("view.html") // use content of view to create new template
	check(err)
	err = html.Execute(writer, guestbook) // write template content to ResponseWriter
	check(err)
}

// to get signatures from txt file
func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName) // opening file
	if os.IsNotExist(err) {        // if error returned saying file not exists
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err()) // report any scanning error and exit
	return lines

}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
