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

func (w worker) processInbox() {
	for mv := range w.Inbox {
		fmt.Printf("[Worker %d] Recieved move: %s\n", w.ID, mv)
		w.Puzzle.Board[mv.X][mv.Y] = mv.Mark
	}
}

func (w worker) Work() {
	fmt.Printf("[Worker %d] Starting work.\n", w.ID)
	go w.processInbox()

	for {
		job := <-w.Jobs
		//fmt.Printf("[Worker %d] job is: %s\n", w.ID, job)

		switch job {
		case boxesAndSpaces:
			w.BoxesAndSpaces(w.Puzzle.RowHints, true)
			w.BoxesAndSpaces(w.Puzzle.ColumnHints, false)
		case forcing:
			fmt.Printf("[Worker %d] Running Forcing\n", w.ID)
		case glue:
			fmt.Printf("[Worker %d] Running Glue\n", w.ID)
		case joining:
			fmt.Printf("[Worker %d] Running Joining\n", w.ID)
		case splitting:
			fmt.Printf("[Worker %d] Running Splitting\n", w.ID)
		case punctuating:
			fmt.Printf("[Worker %d] Running Punctuating\n", w.ID)
		case mercury:
			fmt.Printf("[Worker %d] Running Mercury\n", w.ID)
		default:
			log.Fatalf("[Worker %d] Unknown job: int(%d) string(%s)", w.ID, job, job)
		}

		w.Jobs <- job
		//fmt.Printf("[Worker %d] Done working. Returning job: %s.\n", w.ID, job)
		break
	}
}

func (w worker) BoxesAndSpaces(hintList [][]int, horizontal bool) {
	/*
		! With puzzle1.json, row 2 is filled in incorrectly according to the final solution.
			It is still correct according to the solving method. Will need to keep an eye on
			occurances like this and develop a method to handle these 'contradictions'.
	*/
	fmt.Printf("[Worker %d] is running Boxes/Spaces (horizontal=%t)\n", w.ID, horizontal)

	for rowIndex, hints := range hintList {
		num := len(hints)
		var listCount int
		if horizontal {
			listCount = w.Puzzle.Width
		} else {
			listCount = w.Puzzle.Height
		}
		L, R := make([]mark, listCount), make([]mark, listCount)
		a, b := 0, 0

		for i, j := 0, num-1; i < num && j >= 0; i, j = i+1, j-1 {
			// left fill
			for z := a; z < a+hints[i]; z++ {
				L[z] = maybeFilled
			}
			a += hints[i] + 1
			// only pad with a cross if not off the edge of the board
			if a-1 < listCount {
				L[a-1] = maybeCrossed
			}

			// right fill
			for z := b; z < b+hints[j]; z++ {
				R[listCount-1-z] = maybeFilled
			}
			b += hints[j] + 1
			// only pad with a cross if not off the edge of the board
			if listCount-b >= 0 {
				R[listCount-b] = maybeCrossed
			}

		}
		//fmt.Printf("[Worker %d] L: %v || R: %v\n", w.ID, L, R)

		for i := 0; i < listCount; i++ {
			if L[i] == R[i] {
				var x, y int
				if horizontal {
					x = rowIndex
					y = i
				} else {
					x = i
					y = rowIndex
				}

				if L[i] == maybeFilled {
					w.Outbox <- move{maybeFilled, 0, x, y}
				} else if L[i] == maybeCrossed {
					w.Outbox <- move{maybeCrossed, 0, x, y}
				}
			}
		}
	}

	fmt.Printf("[Worker %d] Done running Boxes/Spaces\n", w.ID)
}

func newWorker(n nonogram, id int, masterInbox chan move) (w worker) {
	w.ID = id
	w.Inbox = make(chan move)
	w.Outbox = masterInbox
	w.Puzzle = n
	return
}
