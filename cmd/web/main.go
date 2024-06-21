package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // routes register

	fileServer := http.FileServer(http.Dir("./ui/static/"))             // static files
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer)) // static files

	mux.HandleFunc("GET /{$}", home)                      // Restrict this route to exact matches on / only.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // Add the {id} wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("Listening on http://localhost:4000...")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
