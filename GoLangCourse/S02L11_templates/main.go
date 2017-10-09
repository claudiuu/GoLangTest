package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var tplContainer *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	tplContainer = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	// testTemplateWithSlice()
	// testTemplateWithMap()
	testTemplateWithStruct()
}

func testTemplateWithStruct() {
	p1 := person{"Claudiu", 28}
	p2 := person{"Olga", 29}
	p3 := person{"Lissa", 6}
	persons := []person{p1, p2, p3}

	err := tplContainer.ExecuteTemplate(os.Stdout, "template_struct.gohtml", persons)
	logError(err)
}

func testTemplateWithMap() {
	numbers := map[string]string{
		"one":   "unu",
		"two":   "doi",
		"three": "trei",
	}
	err := tplContainer.ExecuteTemplate(os.Stdout, "template_map.gohtml", numbers)
	logError(err)
}

func testTemplateWithSlice() {
	names := []string{"Claudiu", "Olga", "Sasha"}
	err := tplContainer.ExecuteTemplate(os.Stdout, "template_slice.gohtml", names)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("\n#####################################################")
	err = tplContainer.ExecuteTemplate(os.Stdout, "template_slice2.gohtml", names)
	logError(err)
}

func logError(e error) {
	if e != nil {
		log.Println(e)
	}
}
