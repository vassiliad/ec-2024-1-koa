package utilities

import (
	"bufio"
	"os"
	"strings"
)

type Point struct {
	X, Y, Strength int
}
type Puzzle struct {
	Segments map[int]Point
	Targets  []Point
	Width    int
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	segments := map[int]Point{}
	targets := []Point{}
	width := 0

	height := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		for idx, c := range line {
			if c == rune('T') {
				targets = append(targets, Point{X: idx, Y: height, Strength: 1})
			} else if c == rune('H') {
				targets = append(targets, Point{X: idx, Y: height, Strength: 2})
			} else if c >= rune('A') && c <= rune('Z') {
				segments[int(c-rune('A'))+1] = Point{X: idx, Y: height}
			}
			width = idx
		}

		height += 1
	}

	for k := range segments {
		x := segments[k]
		x.Y = height - x.Y

		segments[k] = x
	}

	for i := range targets {
		x := targets[i]
		x.Y = height - x.Y

		targets[i] = x
	}

	return &Puzzle{
		Segments: segments,
		Targets:  targets,
		Width:    width,
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
