package server

import (
	"net/http"
	"text/template"
)

func Error(w http.ResponseWriter, code int) {
	tml, err := template.ParseFiles("../html/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.WriteHeader(code)
	err = tml.Execute(w, http.StatusText(code))
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
