package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Puzzle struct {
	Blocks []int
	Width  int
	Height int
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	blocks := []int{}
	width := -1

	first_line := true

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if first_line {
			width = len(line)
		}

		for _, ch := range line {
			if ch == '.' {
				blocks = append(blocks, -1)
			} else if ch == '#' {
				blocks = append(blocks, 0)
			} else {
				return &Puzzle{
					Blocks: blocks,
					Width:  width,
					Height: len(blocks) / width,
				}, fmt.Errorf("invalid character %c", ch)
			}
		}
	}

	return &Puzzle{
		Blocks: blocks,
		Width:  width,
		Height: len(blocks) / width,
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
