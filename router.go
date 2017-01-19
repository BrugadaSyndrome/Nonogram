package main

import (
	"bytes"
	"fmt"
	"net/http"
)

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

		var buf bytes.Buffer
		buf.Write([]byte("Master and Workers are ready to go."))
		http.Post("http://localhost:8080/log", "text/plain", &buf)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

type logData struct {
	msg string
}

func updateLog(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		w.WriteHeader(http.StatusOK)

		var p []byte
		n, err := req.Body.Read(p)
		checkError(err, "?")
		fmt.Println(n)

		//fmt.Println( req.Body)
		//fmt.Println(req.ContentLength)
		//fmt.Println(req.Form)
		//fmt.Println(req.Header)
		//fmt.Println(req.Host)
		//fmt.Println(req.Method)
		//fmt.Println(req.MultipartForm)
		//fmt.Println(req.PostForm)
		//fmt.Println(req.Proto)
		//fmt.Println(req.RemoteAddr)
		//fmt.Println(req.RequestURI)
		//fmt.Println(req.Response)
		//fmt.Println(req.TLS)
		//fmt.Println(req.Trailer)
		//fmt.Println(req.TransferEncoding)
		//fmt.Println(req.URL)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
