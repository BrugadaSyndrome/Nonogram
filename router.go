package main

import "net/http"

type indexData struct {
	Log     []string
	Master  master
	Title   string
	Workers []worker
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" && req.Method == "GET" {
		w.WriteHeader(http.StatusOK)

		n := loadNonogram("./static/puzzles/puzzle2.json")
		master := newMaster(n, 1)

		context := indexData{
			Log:     []string{},
			Master:  master,
			Title:   "Nonogram Solver",
			Workers: master.Workers,
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
