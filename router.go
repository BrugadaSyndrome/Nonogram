package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type indexData struct {
	Master master
	Title  string
}

func handleIndex(ctx *nonogramContext, w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)

		context := indexData{
			Master: *ctx.Master,
			Title:  "Nonogram Solver",
		}

		err := templates.Execute(w, context)
		checkError(err, "Failed to execute templates.")

		/*
			## TODO/BUG ##
			- This needs to be called once and only once!
			- [ ] Move to its own handler
			- [ ] That handler will be called by the client when it is ready to see the puzzle solved
		*/
		ctx.Master.Manage()
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func handleMoves(ctx *nonogramContext, w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)

		fmt.Println("Moves requested")
		enc := json.NewEncoder(w)

		ctx.Master.Mux.Lock()
		fmt.Println("handleMoves() has control.")
		err := enc.Encode(ctx.Master.MoveList)
		checkError(err, "Unable to prepare JSON.")
		ctx.Master.MoveList = []map[string]int{} // empty list
		ctx.Master.Mux.Unlock()
		fmt.Println("handleMoves() gives up control.")

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
