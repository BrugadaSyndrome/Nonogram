package main

import (
	"fmt"
	"log"
)

// Worker
// ID of worker
// Inbox is the channel that the master will send updates to
// Update is the channel that the workers will send updates to the master
// Puzzle workers working replica of masters puzzle
type worker struct {
	ID     int
	Inbox  chan move
	Jobs   chan method
	Log    []string
	Outbox chan move
	Puzzle nonogram
}

func (w worker) Work() {
	fmt.Printf("[Worker %d] starting work.\n", w.ID)

	job := <-w.Jobs
	fmt.Printf("[Worker %d] got job: %d\n", w.ID, job)

	switch job {
	case boxes:
		fmt.Printf("[Worker %d] method is: %s\n", w.ID, job)
	case spaces:
		fmt.Printf("[Worker %d] method is: %s\n", w.ID, job)
	case forcing:
		fmt.Printf("[Worker %d] method is: %s\n", w.ID, job)
	case glue:
		fmt.Printf("[Worker %d] method is: %s\n", w.ID, job)
	case joining:
		fmt.Printf("[Worker %d] method is: %s\n", w.ID, job)
	case splitting:
		fmt.Printf("[Worker %d] method is: %s\n", w.ID, job)
	case punctuating:
		fmt.Printf("[Worker %d] method is: %s\n", w.ID, job)
	case mercury:
		fmt.Printf("[Worker %d] method is: %s\n", w.ID, job)
	default:
		log.Fatalf("Worker got unknown job: %d", job)
	}

	mv := move{w.ID, filled, w.ID, w.ID}
	w.Outbox <- mv
	fmt.Printf("[Worker %d] sent move: %s\n", w.ID, mv)

	w.Jobs <- job
	fmt.Printf("[Worker %d] returned job: %d\n", w.ID, job)

	fmt.Printf("[Worker %d] done working.\n", w.ID)
}

func newWorker(n nonogram, id int, masterInbox chan move) (w worker) {
	w.ID = id
	w.Inbox = make(chan move)
	w.Outbox = masterInbox
	w.Puzzle = n
	return
}
