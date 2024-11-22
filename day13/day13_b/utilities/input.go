package utilities

import (
	"bufio"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}
type Puzzle struct {
	Board         [][]rune
	Width, Height int
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	board := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		row := []rune(line)
		board = append(board, row)
	}

	return &Puzzle{
		Board:  board,
		Width:  len(board[0]),
		Height: len(board),
	}, scanner.Err()
}

func ReadString(text string) (*Puzzle, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) (*Puzzle, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
