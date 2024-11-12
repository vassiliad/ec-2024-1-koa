package utilities

import (
	"bufio"
	"os"
	"strings"
)

type Puzzle struct {
	Branches map[string][]string
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	branches := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, ":")
		branch := tokens[0]
		leads_to := strings.Split(tokens[1], ",")

		branches[branch] = leads_to
	}

	return &Puzzle{
		Branches: branches,
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
