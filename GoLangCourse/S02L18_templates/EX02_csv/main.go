package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	records := readFromCSV()
	// discard the header
	records = records[1:]

	err := tpl.ExecuteTemplate(os.Stdout, "page.gohtml", records)
	logError(err)
}

func readFromCSV() (records [][]string) {
	file, err := os.Open("res/table.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err = reader.ReadAll()
	logError(err)
	return records
}

func logError(e error) {
	if e != nil {
		log.Println(e)
	}
}
