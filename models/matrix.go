package models

import (
	"bufio"
	"crunch03/globals"
	"errors"
	"fmt"
	"golang.org/x/term"
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
			m.Body[i][j].Live = newMatrix[i][j].Live
			if m.Body[i][j].Live {
				m.Body[i][j].IsVisited = true
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
	x, y := 0, 0
	if globals.IsFullScreen {
		y, x, _ = term.GetSize(int(os.Stdin.Fd()))
	}
	y = y / 2

	fmt.Println("Enter the grid:")

	scanner := bufio.NewScanner(os.Stdin)
	body := make([][]Cell, max(rows, x))

	for i := 0; i < max(rows, x); i++ {
		body[i] = make([]Cell, max(cols, y))
		if i >= rows { // fullscreen
			continue
		}
		if !scanner.Scan() {
			return nil, errors.New("not enough rows in input")
		}
		line := strings.TrimSpace(scanner.Text())
		if len(line) != cols {
			return nil, fmt.Errorf("number of columns in row %d (%d) does not match specified size (%d)", i, len(line), cols)
		}
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

// NewMatrix creates a new Matrix from user's input or random according to global X and Y constants
func NewMatrix(isRandom bool) (*Matrix, error) {
	var body [][]Cell
	fmt.Println("dfSDF", isRandom)
	if !isRandom {
		fmt.Println("SDF")
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
		body = generateRandomGrid(globals.RandomX, globals.RandomY)
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

func generateRandomGrid(rows, cols int) [][]Cell {
	x, y := 0, 0
	if globals.IsFullScreen {
		y, x, _ = term.GetSize(int(os.Stdin.Fd()))
	}
	y = y / 2
	body := make([][]Cell, max(rows, x))
	for i := 0; i < max(rows, x); i++ {
		body[i] = make([]Cell, max(cols, y))
		if i >= rows {
			continue
		}
		for j := 0; j < cols; j++ {
			r := rand.Intn(2)
			if r == 1 {
				body[i][j].Live = true
				body[i][j].IsVisited = true
			}
		}
	}
	fmt.Printf("size  %d %d\n", len(body), len(body[0]))
	return body
}
