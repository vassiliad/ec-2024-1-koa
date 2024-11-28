package utilities

import (
	"bufio"
	"image"
	"os"
	"strings"
)

type Puzzle struct {
	Stars []image.Point
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	stars := []image.Point{}
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, " ")

		if len(line) == 0 {
			continue
		}

		y++

		for x, c := range []rune(line) {
			if c == '*' {
				stars = append(stars, image.Pt(x+1, y))
			}
		}
	}

	// VV: Rewrite stars so that bottom left corner is 1, 1
	for i := range stars {
		stars[i].Y = (y - stars[i].Y) + 1
	}

	return &Puzzle{
		Stars: stars,
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
