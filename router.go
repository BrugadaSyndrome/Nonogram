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

	n := nonogram{}
	n.Height = 5
	n.Width = 5
	n.Board = make([][]mark, n.Height)
	for i := 0; i < n.Height; i++ {
		n.Board[i] = make([]mark, n.Width)
	}

	m := master{}
	m.Puzzle = n

	w1 := worker{}
	w1.ID = 1
	w1.Puzzle = n

	w2 := worker{}
	w2.ID = 2
	w2.Puzzle = n

	context := indexData{
		Log:     []string{"msg 1", "msg 2", "msg 3"},
		Master:  m,
		Title:   "Nonogram Solver",
		Workers: []worker{w1, w2},
	}

	err := templates.Execute(w, context)
	checkError(err, "Failed to execute templates.")
}
