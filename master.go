package Nonogram

import "fmt"

// FillCell is called to fill cell with the value markFilled if cell value is markEmpty
func (ng *Nonogram) FillCell(cell Position, reply *bool) error {
	ng.Mutex.Lock()
	defer ng.Mutex.Unlock()

	if ng.Board[cell.X][cell.Y] == markEmpty {
		ng.Board[cell.X][cell.Y] = markFilled
		debugMessage(debugNormal, fmt.Sprintf("Filled (%d,%d)", cell.X, cell.Y))
		*reply = true
	} else {
		debugMessage(debugNormal, fmt.Sprintf("Can't fill (%d,%d), not empty", cell.X, cell.Y))
		*reply = false
	}

	return nil
}

// CrossCell is called to fill cell with the value markCrossed if cell value is markEmpty
func (ng *Nonogram) CrossCell(cell Position, reply *bool) error {
	ng.Mutex.Lock()
	defer ng.Mutex.Unlock()

	if ng.Board[cell.X][cell.Y] == markEmpty {
		ng.Board[cell.X][cell.Y] = markCrossed
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
func (ng *Nonogram) loadPuzzle(file string, reply *bool) error {

	*reply = false
	return nil
}

func (ng *Nonogram) savePuzzle(file string, reply *bool) error {

	*reply = false
	return nil
}
