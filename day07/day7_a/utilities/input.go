package utilities

import (
	"bufio"
	"os"
	"strings"
)

type Plan struct {
	Name    string
	Actions []string
}
type Puzzle struct {
	Plans []Plan
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	plans := []Plan{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, ":")
		name := tokens[0]
		actions := strings.Split(tokens[1], ",")

		plans = append(plans, Plan{Name: name, Actions: actions})
	}

	return &Puzzle{
		Plans: plans,
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
