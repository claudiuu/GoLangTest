package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "index")
	})
	http.HandleFunc("/dog/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "dog path")
	})
	http.HandleFunc("/me/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello, Claudiu")
	})
	http.ListenAndServe(":8080", nil)
}
