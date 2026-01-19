package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const PORT = "8080"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, PORT)
	log.Fatal(srv.ListenAndServe())
}
