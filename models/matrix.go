package models

type Board struct {
	Body      [][]Cell
	Size      int
	TickCount int
	LiveCells int
}
