package models

type Matrix struct {
	body      [][]Cell
	size      int
	tickCount int
	liveCells int
	delayMs   int
	isVerbose bool
	isDelayMs bool
}

func move(m *Matrix) {
	newMatrix := make([][]Cell, len(m.body))
	for i := range newMatrix {
		newMatrix[i] = make([]Cell, len(m.body[0]))
	}

	for i, row := range m.body {
		for j, _ := range row {
			liveNeighborsCount := liveNeighbors(m.body, i, j)
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
