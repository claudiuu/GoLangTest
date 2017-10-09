package main

import (
	"log"
	"os"
	"text/template"
)

var tplContainer *template.Template

/*
* Init function will be executed when the program initializes.
* It is a special function just like main
 */
func init() {
	tplContainer = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	// testParseFiles()
	testParseGlob()
}

func testParseGlob() {

	// this will execute the first template it finds
	err := tplContainer.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tplContainer.ExecuteTemplate(os.Stdout, "three.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}

func testParseFiles() {
	// the template file can have any extensions
	// but the recommended standard is .gohtml
	// tpl is a pointer to template
	// it is a container of more templates, all that have been parsed
	tpl, err := template.ParseFiles("templates/tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	idxFile, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}

	defer idxFile.Close()

	err = tpl.Execute(idxFile, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
