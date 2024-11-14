package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	Priests int
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	priests := -1

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		var e error
		priests, e = strconv.Atoi(line)

		if e != nil {
			return nil, e
		}

		break
	}

	if priests < 1 {
		return &Puzzle{Priests: priests}, fmt.Errorf("illegal number of priests %d", priests)
	}

	return &Puzzle{
		Priests: priests,
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
