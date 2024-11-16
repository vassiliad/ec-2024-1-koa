package utilities

import (
	"bufio"
	"os"
	"strings"
)

type Puzzle struct {
	Grid [][]rune
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	grid := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		grid = append(grid, []rune(line))
	}

	return &Puzzle{
		Grid: grid,
	}, nil
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
