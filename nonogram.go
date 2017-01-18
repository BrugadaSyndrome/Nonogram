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
	Board       [][]mark
	ColumnHints [][]int
	Height      int
	RowHints    [][]int
	Width       int
}

func sampleNonogram() (n nonogram) {
	height := 5
	width := 5

	n.Board = make([][]mark, height)
	for i := 0; i < height; i++ {
		n.Board[i] = make([]mark, width)
	}
	//fmt.Println(n.Board)
	n.ColumnHints = make([][]int, width)
	n.ColumnHints[0] = []int{5}
	n.ColumnHints[1] = []int{1, 1}
	n.ColumnHints[2] = []int{1}
	n.ColumnHints[3] = []int{1, 1}
	n.ColumnHints[4] = []int{5}
	//fmt.Println(n.ColumnHints)
	/*
		n.ColumnHints = make([][]int, width)
		for i := 0; i < width; i++ {
			// need hints from user...
			n.ColumnHints[i] = make([]int, 0)
		}
	*/
	n.Height = 5
	n.RowHints = make([][]int, height)
	n.RowHints[0] = []int{1, 3}
	n.RowHints[1] = []int{1, 1}
	n.RowHints[2] = []int{3}
	n.RowHints[3] = []int{2, 2}
	n.RowHints[4] = []int{1, 1, 1}
	/*
		n.RowHints = make([][]int, height)
		for i := 0; i < height; i++ {
			// need hints from user...
			n.RowHints[i] = make([]int, 0)
		}
	*/
	//fmt.Println(n.RowHints)
	n.Width = 5
	return
}
