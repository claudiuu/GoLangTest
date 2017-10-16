package main

import (
	"net/http"

	"github.com/claudiuu/GoLangCourse/S04L32_http/Ex01_form/httphandler"
)

func main() {
	var handler httphandler.FormHandler
	http.ListenAndServe(":8080", handler)
}
