package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

var (
	templates *template.Template
)

func init() {
	// parse all template files
	var allFiles []string
	files, err := ioutil.ReadDir("./static/templates")
	checkError(err, "Unable to read from template directory.")

	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".tmpl") {
			allFiles = append(allFiles, "./static/templates/"+filename)
		}
	}

	functions := template.FuncMap{
		"add":         add,
		"ascending":   ascending,
		"descending":  descending,
		"longestList": longestList,
		"subtract":    subtract,
	}

	templates, err = template.New(filepath.Base(allFiles[0])).Funcs(functions).ParseFiles(allFiles...)
	checkError(err, "Unable to parse all templates.")
}

func main() {
	// set up a file server for static files
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	ctx := &nonogramContext{Master: newMaster(loadNonogram("./static/puzzles/puzzle1.json"), 1)}
	ctx.Master.Manage()

	// handle URLs
	http.Handle("/", nonogramHandler{ctx, handleIndex})
	http.Handle("/moves", nonogramHandler{ctx, handleMoves})

	// run server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func checkError(err error, errMessage string) {
	if err != nil {
		log.Fatalf("\n[err] %v\n[msg] %s", err, errMessage)
	}
}
