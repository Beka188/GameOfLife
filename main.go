package main

import "fmt"

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

func (m *matrix) updateLiveNeighborsCount() {
	for i := 0; i < len(m.body); i++ {
		for j := 0; j < len(m.body[0]); j++ {
			if m.body[i][j].live {
				m.liveCells++
			}
		}
	}
}

func move(m *matrix) {
	newMatrix := make([][]cell, len(m.body))
	for i := range newMatrix {
		newMatrix[i] = make([]cell, len(m.body[0]))
	}

	for i, row := range m.body {
		for j, _ := range row {
			liveNeighborsCount := liveNeighbors(m.body, i, j)
			//fmt.Printf("%d %d   == > %d\n", i, j, liveNeighborsCount)
			if m.body[i][j].live {
				if liveNeighborsCount < 2 {
					newMatrix[i][j].live = false
				} else if liveNeighborsCount <= 3 {
					newMatrix[i][j].live = true
				} else {
					newMatrix[i][j].live = false
				}
			} else {
				if liveNeighborsCount == 3 {
					newMatrix[i][j].live = true
				}
			}
		}
	}
	m.liveCells = 0
	for i, row := range m.body {
		for j, _ := range row {
			m.body[i][j] = newMatrix[i][j]
			if m.body[i][j].live {
				m.liveCells++
			}
		}
	}
	m.tickCount++
}

func liveNeighbors(matrix [][]cell, x, y int) (count int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) {
				if matrix[i][j].live {
					count++
				}
			}
		}
	}
	return
}

func main() {
	arr := make([][]cell, 6)
	for i := range arr {
		arr[i] = make([]cell, 6)
	}
	arr[1][2].live = true
	arr[1][3].live = true
	arr[2][1].live = true
	arr[2][2].live = true
	arr[3][2].live = true
	arr[3][3].live = true
	arr[4][2].live = true
	arr[4][3].live = true
	m := matrix{
		body:      arr,
		size:      6,
		tickCount: 0,
		liveCells: 0,
		delayMs:   0,
		isVerbose: false,
		isDelayMs: false,
	}

	printMatrix(m.body)
	//fmt.Println()
	move(&m)
	printMatrix(m.body)
	//printMatrix(matrix)
	//
	//fmt.Println()
	//
	//matrix = move(matrix)
	//printMatrix(matrix)

}

func printMatrix(matrix [][]cell) {
	for _, row := range matrix {
		for j, cell := range row {
			if cell.live {
				fmt.Printf("x")
			} else {
				fmt.Printf(".")
			}
			if j != len(row)-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func (m *matrix) countLiveCells() {
	m.liveCells = 0
	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			if m.body[i][j].live {
				m.liveCells++
			}
		}
	}
}
