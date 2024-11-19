package utilities

import (
	"bufio"
	"os"
	"strings"
)

type Puzzle struct {
	Recipe map[string][]string
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	recipe := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ":")
		mutations := strings.Split(parts[1], ",")

		recipe[parts[0]] = mutations
	}

	return &Puzzle{
		Recipe: recipe,
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
