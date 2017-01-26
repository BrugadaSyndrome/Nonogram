package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Method is an enum for methods of solving a nonogram puzzle
type method int

const (
	boxes method = iota
	spaces
	forcing
	glue
	joining
	splitting
	punctuating
	mercury
	numMethods
)

func (m method) String() string {
	var txt string
	switch m {
	case 0:
		txt = "boxes"
	case 1:
		txt = "spaces"
	case 2:
		txt = "forcing"
	case 3:
		txt = "glue"
	case 4:
		txt = "joining"
	case 5:
		txt = "splitting"
	case 6:
		txt = "punctuating"
	case 7:
		txt = "mercury"
	default:
		txt = "Undefined!"
	}
	return fmt.Sprintf("%s", txt)
}

// Master
// Inbox is the channel that the workers will send update to
// Outboxes is a list of channels that workers will listen on to receive updates from the master
// Puzzle master instance of puzzle
type master struct {
	Inbox      chan move
	Jobs       chan method
	NumWorkers int
	Outboxes   []chan move
	Prepare    chan move
	Puzzle     nonogram
	Workers    []worker
}

func (m master) Manage() {
	fmt.Println("Master is managing.")
	go m.processInbox()
	go m.prepareJSON()

	fmt.Println("Starting workers")
	for _, w := range m.Workers {
		go w.Work()
	}
}

func (m master) processInbox() {
	for mv := range m.Inbox {

		fmt.Printf("[Master] Recieved move: %s\n", mv)

		m.Puzzle.Board[mv.X][mv.Y] = mv.Mark
		m.Prepare <- mv

	}
}

func (m master) prepareJSON() {
	var moveList []map[string]int
	enc := json.NewEncoder(os.Stdout)
	for mv := range m.Prepare {
		moveList = append(moveList, mv.Map())
		err := enc.Encode(moveList)
		checkError(err, "Unable to prepare JSON.")
	}
}

func newMaster(n nonogram, numWorkers int) (m master) {
	m.Inbox = make(chan move, numWorkers)
	m.Jobs = make(chan method, numMethods)
	for i := 0; i < int(numMethods); i++ {
		m.Jobs <- method(i)
	}
	m.NumWorkers = numWorkers
	m.Outboxes = make([]chan move, numWorkers)
	m.Prepare = make(chan move, numWorkers)
	m.Puzzle = n
	m.Workers = make([]worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		m.Workers[i] = newWorker(n, i+1, m.Inbox)
		m.Workers[i].Jobs = m.Jobs
		m.Outboxes[i] = m.Workers[i].Inbox
	}

	return
}
