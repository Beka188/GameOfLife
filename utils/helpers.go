package utils

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
