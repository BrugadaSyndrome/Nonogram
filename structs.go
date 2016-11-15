package main

import (
	"fmt"
	"sync"
)

type Nothing struct{}

type Position struct {
	X int
	Y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)\n", p.X, p.Y)
}

type Nonogram struct {
	Board       [][]int
	ColumnHints [][]int
	Height      int
	RowHints    [][]int
	Width       int
}

func (ng Nonogram) String() string {
	var boardString string

	// border
	boardString += "   "
	for i := 0; i < len(ng.ColumnHints); i++ {
		boardString += "_ "
	}
	boardString += "\n"

	// marks and row hints
	for i := 0; i < ng.Height; i++ {
		// marks
		boardString += " | "
		for j := 0; j < ng.Width; j++ {
			boardString += fmt.Sprintf("%d ", ng.Board[i][j])
		}
		// row hints
		boardString += "| "
		for j := 0; j < len(ng.RowHints[i]); j++ {
			boardString += fmt.Sprintf("%d ", ng.RowHints[i][j])
		}
		boardString += "\n"
	}

	// border
	boardString += "   "
	for i := 0; i < len(ng.ColumnHints); i++ {
		boardString += "_ "
	}

	// need to determine longest column of hints for boundry detection
	var max int
	for i := 0; i < len(ng.ColumnHints); i++ {
		if len(ng.ColumnHints[i]) > max {
			max = len(ng.ColumnHints[i])
		}
	}

	// columns hints
	boardString += "\n"
	for j := 0; j <= max; j++ {
		boardString += "   "
		for i := 0; i < len(ng.ColumnHints); i++ {
			if j < len(ng.ColumnHints[i]) {
				boardString += fmt.Sprintf("%d ", ng.ColumnHints[i][j])
			} else {
				boardString += "  "
			}
		}
		boardString += "\n"
	}

	return boardString
}

type Master struct {
	Workers  []string
	Nonogram *Nonogram
	Mutex    sync.Mutex
}

type Worker struct {
	Master   string
	Nonogram *Nonogram
	Mutex    sync.Mutex
}
