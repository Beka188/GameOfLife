package models

type Matrix struct {
	Body      [][]Cell
	Size      int
	TickCount int
	LiveCells int
}

// NewMatrix creates a new Matrix from user's input or random according to global X and Y constants
