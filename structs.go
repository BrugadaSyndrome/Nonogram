package Nonogram

import "sync"

type Nothing struct{}

type Position struct {
	X int
	Y int
}

type Nonogram struct {
	Board [Height][Width]int
	Mutex sync.Mutex
}
