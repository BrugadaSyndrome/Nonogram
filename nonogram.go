package main

// Mark is an enum for what is put into a cell on a board
type mark int

const (
	empty mark = iota
	filled
	crossed
)

// Move represents where a mark is placed
// Mark is the type of mark
// X is the column position of the mark
// Y is the row position of the mark
type move struct {
	Mark mark
	X    int
	Y    int
}

// Nonogram represents the state of a nonogram puzzle
// Board stores the marks made on the cells of the puzzle
// Height number of rows in the puzzle
// Hints stores the hints needed to solve the puzzle
// Width number of columns in the puzzle
type nonogram struct {
	Board  [][]mark
	Height int
	Hints  [][]int
	Width  int
}
