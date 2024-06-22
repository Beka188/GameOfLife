package models

import "crunch03/globals"

func liveNeighbors(m [][]Cell, x, y int) (count int) {
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
