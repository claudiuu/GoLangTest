package main

import (
	"html/template"
	"net/http"
)

var templateContainer *template.Template

func init() {
	templateContainer = template.Must(template.ParseGlob("res/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	templateContainer.ExecuteTemplate(res, "index.gohtml", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	templateContainer.ExecuteTemplate(res, "dog.gohtml", nil)
}

func me(res http.ResponseWriter, req *http.Request) {
	templateContainer.ExecuteTemplate(res, "me.gohtml", nil)
}
