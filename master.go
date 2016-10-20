package main

import "fmt"

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
func (m *Master) loadPuzzle(file string, reply *Nothing) error {

	return nil
}

func (m *Master) savePuzzle(file string, reply *Nothing) error {

	return nil
}

/*
func (ng *Nonogram) String(junk Nothing, reply *string) error {
	*reply = fmt.Sprintf("Width: %d\nHeight: %d\n", ng.Width, ng.Height)
	return nil
}
*/
