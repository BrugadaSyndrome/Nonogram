package main

import (
	"fmt"
	"net/http"
)

type indexData struct {
	Master  master
	Seconds int
	Title   string
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)

		n := loadNonogram("./static/puzzles/puzzle2.json")
		serverMaster := newMaster(n, 2)

		context := indexData{
			Master:  serverMaster,
			Seconds: 3,
			Title:   "Nonogram Solver",
		}

		err := templates.Execute(w, context)
		checkError(err, "Failed to execute templates.")

		serverMaster.Manage()
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func handleMoves(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Moves requested")
	}
}
