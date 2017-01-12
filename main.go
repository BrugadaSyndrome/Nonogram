package main

import (
	"html/template"
	"io/ioutil"
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

// INDEX //
type indexData struct {
	Title   string
	Master  string
	Log     []string
	Workers []worker
}

func index(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	fin, err := ioutil.ReadFile("static/index.html")
	checkError(err, "Failed to read index.html.")

	templateString := string(fin)
	indexTemplate, err := template.New("index").Parse(templateString)
	checkError(err, "Failed to parse templateString.")

	data := indexData{
		Title:   "Nonogram Solver",
		Master:  "Master Board will be here",
		Log:     []string{},
		Workers: []worker{worker{1}, worker{2}},
	}
	err = indexTemplate.Execute(w, data)
	checkError(err, "Failed to execute indexTemplate.")
}

type worker struct {
	ID int
}

func checkError(err error, errMessage string) {
	if err != nil {
		log.Fatal(errMessage)
	}
}
