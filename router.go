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

	worker1 := worker{}
	worker1.ID = 1
	worker1.Puzzle.Height = 1
	worker1.Puzzle.Width = 2
	worker1.Puzzle.ExecuteTemplate()

	templateString := string(fin)
	indexTemplate, err := template.New("index").Parse(templateString)
	checkError(err, "Failed to parse indexTemplate.")

	data := indexData{
		Title:   "Nonogram Solver",
		Master:  "Master data will go here",
		Log:     []string{},
		Workers: []worker{worker1},
	}

	err = indexTemplate.Execute(w, data)
	checkError(err, "Failed to execute indexTemplate.")
}
