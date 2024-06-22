package models

type cell struct {
	live bool
}

type matrix struct {
	body      [][]cell
	size      int
	tickCount int
	liveCells int
	delayMs   int
	isVerbose bool
	isDelayMs bool
}
