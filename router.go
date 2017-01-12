package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

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

	worker1 := worker{}
	worker2 := worker{}

	data := indexData{
		Title:   "Nonogram Solver",
		Master:  "Master Board will be here",
		Log:     []string{},
		Workers: []worker{worker1, worker2},
	}
	err = indexTemplate.Execute(w, data)
	checkError(err, "Failed to execute indexTemplate.")
}
