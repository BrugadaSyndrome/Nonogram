package main

import (
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

		fmt.Printf("handleIndex: [%p] %+v\n", ctx.Master, ctx.Master.Puzzle.Board)

		err := templates.Execute(w, context)
		checkError(err, "Failed to execute templates.")

		ctx.Master.Manage()
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func handleMoves(ctx *nonogramContext, w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)

		context := indexData{
			Master: *ctx.Master,
			Title:  "Nonogram Solver",
		}

		fmt.Printf("handleMoves: [%p] %+v\n", ctx.Master, ctx.Master.Puzzle.Board)

		err := templates.Execute(w, context)
		checkError(err, "Failed to execute templates.")

		/*
			fmt.Println("Moves requested")
			var buffer bytes.Buffer
			enc := json.NewEncoder(&buffer)

			ctx.Master.Mux.Lock()
			fmt.Println("handleMoves() has control.")
			err := enc.Encode(ctx.Master.MoveList)
			checkError(err, "Unable to prepare JSON.")
			ctx.Master.MoveList = []map[string]int{} // empty list
			ctx.Master.Mux.Unlock()
			fmt.Println("handleMoves() gives up control.")

			fmt.Println(buffer.String())
		*/
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
