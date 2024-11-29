package utilities

import (
	"bufio"
	"os"
	"strings"
)

type Puzzle struct {
	Board      [][]rune
	Directions []rune
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	board := [][]rune{}
	directions := []rune{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if len(directions) == 0 {
			directions = []rune(line)
		} else {
			board = append(board, []rune(line))
		}
	}

	return &Puzzle{
		Directions: directions,
		Board:      board,
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
