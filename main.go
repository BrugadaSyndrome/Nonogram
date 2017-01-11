package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// INDEX //
type indexData struct {
	Title   string
	Content string
}

func index(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	fin, err := ioutil.ReadFile("index.html")
	check(err, "Failed to read index.html.")

	templateString := string(fin)
	indexTemplate, err := template.New("index").Parse(templateString)
	check(err, "Failed to parse templateString.")

	data := indexData{
		Title:   "Nonogram Solver",
		Content: "Welcome to root!",
	}
	err = indexTemplate.Execute(w, data)
	check(err, "Failed to execute indexTemplate.")
}

// MISC //
func check(err error, errMessage string) {
	if err != nil {
		log.Fatal(errMessage)
	}
}
