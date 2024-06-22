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
