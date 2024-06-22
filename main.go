package main

import (
	"crunch03/models"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
)

func main() {
	m, err := models.NewMatrix()
	if err != nil {
		panic(err)
	} else {
		printMatrix(*m)
	}
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyArrowRight {
			models.Move(m)
			printMatrix(*m)
		} else if char == 'q' || key == keyboard.KeyCtrlC {
			break
		}
	}
}

func printMatrix(m models.Matrix) {
	fmt.Printf("%d\n\n", m.TickCount)
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

//func (m *matrix) countLiveCells() {
//	m.liveCells = 0
//	for i := 0; i < m.size; i++ {
//		for j := 0; j < m.size; j++ {
//			if m.body[i][j].live {
//				m.liveCells++
//			}
//		}
//	}
//}
