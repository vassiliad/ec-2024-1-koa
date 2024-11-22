package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Up = iota
	Down
	Right
	Left
	Forward
	Backward
)

type Move struct {
	Direction int
	Repeat    int
}
type Puzzle struct {
	Moves []Move
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	moves := []Move{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		dir := -1

		for _, step := range strings.Split(line, ",") {
			parts := []rune(step)

			switch parts[0] {
			case rune('U'):
				dir = Up
			case rune('D'):
				dir = Down
			case rune('R'):
				dir = Right
			case rune('L'):
				dir = Left
			case rune('F'):
				dir = Forward
			case rune('B'):
				dir = Backward
			default:
				return &Puzzle{Moves: moves}, fmt.Errorf("invalid step %s", step)
			}

			val, err := strconv.Atoi(string(parts[1:]))

			if err != nil {
				return &Puzzle{Moves: moves}, fmt.Errorf("invalid step %s due to %+v", step, err)
			}

			moves = append(moves, Move{Direction: dir, Repeat: val})
		}
	}

	return &Puzzle{
		Moves: moves,
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
