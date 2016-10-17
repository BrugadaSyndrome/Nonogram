package Nonogram

import (
	"os"
	"testing"
)

func TestMarkEmptyCell(t *testing.T) {
	var reply bool
	// Fill
	var cell = Position{1, 1}
	call(masterAddress, "Nonogram.FillCell", cell, &reply)
	// Cross
	cell = Position{2, 2}
	call(masterAddress, "Nonogram.CrossCell", cell, &reply)
}

func TestMarkNonEmptyCell(t *testing.T) {
	var reply bool
	// Fill
	var cell = Position{1, 1}
	call(masterAddress, "Nonogram.FillCell", cell, &reply)
	call(masterAddress, "Nonogram.CrossCell", cell, &reply)
	// Cross
	cell = Position{2, 2}
	call(masterAddress, "Nonogram.CrossCell", cell, &reply)
	call(masterAddress, "Nonogram.FillCell", cell, &reply)
}

func TestMain(m *testing.M) {
	// Set up
	puzzle := new(Nonogram)
	startNonogramMaster(puzzle, masterAddress)

	// Run tests
	os.Exit(m.Run())
}
