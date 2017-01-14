package main

import "net/http"

type indexData struct {
	Title   string
	Master  string
	Log     []string
	Workers []worker
}

type logData struct {
	Log []string
}

func index(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	context := indexData{
		Title: "Nonogram Solver",
		Log:   []string{"[w1] log", "[w2] log", "[m] log"},
	}

	err := templates.Execute(w, context)
	checkError(err, "Failed to execute templates.")
}
