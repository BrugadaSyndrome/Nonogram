package main

import (
	"log"
	"net/http"
)

func main() {
	// set up a file server for static files
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// handle URLs
	http.HandleFunc("/", index)

	// run server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func checkError(err error, errMessage string) {
	if err != nil {
		log.Fatalf("\n[err] %v\n[msg] %s", err, errMessage)
	}
}
