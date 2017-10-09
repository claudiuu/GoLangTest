package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number string
	Name   string
	Grade  int
}

type semester struct {
	Name    string
	Courses []course
}

type year struct {
	TheYear   string
	Semesters []semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	y1 := year{
		"2015-2016",
		[]semester{
			semester{
				"S01-2015/2016",
				[]course{
					course{"CS-01", "Programing in Java", 10},
					course{"CS-02", "Programing in Go", 7},
				},
			},
			semester{
				"S02-2015/2016",
				[]course{
					course{"CS-03", "Web Development", 8},
					course{"CS-04", "Catch the python", 6},
				},
			},
		},
	}

	y2 := year{
		"2016-2017",
		[]semester{
			semester{
				"S01-2016/2017",
				[]course{
					course{"CS-05", "Programing in C", 10},
					course{"CS-06", "Programing in C++", 7},
				},
			},
			semester{
				"S02-2016/2017",
				[]course{
					course{"CS-03", "Web Development", 8},
					course{"CS-07", "Servers", 7},
				},
			},
		},
	}

	context := struct {
		Years []year
	}{
		[]year{y1, y2},
	}

	err := tpl.Execute(os.Stdout, context)
	logError(err)
}

func logError(e error) {
	if e != nil {
		log.Println(e)
	}
}
