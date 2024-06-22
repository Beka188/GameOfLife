package utils

import (
	"bufio"
	"crunch03/globals"
	"crunch03/models"
	"errors"
	"fmt"
	"golang.org/x/term"
	"math/rand"
	"os"
	"strings"
)

func NewMatrix(isRandom bool) (*models.Matrix, error) {
	var body [][]models.Cell
	fmt.Println("dfSDF", isRandom)
	if globals.FileName != "" {
		var err error
		body, err = readInputFromFile()
		if err != nil {
			return nil, err
		}
	} else if !isRandom {
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
	return &models.Matrix{
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
	body := make([][]models.Cell, max(rows, x))
	for i := 0; i < max(rows, x); i++ {
		body[i] = make([]models.Cell, max(cols, y))
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
func promptGrid(rows, cols int) ([][]models.Cell, error) {
	x, y := 0, 0
	if globals.IsFullScreen {
		y, x, _ = term.GetSize(int(os.Stdin.Fd()))
	}
	y = y / 2

	fmt.Println("Enter the grid:")

	scanner := bufio.NewScanner(os.Stdin)
	body := make([][]models.Cell, max(rows, x))

	for i := 0; i < max(rows, x); i++ {
		body[i] = make([]models.Cell, max(cols, y))
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
	// Read the grid data
	body := make([][]models.Cell, rows)
	for i := 0; i < rows && i <= len(cont); i++ {
		line := strings.TrimSpace(cont[i+1])
		if len(line) != cols {
			return nil, fmt.Errorf("number of columns in row %d (%d) does not match specified size (%d)", i, len(line), cols)
		}
		body[i] = make([]models.Cell, cols)
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
	return body, nil
}
