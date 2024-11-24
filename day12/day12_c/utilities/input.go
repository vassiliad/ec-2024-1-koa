package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}
type Puzzle struct {
	Segments      map[int]Point
	Targets       []Point
	Width, Height int
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	segments := map[int]Point{
		1: {X: 0, Y: 0},
		2: {X: 0, Y: 1},
		3: {X: 0, Y: 2},
	}
	targets := []Point{}
	width := 0
	height := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")
		x, err := strconv.Atoi(parts[0])

		if err != nil {
			return &Puzzle{
				Segments: segments,
				Targets:  targets,
				Width:    width,
				Height:   height,
			}, err
		}

		y, err := strconv.Atoi(parts[1])

		if err != nil {
			return &Puzzle{
				Segments: segments,
				Targets:  targets,
				Width:    width,
				Height:   height,
			}, err
		}
		targets = append(targets, Point{X: x, Y: y})

		width = max(width, x)
		height = max(height, y)
	}

	return &Puzzle{
		Segments: segments,
		Targets:  targets,
		Width:    width,
		Height:   height,
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
