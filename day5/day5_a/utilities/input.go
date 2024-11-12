package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	Queues [][]int
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	queues := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		for idx, tok := range strings.Split(line, " ") {

			number, err := strconv.Atoi(tok)

			if err != nil {
				return &Puzzle{
					Queues: queues,
				}, err
			}

			if len(queues) <= idx {
				queues = append(queues, make([]int, 0))
			}

			queues[idx] = append(queues[idx], number)
		}

	}

	return &Puzzle{
		Queues: queues,
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
