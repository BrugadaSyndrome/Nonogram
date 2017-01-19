package main

import "net/http"

type indexData struct {
	Log     []string
	Master  master
	Title   string
	Workers []worker
}

type logData struct {
	Log []string
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" && req.Method == "GET" {
		w.WriteHeader(http.StatusOK)

		n := loadNonogramFromJSON("./static/puzzles/puzzle1.json")
		master, workers := newMaster(n, 2)

		context := indexData{
			Log:     []string{"msg 1", "msg 2", "msg 3", "msg 4", "msg 5", "msg 6", "msg 7", "msg 8", "msg 9", "msg 10"},
			Master:  master,
			Title:   "Nonogram Solver",
			Workers: workers,
		}

		err := templates.Execute(w, context)
		checkError(err, "Failed to execute templates.")

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
