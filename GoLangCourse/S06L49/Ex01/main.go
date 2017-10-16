package main

import (
	"html/template"
	"io"
	"net/http"
)

var templateContainer *template.Template

func init() {
	templateContainer = template.Must(template.ParseFiles("res/index.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/res/", http.StripPrefix("/res", http.FileServer(http.Dir("res/img/"))))
	http.Handle("/favicon.ico", http.HandlerFunc(fav))

	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "blah, index")
}

func dog(res http.ResponseWriter, req *http.Request) {
	templateContainer.Execute(res, nil)
}

func dogServer(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "res/img/dog.jpg")
}

func fav(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "res/img/favicon.ico")
}
