package main

// Master
// Inbox is the channel that the workers will send update to
// Outboxes is a list of channels that workers will listen on to receive updates from the master
// Puzzle master instance of puzzle
type master struct {
	Inbox    <-chan mark
	Outboxes []chan<- mark
	Puzzle   nonogram
}
