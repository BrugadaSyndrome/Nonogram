package main

import "fmt"

// Master
// Inbox is the channel that the workers will send update to
// Outboxes is a list of channels that workers will listen on to receive updates from the master
// Puzzle master instance of puzzle
type master struct {
	Inbox      <-chan move
	NumWorkers int
	Outboxes   []chan<- move
	Puzzle     nonogram
}

func (m master) Manage() {
	fmt.Println("Master is managing.")
	for {

	}
}

func newMaster(n nonogram, numWorkers int) (m master, w []worker) {
	m.Inbox = make(<-chan move, numWorkers)
	m.NumWorkers = numWorkers
	m.Outboxes = make([]chan<- move, numWorkers)
	m.Puzzle = n

	w = make([]worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		w[i] = newWorker(n, i+1, m.Inbox)
		m.Outboxes[i] = w[i].Inbox
		go w[i].Work()
	}

	return
}
