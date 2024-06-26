package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux() // routes register

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")}) // static files | {} - is used to create a composite literal
	mux.Handle("/static", http.NotFoundHandler())                               // to user give an error when he goes to static
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))         // static files

	mux.HandleFunc("GET /{$}", home)                      // Restrict this route to exact matches on / only.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // Add the {id} wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("Listening on http://localhost:4000...")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

type neuteredFileSystem struct { // the code below closes files server listing to user ("/stats"), handles index.html and gives user a proper error
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
