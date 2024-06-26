package utils

import (
	"bufio"
	"crunch03/globals"
	"crunch03/models"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"golang.org/x/term"
)

func newMatrix() (*models.Board, error) {
	var body [][]models.Cell
	if globals.FileName != "" {
		var err error
		body, err = readInputFromFile()
		if err != nil {
			return nil, err
		}
	} else if !globals.IsRandom {
		rows, cols, err := promptSize()
		if err != nil {
			return nil, err
		}
		for {
			body, err = promptGrid(rows, cols)
			if err != nil {
				fmt.Println("Invalid grid. Please re-enter the grid: (. - for dead cell, # - for live cell) ")
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
	return &models.Board{
		Body:      body,
		TickCount: 1,
		LiveCells: liveCells,
	}, nil
}

func generateRandomGrid(rows, cols int) [][]models.Cell {
	x, y := 0, 0
	if globals.IsFullScreen {
		y, x, _ = term.GetSize(int(os.Stdin.Fd()))
	}
	y = y / 2
	body := make([][]models.Cell, customMax(rows, x))
	for i := 0; i < customMax(rows, x); i++ {
		body[i] = make([]models.Cell, customMax(cols, y))
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
	return body
}

// promptSize asks for the size of the grid and works til user enters the correct format
func promptSize() (int, int, error) {
	var rows, cols int
	for {
		fmt.Print("Enter the size of the matrix (rows cols): ")
		_, err := fmt.Scanf("%d %d", &rows, &cols)
		if err != nil {
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
func promptGrid(rows, cols int) ([][]models.Cell, error) {
	x, y := 0, 0
	if globals.IsFullScreen {
		y, x, _ = term.GetSize(int(os.Stdin.Fd()))
	}
	y = y / 2

	fmt.Println("Enter the grid: (. - for dead cell, # - for live cell) ")

	scanner := bufio.NewScanner(os.Stdin)
	body := make([][]models.Cell, customMax(rows, x))

	for i := 0; i < customMax(rows, x); i++ {
		body[i] = make([]models.Cell, customMax(cols, y))
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

// readInputFromFile reads grid size and grid data from a file
func readInputFromFile() ([][]models.Cell, error) {
	x, y := 0, 0
	if globals.IsFullScreen {
		y, x, _ = term.GetSize(int(os.Stdin.Fd()))
	}
	y = y / 2

	content, err := os.ReadFile(globals.FileName)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	cont := strings.Split(string(content), "\n")
	sizeLine := strings.TrimSpace(cont[0])
	var rows, cols int
	if _, err := fmt.Sscanf(sizeLine, "%d %d", &rows, &cols); err != nil {
		return nil, fmt.Errorf("invalid size format in file: %w", err)
	}
	body := make([][]models.Cell, customMax(x, rows))
	for i := 0; i < customMax(rows, x); i++ {
		body[i] = make([]models.Cell, customMax(cols, y))
		if i >= rows {
			continue
		}
		line := strings.TrimSpace(cont[i+1])
		if len(line) < cols {
			return nil, fmt.Errorf("number of columns in row %d (%d) does not match specified size (%d)", i, len(line), cols)
		}
		for j := 0; j < customMax(cols, y); j++ {
			if j >= len(line) {
				continue
			}
			char := line[j]
			switch char {
			case '#':
				body[i][j].Live = true
				body[i][j].IsVisited = true
			case '.':
				body[i][j].Live = false
			default:
				return nil, fmt.Errorf("invalid character in input at row %d, column %d: %c", i, j, char)
			}
		}
	}

	fmt.Printf("%d   %d\n", len(body), len(body[0]))
	return body, nil
}

func customMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
