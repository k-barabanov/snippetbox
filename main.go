package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hi, it`s Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing a note..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		//w.Header().Add("Rofler", "keker")
		//w.Header()["Date"] = nil

		http.Error(w, "Method not allowed", 405)

		return
	}
	w.Write([]byte("Create new note..."))
}

func main() {
	mux := http.NewServeMux() // регистрация маршрутов
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Listening on http://localhost:4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
