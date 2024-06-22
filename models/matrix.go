package models

import (
	"crunch03/globals"
	"fmt"
)

type Matrix struct {
	Body      [][]Cell
	Size      int
	TickCount int
	LiveCells int
}

func Move(m *Matrix) {
	newMatrix := make([][]Cell, len(m.Body))
	for i := range newMatrix {
		newMatrix[i] = make([]Cell, len(m.Body[0]))
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

// NewMatrix creates a new Matrix from user's input or random according to global X and Y constants
func NewMatrix(isRandom bool) (*Matrix, error) {
	var body [][]Cell
	fmt.Println("dfSDF", isRandom)
	if !isRandom {
		fmt.Println("SDF")
		rows, cols, err := promptSize()
		if err != nil {
			return nil, err
		}
		for {
			body, err = promptGrid(rows, cols)
			if err != nil {
				fmt.Println("Invalid grid. Please re-enter the grid:")
				continue
			}
			break
		}
	} else {
		body = generateRandomGrid(globals.RandomX, globals.RandomY)
	}
	liveCells := 0
	for _, row := range body {
		for _, cell := range row {
			if cell.Live {
				liveCells++
			}
		}
	}
	return &Matrix{
		Body:      body,
		TickCount: 1,
		LiveCells: liveCells,
	}, nil
}
