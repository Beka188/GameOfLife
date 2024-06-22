package models

import (
	"bufio"
	"crunch03/globals"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Matrix struct {
	Body      [][]Cell
	Size      int
	TickCount int
	LiveCells int
}

func Move(m *Matrix) {
	newMatrix := make([][]Cell, len(m.Body))
	for i := range newMatrix {
		newMatrix[i] = make([]Cell, len(m.Body[0]))
	}

	for i, row := range m.Body {
		for j, _ := range row {
			liveNeighborsCount := liveNeighbors(m.Body, i, j)
			fmt.Println(i, j, liveNeighborsCount)
			if m.Body[i][j].Live {
				if liveNeighborsCount < 2 {
					newMatrix[i][j].Live = false
				} else if liveNeighborsCount <= 3 {
					newMatrix[i][j].Live = true
				} else {
					newMatrix[i][j].Live = false
				}
			} else {
				if liveNeighborsCount == 3 {
					newMatrix[i][j].Live = true
				}
			}
		}
	}
	m.LiveCells = 0
	for i, row := range m.Body {
		for j, _ := range row {
			m.Body[i][j] = newMatrix[i][j]
			if m.Body[i][j].Live {
				m.LiveCells++
			}
		}
	}
	m.TickCount++
}

// promptSize asks for the size of the grid and works til user enters the correct format
func promptSize() (int, int, error) {
	var rows, cols int
	for {
		fmt.Print("Enter the size of the matrix (rows cols): ")
		if _, err := fmt.Scanf("%d %d", &rows, &cols); err != nil {
			fmt.Println("Invalid dimensions. Please enter two integers.")
			continue
		}
		if rows < 3 || cols < 3 {
			fmt.Println("Minimum size is 3x3. Please enter valid dimensions.")
			continue
		}
		break
	}
	return rows, cols, nil
}

// promptGrid asks for the grid input and validates it
func promptGrid(rows, cols int) ([][]Cell, error) {
	fmt.Println("Enter the grid:")

	scanner := bufio.NewScanner(os.Stdin)
	body := make([][]Cell, rows)

	for i := 0; i < rows; i++ {
		if !scanner.Scan() {
			return nil, errors.New("not enough rows in input")
		}
		line := strings.TrimSpace(scanner.Text())
		if len(line) != cols {
			return nil, fmt.Errorf("number of columns in row %d (%d) does not match specified size (%d)", i, len(line), cols)
		}
		body[i] = make([]Cell, cols)
		for j, char := range line {
			switch char {
			case '#':
				body[i][j].Live = true
			case '.':
				body[i][j].Live = false
			default:
				return nil, fmt.Errorf("invalid character in input at row %d, column %d: %c", i, j, char)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}

	return body, nil
}

// NewMatrix creates a new Matrix from user's input
func NewMatrix(isRandom bool) (*Matrix, error) {
	var body [][]Cell
	if !isRandom {
		rows, cols, err := promptSize()
		if err != nil {
			return nil, err
		}

		for {
			body, err = promptGrid(rows, cols)
			if err != nil {
				fmt.Println("Invalid grid. Please re-enter the grid:")
				continue
			}
			break
		}
	} else {
		body = make([][]Cell, globals.RandomX)
		for i := 0; i < globals.RandomX; i++ {
			body[i] = make([]Cell, globals.RandomY)

			for j := 0; j < globals.RandomY; j++ {
				r := rand.Intn(2)
				if r == 1 {
					body[i][j].Live = true
				} else {
					body[i][j].Live = false
				}
			}
		}

	}

	liveCells := 0
	for _, row := range body {
		for _, cell := range row {
			if cell.Live {
				liveCells++
			}
		}
	}
	return &Matrix{
		Body:      body,
		TickCount: 1,
		LiveCells: liveCells,
	}, nil
}
