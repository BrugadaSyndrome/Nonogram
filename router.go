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
	w.WriteHeader(http.StatusOK)

	n := sampleNonogram()

	m := master{}
	m.Puzzle = n

	w1 := worker{}
	w1.ID = 1
	w1.Puzzle = n

	w2 := worker{}
	w2.ID = 2
	w2.Puzzle = n

	context := indexData{
		Log:     []string{"msg 1", "msg 2", "msg 3", "msg 4", "msg 5", "msg 6", "msg 7", "msg 8", "msg 9", "msg 10"},
		Master:  m,
		Title:   "Nonogram Solver",
		Workers: []worker{w1, w2},
	}

	err := templates.Execute(w, context)
	checkError(err, "Failed to execute templates.")
}
