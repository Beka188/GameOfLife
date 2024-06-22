package main

import "fmt"

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
