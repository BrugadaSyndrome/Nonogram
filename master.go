package main

import (
	"bufio"
	"fmt"
	"os"
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

// TODO: make file format for puzzles
/*
width height
1
1 2
...
1 3 //width x height times
*/
func (m *Master) LoadPuzzle(file string, reply *bool) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	puzzle := new(Nonogram)
	puzzle.Width = 5
	puzzle.Height = 5
	puzzle.Board = make([][]int, puzzle.Height)
	for i := range puzzle.Board {
		puzzle.Board[i] = make([]int, puzzle.Width)
	}

	fin, err := os.Open(file)

	if err != nil {
		scanner := bufio.NewScanner(fin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			fmt.Println(scanner.Bytes())
		}

		debugMessage(debugNone, fmt.Sprintf("[Master] Puzzle loaded: %s", file))
		fin.Close()
		*reply = true
	} else {
		debugMessage(debugNormal, fmt.Sprintf("[Master] Puzzle not loaded: %s", file))
		fmt.Println(err)
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

func (m *Master) SetAddress(address string, reply *bool) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.Address = address
	*reply = true

	debugMessage(debugNormal, fmt.Sprintf("[Master] Address set: %s", address))
	return nil
}
