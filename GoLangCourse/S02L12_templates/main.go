package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tplContainer *template.Template

var fm = template.FuncMap{
	"uc":     strings.ToUpper,
	"append": appendX,
}

func init() {
	// create a new empty template, pass in the functions map, parse the files
	tplContainer = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func main() {
	testTemplateWithFunc()
}

func appendX(s string) string {
	return s + "X"
}

func testTemplateWithFunc() {
	numbers := map[string]string{
		"one":   "unu",
		"two":   "doi",
		"three": "trei",
	}
	err := tplContainer.ExecuteTemplate(os.Stdout, "template.gohtml", numbers)
	logError(err)
}

func logError(e error) {
	if e != nil {
		log.Println(e)
	}
}
