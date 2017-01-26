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
	fmt.Printf("[Worker %d] Starting work.\n", w.ID)

	job := <-w.Jobs

	switch job {
	case boxes:
		fmt.Printf("[Worker %d] job is: %s\n", w.ID, job)
		w.Boxes()
	case spaces:
		fmt.Printf("[Worker %d] job is: %s\n", w.ID, job)
	case forcing:
		fmt.Printf("[Worker %d] job is: %s\n", w.ID, job)
	case glue:
		fmt.Printf("[Worker %d] job is: %s\n", w.ID, job)
	case joining:
		fmt.Printf("[Worker %d] job is: %s\n", w.ID, job)
	case splitting:
		fmt.Printf("[Worker %d] job is: %s\n", w.ID, job)
	case punctuating:
		fmt.Printf("[Worker %d] job is: %s\n", w.ID, job)
	case mercury:
		fmt.Printf("[Worker %d] job is: %s\n", w.ID, job)
	default:
		log.Fatalf("Worker got unknown job: int(%d) string(%s)", job, job)
	}

	w.Jobs <- job
	fmt.Printf("[Worker %d] Done working. Returning job: %s.\n", w.ID, job)
}

func (w worker) Boxes() {
	/*
		- What if the board already has filled cells?
	*/
	fmt.Printf("[Worker %d] is running Boxes\n", w.ID)

	for rowIndex, row := range w.Puzzle.RowHints {
		L := make([]mark, w.Puzzle.Width)
		i := 0
		for _, hint := range row {
			for z := i; z < i+hint; z++ {
				L[z] = filled
			}
			i += hint
			L[i] = crossed
			i++
		}
		//fmt.Printf("[Worker %d] L: %v\n", w.ID, L)

		for i, j := 0, len(L)-1; i < len(L) && j >= 0; i, j = i+1, j-1 {
			if L[i] == L[j] && L[i] == filled && L[j] == filled {
				w.Outbox <- move{filled, rowIndex, i}
			}
		}
	}

	fmt.Printf("[Worker %d] Done running boxes.\n", w.ID)
}

func newWorker(n nonogram, id int, masterInbox chan move) (w worker) {
	w.ID = id
	w.Inbox = make(chan move)
	w.Outbox = masterInbox
	w.Puzzle = n
	return
}
