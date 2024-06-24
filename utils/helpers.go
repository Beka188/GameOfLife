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
	// fmt.Printf("\n")
}
