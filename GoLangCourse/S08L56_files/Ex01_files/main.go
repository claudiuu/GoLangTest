package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles("res/templates/index.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(handle))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handle(res http.ResponseWriter, req *http.Request) {
	var body string
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("theFile")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		fmt.Println("\nheader:", h)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		body = string(bs)

		fmt.Println("Writing the content to a new file on the server")
		nf, err := os.Create("./res/user/" + h.Filename)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		_, err = nf.Write(bs)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}

	templates.Execute(res, body)
}
