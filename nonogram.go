package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type nonogramContext struct {
	Master *master
}

type nonogramHandler struct {
	Context *nonogramContext
	H       func(*nonogramContext, http.ResponseWriter, *http.Request)
}

func (nh nonogramHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nh.H(nh.Context, w, r)
}

// Mark is an enum for what is put into a cell on a board
type mark int

const (
	empty mark = iota
	filled
	crossed
)

func (m mark) String() string {
	var txt string
	switch m {
	case 0:
		txt = "empty"
	case 1:
		txt = "fill"
	case 2:
		txt = "cross"
	default:
		txt = "Undefined!"
	}
	return fmt.Sprintf("%s", txt)
}

// Move represents where a mark is placed
// Mark is the type of mark
// X is the column position of the mark
// Y is the row position of the mark
type move struct {
	Mark mark
	X    int
	Y    int
}

func (mv move) String() string {
	return fmt.Sprintf("%s (%d,%d)", mv.Mark, mv.X, mv.Y)
}

func (mv move) Map() (tmpMap map[string]int) {
	tmpMap = make(map[string]int)
	tmpMap["Mark"] = int(mv.Mark)
	tmpMap["X"] = mv.X
	tmpMap["Y"] = mv.Y
	return
}

// Nonogram represents the state of a nonogram puzzle
// Board stores the marks made on the cells of the puzzle
// ColumnHints stores the hints needed to solve the puzzle
// Height number of rows in the puzzle
// RowHints stores the hints needed to solve the puzzle
// Width number of columns in the puzzle
type nonogram struct {
	Board       [][]mark `json:"board"`
	ColumnHints [][]int  `json: "columnHints"`
	Height      int      `json: "height"`
	RowHints    [][]int  `json: "rowHints"`
	Width       int      `json: "width"`
}

func loadNonogram(path string) (n nonogram) {
	file, err := ioutil.ReadFile(path)
	checkError(err, "Unable to load nonogram from JSON file.")

	err = json.Unmarshal(file, &n)
	checkError(err, "Could not unmarshal JSON file.")

	n.Board = make([][]mark, n.Height)
	for i := 0; i < n.Height; i++ {
		n.Board[i] = make([]mark, n.Width)
	}
	return
}
