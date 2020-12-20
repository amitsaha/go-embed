package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tmpl := template.New("test")
	tmpl, err := tmpl.Parse(string(tmplMainGo))
	if err != nil {
		log.Fatal("Error Parsing template: ", err)
		return
	}

	type data struct {
		Name string
	}

	d := data{Name: "Jane"}
	err1 := tmpl.Execute(os.Stdout, d)
	if err1 != nil {
		log.Fatal("Error executing template: ", err1)
	}
}
