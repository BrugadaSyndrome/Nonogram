package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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
	From int
	Mark mark
	X    int
	Y    int
}

func (m move) String() string {
	return fmt.Sprintf("%s cell (%d,%d)", m.Mark, m.X, m.Y)
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

func sampleNonogram() (n nonogram) {
	height := 5
	width := 5

	n.Board = make([][]mark, height)
	for i := 0; i < height; i++ {
		n.Board[i] = make([]mark, width)
	}

	n.ColumnHints = make([][]int, width)
	n.ColumnHints[0] = []int{1, 2}
	n.ColumnHints[1] = []int{1, 1}
	n.ColumnHints[2] = []int{1, 1, 1}
	n.ColumnHints[3] = []int{1, 2}
	n.ColumnHints[4] = []int{5}

	n.Height = 5

	n.RowHints = make([][]int, height)
	n.RowHints[0] = []int{1, 3}
	n.RowHints[1] = []int{1, 1}
	n.RowHints[2] = []int{3}
	n.RowHints[3] = []int{2, 2}
	n.RowHints[4] = []int{1, 1, 1}

	n.Width = 5
	return
}
