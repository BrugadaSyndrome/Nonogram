package main

import (
	"fmt"
)

// Worker
// ID of worker
// Inbox is the channel that the master will send updates to
// Update is the channel that the workers will send updates to the master
// Puzzle workers working replica of masters puzzle
type worker struct {
	ID     int
	Inbox  chan<- move
	Outbox <-chan move
	Puzzle nonogram
}

func (w worker) Work() {
	fmt.Printf("Worker %d is working.", w.ID)
	for {

	}
}

func newWorker(n nonogram, id int, masterInbox <-chan move) (w worker) {
	w.ID = id
	w.Inbox = make(chan<- move)
	w.Outbox = masterInbox
	w.Puzzle = n
	return
}
