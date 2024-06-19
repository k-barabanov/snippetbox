package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/" {
	//	http.NotFound(w, r)
	//	return
	//}
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox\n"))

}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...\n", id)
	//msg := fmt.Sprintf("Display a specific snippet with ID %d...\n", id) // the same functionality
	//w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet...\n"))

	//if r.Method != http.MethodPost {
	//	w.Header().Set("Allow", http.MethodPost)
	//	//w.Header().Add("Rofler", "keker")
	//	//w.Header()["Date"] = nil
	//	http.Error(w, "Method not allowed", 405)
	//	return
	//}

	// to cover JSON
	/*
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"name":"Alex"}`))
	*/
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet...\n"))
}

func main() {
	mux := http.NewServeMux()                             // routes register
	mux.HandleFunc("GET /{$}", home)                      // Restrict this route to exact matches on / only.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // Add the {id} wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("Listening on http://localhost:4000...")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
