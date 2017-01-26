package main

import (
	"bytes"
	"encoding/json"
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
		fmt.Printf("handleIndex: %p\n", serverMaster)

		context := indexData{
			Master:  *serverMaster,
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
		var buffer bytes.Buffer
		enc := json.NewEncoder(&buffer)

		fmt.Printf("handleMoves: %p\n", serverMaster)

		serverMaster.Mux.Lock()
		fmt.Println("handleMovess() has control.")
		err := enc.Encode(serverMaster.MoveList)
		checkError(err, "Unable to prepare JSON.")
		//serverMaster.MoveList = []map[string]int{} // empty list
		serverMaster.Mux.Unlock()
		fmt.Println("handleMovess() gives up control.")

		fmt.Println(buffer.String())
	}
}
