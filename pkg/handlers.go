package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Println(0.1)
		Error(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		fmt.Println(0.2)
		Error(w, http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("../html/index.html")
	if err != nil {
		fmt.Println(1)
		Error(w, http.StatusInternalServerError)
		return
	}
	var test []Artists
	test = InfArtist()
	err = tmp.Execute(w, test)
	if err != nil {
		fmt.Println(2)
		Error(w, http.StatusInternalServerError)
		return
	}
}

func result(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println(0.3)
		Error(w, http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("../html/artist.html")
	if err != nil {
		fmt.Println(3)
		Error(w, http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 || id > 52 {
		Error(w, http.StatusNotFound)
		return
	}

	test := NameArtist(id)
	concert := LocationDateArtist(id)
	test.Relations = concert
	err = tmp.Execute(w, test)
	if err != nil {
		fmt.Println(4)
		Error(w, http.StatusInternalServerError)
		return
	}
}
