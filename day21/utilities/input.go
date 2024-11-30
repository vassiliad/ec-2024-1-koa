package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	Board      [][]rune
	Directions []rune
	Rounds     int
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	board := [][]rune{}
	directions := []rune{}
	rounds := 1024

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if len(directions) == 0 {
			parts := strings.Split(line, ":")
			directions = []rune(parts[1])
			var err error = nil

			if rounds, err = strconv.Atoi(parts[0]); err != nil {
				return &Puzzle{
					Directions: directions,
					Board:      board,
				}, err
			}
		} else {
			board = append(board, []rune(line))
		}
	}

	return &Puzzle{
		Directions: directions,
		Board:      board,
		Rounds:     rounds,
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
