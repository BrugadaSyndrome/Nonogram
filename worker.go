package main

// Worker
// ID of worker
// Inbox is the channel that the master will send updates to
// Update is the channel that the workers will send updates to the master
// Puzzle workers working replica of masters puzzle
type worker struct {
	ID     int
	Inbox  chan<- mark
	Outbox <-chan mark
	Puzzle nonogram
}
