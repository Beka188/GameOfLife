package utils

import (
	"crunch03/globals"
	"crunch03/models"
	"fmt"
)

func printMatrix(m models.Board) {
	if globals.IsVerbose {
		fmt.Printf("\nTick: %d\nGrid Size: %dx%d\nLive Cells: %d\nDelayMs: %v", m.TickCount, len(m.Body), len(m.Body[0]), m.LiveCells, globals.Interval)
	}
	fmt.Printf("\n\n")
	for i, row := range m.Body {
		for j, cell := range row {
			if cell.Live {
				if !globals.IsColored {
					fmt.Printf("x")
				} else {
					fmt.Print(globals.ColorLive + "x" + globals.ResetColor)
				}
			} else if cell.IsVisited && globals.IsFootPrint {
				if !globals.IsColored {
					fmt.Printf("∘")
				} else {
					fmt.Printf(globals.ColorFootprint + "∘" + globals.ResetColor)
				}
			} else {
				if !globals.IsColored {
					fmt.Printf(".")
				} else {
					fmt.Printf(globals.ColorEmpty + "." + globals.ResetColor)
				}
			}
			if j != len(row)-1 {
				fmt.Printf(" ")
			}
		}
		if i != len(m.Body)-1 {
			fmt.Printf("\n")
		}
	}
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
