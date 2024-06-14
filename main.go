package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yolo, Swagger!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Listening on http://localhost:4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
