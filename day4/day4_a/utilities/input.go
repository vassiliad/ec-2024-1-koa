package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	Nails []int
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	nails := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		nail, err := strconv.Atoi(line)

		if err != nil {
			return &Puzzle{
				Nails: nails,
			}, err
		}

		nails = append(nails, nail)
	}

	return &Puzzle{
		Nails: nails,
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
