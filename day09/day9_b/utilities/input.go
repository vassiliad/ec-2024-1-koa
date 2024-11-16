package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	Brightnesses []int
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	brightnesses := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		b, e := strconv.Atoi(line)

		if e != nil {
			return nil, e
		}

		brightnesses = append(brightnesses, b)

		if b < 0 {
			return &Puzzle{
				Brightnesses: brightnesses,
			}, fmt.Errorf("negative brightness %d", b)
		}

	}

	return &Puzzle{
		Brightnesses: brightnesses,
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
