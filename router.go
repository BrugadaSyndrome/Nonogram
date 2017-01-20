package main

import "net/http"

type indexData struct {
	Log     []string
	Master  master
	Title   string
	Workers []worker
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)

		n := loadNonogramFromJSON("./static/puzzles/puzzle1.json")
		master, workers := newMaster(n, 2)

		context := indexData{
			Log:     []string{},
			Master:  master,
			Title:   "Nonogram Solver",
			Workers: workers,
		}

		err := templates.Execute(w, context)
		checkError(err, "Failed to execute templates.")

		master.Manage()

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

type logData struct {
	Log []string
}
