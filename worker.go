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
		//fmt.Printf("[Worker %d] Recieved move: %s\n", w.ID, mv)
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
			w.Forcing()
			return
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
	}
}

func (w worker) BoxesAndSpaces(hintList [][]int, horizontal bool) {
	/*
		Does this method need to check the cells on the board?
		Does checking the cells help this method solve better/faster?
	*/
	if horizontal {
		fmt.Printf("[Worker %d] is running Boxes/Spaces (horizontal)\n", w.ID)
	} else {
		fmt.Printf("[Worker %d] is running Boxes/Spaces (vertical)\n", w.ID)
	}

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
				//if w.Puzzle.Board[rowIndex][z] == empty {
				L[z] = maybeFilled
				//} else {
				//	L[z] = w.Puzzle.Board[rowIndex][z]
				//}
			}
			a += hints[i] + 1
			// only pad with a cross if not off the edge of the board
			if a-1 < listCount {
				L[a-1] = maybeCrossed
			}

			// right fill
			for z := b; z < b+hints[j]; z++ {
				//if w.Puzzle.Board[rowIndex][z] == empty {
				R[listCount-1-z] = maybeFilled
				//} else {
				//	R[z] = w.Puzzle.Board[rowIndex][z]
				//}
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

func (w worker) Forcing() {
	fmt.Printf("[Worker %d] Running Forcing\n", w.ID)

	for rowIndex, row := range w.Puzzle.Board {
		begin, count := 0, 0
		fmt.Println(row)
		for itemIndex, item := range row {
			if item != crossed {
				count++
			} else {
				fmt.Printf("row: %d | count: %d | range: %d-%d\n", rowIndex, count, begin, itemIndex-1)
				begin = itemIndex + 1
				count = 0
			}
		}
		fmt.Printf("row: %d | count: %d | range: %d-%d\n", rowIndex, count, begin, w.Puzzle.Width-1)
	}

	fmt.Printf("[Worker %d] Done running Forcing\n", w.ID)
}

func newWorker(n nonogram, id int, masterInbox chan move) (w worker) {
	w.ID = id
	w.Inbox = make(chan move)
	w.Outbox = masterInbox
	w.Puzzle = n
	return
}
