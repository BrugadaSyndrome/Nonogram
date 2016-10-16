package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"runtime"
)

// const
// Width: determines the width of the puzzle
// Height: determines the height of the puzzle
const (
	debugNone     = 0
	debugNormal   = 1
	debugDetailed = 2
	debugVerbose  = 3

	markEmpty   = 0
	markFilled  = 1
	markCrossed = 2

	Width  = 5
	Height = 10
)

// var
// DebugLevel: determines the amount of debug text in the console
// Port: port for the server to run on
// MasterAddress: address for the server to run on
var (
	DebugLevel    = debugVerbose
	port          = "8080"
	masterAddress = fmt.Sprintf("%s:%s", getLocalAddress(), port)
)

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

func startPuzzleMaster(ng *Nonogram, address string) {
	rpc.Register(ng)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", address)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	puzzle := new(Nonogram)
	startPuzzleMaster(puzzle, masterAddress)

	var cell = Position{1, 1}
	var reply bool
	call(masterAddress, "Nonogram.FillCell", cell, &reply)
	call(masterAddress, "Nonogram.FillCell", cell, &reply)
	call(masterAddress, "Nonogram.CrossCell", cell, &reply)

	var cell2 = Position{2, 2}
	var reply2 bool
	call(masterAddress, "Nonogram.CrossCell", cell2, &reply2)
	call(masterAddress, "Nonogram.CrossCell", cell2, &reply2)
	call(masterAddress, "Nonogram.FillCell", cell2, &reply2)

}
