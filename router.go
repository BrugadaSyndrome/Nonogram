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

	fin, err := ioutil.ReadFile("static/templates/index.html")
	checkError(err, "Failed to read index.html.")

	templateString := string(fin)
	indexTemplate, err := template.New("index").Parse(templateString)
	checkError(err, "Failed to parse indexTemplate.")

	worker1 := worker{}
	worker1.ID = 1
	worker2 := worker{}
	worker2.ID = 2

	data := indexData{
		Title:   "Nonogram Solver",
		Master:  "Master data will go here",
		Log:     []string{},
		Workers: []worker{worker1, worker2},
	}
	err = indexTemplate.Execute(w, data)
	checkError(err, "Failed to execute indexTemplate.")
}
