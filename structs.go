package main

import (
	"fmt"
	"sync"
)

type Nothing struct{}

type Position struct {
	X int
	Y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)\n", p.X, p.Y)
}

type Nonogram struct {
	Board  [][]int
	Width  int
	Height int
}

func (ng Nonogram) String() string {
	return fmt.Sprintf("Width: %d\nHeight: %d\n", ng.Width, ng.Height)
}

type Master struct {
	Nonogram *Nonogram
	Mutex    sync.Mutex
}

type Worker struct {
	Nonogram *Nonogram
	Mutex    sync.Mutex
}
