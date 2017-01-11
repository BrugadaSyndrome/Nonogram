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
	Master  string
	Log     []string
	Workers []Worker
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
		Master:  "Master Board will be here",
		Log:     []string{"log 1", "log 2"},
		Workers: []Worker{Worker{1}, Worker{2}},
	}
	err = indexTemplate.Execute(w, data)
	check(err, "Failed to execute indexTemplate.")
}

// MISC //
type Worker struct {
	ID int
}

func check(err error, errMessage string) {
	if err != nil {
		log.Fatal(errMessage)
	}
}
