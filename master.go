package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// FillCell is called to fill cell with the value markFilled if cell value is markEmpty
func (m *Master) FillCell(cell Position, reply *bool) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	if m.Nonogram.Board[cell.X][cell.Y] == markEmpty {
		m.Nonogram.Board[cell.X][cell.Y] = markFilled
		debugMessage(debugNormal, fmt.Sprintf("Filled (%d,%d)", cell.X, cell.Y))
		*reply = true
	} else {
		debugMessage(debugNormal, fmt.Sprintf("Can't fill (%d,%d), not empty", cell.X, cell.Y))
		*reply = false
	}

	return nil
}

// CrossCell is called to fill cell with the value markCrossed if cell value is markEmpty
func (m *Master) CrossCell(cell Position, reply *bool) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	if m.Nonogram.Board[cell.X][cell.Y] == markEmpty {
		m.Nonogram.Board[cell.X][cell.Y] = markCrossed
		debugMessage(debugNormal, fmt.Sprintf("Crossed (%d,%d)", cell.X, cell.Y))
		*reply = true
	} else {
		debugMessage(debugNormal, fmt.Sprintf("Can't cross (%d,%d), not empty", cell.X, cell.Y))
		*reply = false
	}

	return nil
}

/*
	Width Height
	Column Hints x Width
	Row Hints x Height
*/
func (m *Master) LoadPuzzle(file string, reply *bool) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	fin, err := os.Open(file)

	if fin != nil {
		scanner := bufio.NewScanner(fin)

		// read in width and height first
		scanner.Scan()
		widthThenHeight := strings.Fields(scanner.Text())

		// make puzzle instance
		puzzle := new(Nonogram)
		puzzle.Width, _ = strconv.Atoi(widthThenHeight[0])
		puzzle.Height, _ = strconv.Atoi(widthThenHeight[1])
		puzzle.Board = make([][]int, puzzle.Height)
		for i := range puzzle.Board {
			puzzle.Board[i] = make([]int, puzzle.Width)
		}
		puzzle.ColumnHints = make([][]int, puzzle.Width)
		puzzle.RowHints = make([][]int, puzzle.Height)

		// read in rest of the puzzle
		//lines := make([]string, puzzle.Width+puzzle.Height)
		i := 0
		for scanner.Scan() {
			hints := strings.Fields(scanner.Text())
			intHints := make([]int, len(hints))
			for j := 0; j < len(hints); j++ {
				intHints[j], _ = strconv.Atoi(hints[j])
			}

			if i < puzzle.Width {
				puzzle.ColumnHints[i] = intHints
			} else {
				puzzle.RowHints[i%puzzle.Height] = intHints
			}
			//fmt.Println(hints, intHints)
			i++
		}

		debugMessage(debugVerbose, "\n"+puzzle.String())

		debugMessage(debugNone, fmt.Sprintf("[Master] Puzzle loaded: %s", file))
		fin.Close()
		*reply = true
	} else {
		debugMessage(debugNormal, fmt.Sprintf("[Master] Puzzle not loaded: %s", file))
		debugMessage(debugNormal, fmt.Sprintf("[Master] Error %v", err))
		*reply = false
	}
	return nil
}

func (m *Master) SavePuzzle(file string, reply *bool) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	*reply = false

	return nil
}
