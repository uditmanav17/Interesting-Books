package main

import (
	"log"
	"os"
	"text/template"
)

type Part struct {
	Name  string
	Count int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// create template based on string - text
// forward data to execute method on template
func execTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	execTemplate("Dot is : {{.}}\n", "ABC")
	execTemplate("Dot is : {{.}}\n", 123)

	// conditionals
	execTemplate("\nSTART {{if .}} Dot is true {{end}} finish\n", true)
	execTemplate("START {{if .}} Dot is true {{end}} finish\n", false)

	// loops
	templateText := "\nBefore loop: {{.}}\n{{range .}}In Loop: {{.}}\n{{end}}After Loop: {{.}}\n"
	execTemplate(templateText, []string{"do", "re", "mi"})

	// struct in templates
	templateText2 := "\nName: {{.Name}}\nCount: {{.Count}}\n"
	newPart := Part{"Bolts", 99} // fields of struct need to be exported
	execTemplate(templateText2, newPart)
}
