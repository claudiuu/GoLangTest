package httphandler

import (
	"html/template"
	"log"
	"net/http"
)

// FormHandler handler
type FormHandler struct {
}

func (fh FormHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	var tpl *template.Template
	tpl = template.Must(template.ParseFiles("res/index.gohtml"))
	tpl.Execute(w, req.Form)
}
