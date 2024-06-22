package models

func liveNeighbors(m [][]Cell, x, y int) (count int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if i >= 0 && i < len(m) && j >= 0 && j < len(m[0]) {
				if m[i][j].Live {
					count++
				}
			}
		}
	}
	return
}
