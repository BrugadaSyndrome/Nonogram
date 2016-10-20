package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"runtime"
)

const (
	debugNone     = 0
	debugNormal   = 1
	debugDetailed = 2
	debugVerbose  = 3

	markEmpty   = 0
	markFilled  = 1
	markCrossed = 2
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

func startNonogramMaster(m *Master, address string) {
	rpc.Register(m)
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
	puzzle.Width = 5
	puzzle.Height = 5

	puzzle.Board = make([][]int, puzzle.Height)
	for i := range puzzle.Board {
		puzzle.Board[i] = make([]int, puzzle.Width)
	}

	master := new(Master)
	master.Nonogram = puzzle

	/*
		for i := 0; i < puzzle.Height; i++ {
			for j := 0; j < puzzle.Width; j++ {
				fmt.Print(puzzle.Board[i][j])
			}
			fmt.Print("\n")
		}
	*/
	fmt.Print(puzzle)

	startNonogramMaster(master, masterAddress)

}
