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
	DebugLevel     = debugVerbose
	CurrentPort    = 8080
	MachineAddress = getLocalAddress()
)

func startNonogramMaster(m *Master) {
	rpc.Register(m)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", m.Address)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	master := new(Master)
	startNonogramMaster(master)

	var reply bool
	master.SetAddress(fmt.Sprintf("%s:%d", MachineAddress, CurrentPort), &reply)
	master.LoadPuzzle("nonogram1.txt", &reply)

	/*
		for i := 0; i < puzzle.Height; i++ {
			for j := 0; j < puzzle.Width; j++ {
				fmt.Print(puzzle.Board[i][j])
			}
			fmt.Print("\n")
		}
	*/

}
