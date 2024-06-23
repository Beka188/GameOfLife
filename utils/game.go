package utils

import (
	"crunch03/globals"
	"crunch03/models"
)

func Move(m *models.Board) {
	newMatrix := make([][]models.Cell, len(m.Body))
	for i := range newMatrix {
		newMatrix[i] = make([]models.Cell, len(m.Body[0]))
	}
	for i, row := range m.Body {
		for j, _ := range row {
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
		for j, _ := range row {
			m.Body[i][j].Live = newMatrix[i][j].Live
			if m.Body[i][j].Live {
				m.Body[i][j].IsVisited = true
				m.LiveCells++
			}
		}
	}
	m.TickCount++
}

func liveNeighbors(m [][]models.Cell, x, y int) (count int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			a, b := i, j
			if globals.IsEdgePortal {
				if i < 0 {
					a = len(m) - 1
				} else if i >= len(m) {
					a = 0
				}
				if j < 0 {
					b = len(m[0]) - 1
				} else if j >= len(m[0]) {
					b = 0
				}
			} else {
				if i < 0 || i >= len(m) || j < 0 || j >= len(m[0]) {
					continue
				}
			}
			if m[a][b].Live {
				count++
			}
		}
	}
	return
}
