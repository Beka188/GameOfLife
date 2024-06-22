package main

import (
	"crunch03/globals"
	"crunch03/models"
	"crunch03/utils"
	"fmt"
	"time"
)

func main() {
	isFlagsValid := utils.ReadFlags()
	if !isFlagsValid {
		return
	}
	m, err := models.NewMatrix(globals.RandomX != 0 && globals.RandomY != 0) // globals.RandomX != 0 || globals.RandomY != 0
	if err != nil {
		panic(err)
	} else {
		printMatrix(*m)
	}
	ticker := time.NewTicker(globals.Interval)
	defer ticker.Stop()

	startGame(m)
}

func printMatrix(m models.Matrix) {
	if globals.IsVerbose {
		fmt.Printf("Tick: %d\nGrid Size: %dx%d\nLive Cells: %d\nDelayMs: %v\n\n", m.TickCount, len(m.Body), len(m.Body[0]), m.LiveCells, globals.Interval)
	}

	for _, row := range m.Body {
		for j, cell := range row {
			if cell.Live {
				fmt.Printf("x")
			} else if cell.IsVisited {
				fmt.Printf("∘")
			} else {
				fmt.Printf(".")
			}
			if j != len(row)-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

}

func startGame(m *models.Matrix) {
	for _ = range time.Tick(globals.Interval) {
		models.Move(m)
		printMatrix(*m)
		if m.LiveCells == 0 {
			break
		}
	}
}
