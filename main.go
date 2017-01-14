package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	templates *template.Template
)

func init() {
	// parse all template files
	var allFiles []string
	files, err := ioutil.ReadDir("./static/templates")
	checkError(err, "Unable to read from template directory.")

	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".tmpl") {
			allFiles = append(allFiles, "./static/templates/"+filename)
		}
	}

	templates, err = template.ParseFiles(allFiles...)
	checkError(err, "Unable to parse all templates.")
}

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
