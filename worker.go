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
	case boxesAndSpaces:
		fmt.Printf("[Worker %d] job is: %s\n", w.ID, job)
		w.BoxesAndSpaces()
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

func (w worker) BoxesAndSpaces() {
	/*
		- What if the board already has filled cells?
		! With puzzle1.json, row 2 is filled in incorrectly according to the final solution.
			It is still correct according to the solving method. Will need to keep an eye on
			occurances like this and develop a method to handle these 'contradictions'.
	*/
	fmt.Printf("[Worker %d] is running Boxes\n", w.ID)

	for rowIndex, hints := range w.Puzzle.RowHints {
		num := len(hints)
		L, R := make([]mark, w.Puzzle.Width), make([]mark, w.Puzzle.Width)
		a, b := 0, 0

		for i, j := 0, num-1; i < num && j >= 0; i, j = i+1, j-1 {
			// left fill
			for z := a; z < a+hints[i]; z++ {
				L[z] = filled
			}
			a += hints[i] + 1
			// only pad with a cross if not off the edge of the board
			if a-1 < w.Puzzle.Width {
				L[a-1] = crossed
			}

			// right fill
			for z := b; z < b+hints[j]; z++ {
				R[w.Puzzle.Width-1-z] = filled
			}
			b += hints[j] + 1
			// only pad with a cross if not off the edge of the board
			if w.Puzzle.Width-b >= 0 {
				R[w.Puzzle.Width-b] = crossed
			}

		}
		//fmt.Printf("[Worker %d] L: %v || R: %v\n", w.ID, L, R)

		for i := 0; i < w.Puzzle.Width; i++ {
			if L[i] == R[i] {
				if L[i] == filled {
					w.Outbox <- move{filled, rowIndex, i}
				} else if L[i] == crossed {
					w.Outbox <- move{crossed, rowIndex, i}
				}
			}
		}
	}

	fmt.Printf("[Worker %d] Done running Boxes.\n", w.ID)
}

func newWorker(n nonogram, id int, masterInbox chan move) (w worker) {
	w.ID = id
	w.Inbox = make(chan move)
	w.Outbox = masterInbox
	w.Puzzle = n
	return
}
