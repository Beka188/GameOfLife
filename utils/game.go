package utils

import (
	"crunch03/globals"
	"crunch03/models"
	"time"
)

func Setup() *models.Board {
	isFlagsValid := readFlags()
	m, err := newMatrix()
	if err != nil || !isFlagsValid {
		panic(err)
	}
	return m
}

func StartGame(m *models.Board) {
	printMatrix(*m)
	ticker := time.NewTicker(globals.Interval)
	defer ticker.Stop()
	if m.LiveCells == 0 {
		return
	}
	for range time.Tick(globals.Interval) {
		move(m)
		printMatrix(*m)
		if m.LiveCells == 0 {
			break
		}
	}
}

func move(m *models.Board) {
	newMatrix := make([][]models.Cell, len(m.Body))
	for i := range newMatrix {
		newMatrix[i] = make([]models.Cell, len(m.Body[0]))
	}
	for i, row := range m.Body {
		for j := range row {
			liveNeighborsCount := liveNeighbors(m.Body, i, j)
			if m.Body[i][j].Live {
				if liveNeighborsCount < 2 {
					newMatrix[i][j].Live = false
				} else if liveNeighborsCount <= 3 {
					newMatrix[i][j].Live = true
				} else {
					newMatrix[i][j].Live = false
				}
			} else {
				if liveNeighborsCount == 3 {
					newMatrix[i][j].Live = true
				}
			}
		}
	}
	m.LiveCells = 0
	for i, row := range m.Body {
		for j := range row {
			m.Body[i][j].Live = newMatrix[i][j].Live
			if m.Body[i][j].Live {
				m.Body[i][j].IsVisited = true
				m.LiveCells++
			}
		}
	}
	m.TickCount++
}
