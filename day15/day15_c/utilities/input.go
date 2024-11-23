package utilities

import (
	"bufio"
	"os"
	"strings"
)

type Puzzle struct {
	Board        [][]rune
	TargetFruits int
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	board := [][]rune{}
	target_fruits := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		row := []rune(line)
		board = append(board, row)
		for _, x := range row {
			if x != rune('#') && x != rune('~') && x != rune('.') {

				target_fruits |= 1 << (x - rune('A'))
			}
		}
	}

	return &Puzzle{
		Board:        board,
		TargetFruits: target_fruits,
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