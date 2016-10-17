package Nonogram

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

func startNonogramMaster(ng *Nonogram, address string) {
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
	startNonogramMaster(puzzle, masterAddress)

}
