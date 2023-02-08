package server

import (
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/artist/", result)
	fileServer := http.FileServer(http.Dir("../static/img/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	log.Println("listen server on http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Println(err)
}
