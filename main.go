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

var (
	DebugLevel     = debugVerbose
	CurrentPort    = 8080
	MachineAddress = getLocalAddress()
)

func startNonogramMaster(master *Master, masterAddress string, puzzlePath string) {
	// start rpc server
	rpc.Register(master)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", masterAddress)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	// load board
	var reply bool
	err := call(masterAddress, "Master.LoadPuzzle", puzzlePath, &reply)
	if err != nil {
		debugMessage(debugNone, "Something went wrong with RPC call")
	}
	if reply == false {
		debugMessage(debugNormal, "Something went wrong when loading board")
	} else {
		debugMessage(debugNormal, "Puzzle is ready to be solved")
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	master := new(Master)
	masterAddress := fmt.Sprintf("%s:%d", MachineAddress, CurrentPort)
	puzzlePath := "nonogram1.txt"

	startNonogramMaster(master, masterAddress, puzzlePath)
}
