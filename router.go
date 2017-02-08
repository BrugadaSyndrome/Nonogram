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
		//fmt.Println("handleMoves() has control.")
		err := enc.Encode(ctx.Master.MoveList)
		checkError(err, "Unable to prepare JSON.")
		// empty the list
		ctx.Master.MoveList = []map[string]int{}
		ctx.Master.Mux.Unlock()
		//fmt.Println("handleMoves() gives up control.")

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func handleSolve(ctx *nonogramContext, w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)

		ctx.Master.Manage()

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
