package main

import (
	"crunch03/globals"
	"crunch03/models"
	"crunch03/utils"
	"fmt"
	"time"
)

func main() {
	utils.ReadFlags()

	m, err := models.NewMatrix()
	if err != nil {
		panic(err)
	} else {
		printMatrix(*m)
	}

	ticker := time.NewTicker(globals.Interval)
	defer ticker.Stop()

	go func() {
		for _ = range ticker.C {
			models.Move(m)
			printMatrix(*m)
		}
	}()
	select {}
}

func printMatrix(m models.Matrix) {
	if m.IsVerbose {
		fmt.Printf("Tick: %d\nGrid Size: %dx%d\nLive Cells: %d\nDelayMs: %dms\n\n", m.TickCount, len(m.Body), len(m.Body[0]), m.LiveCells, globals.Interval)
	}

	for _, row := range m.Body {
		for j, cell := range row {
			if cell.Live {
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
	fmt.Printf("\n")

}
