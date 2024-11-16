package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	AvailableBlocks int
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	available_blocks := -1

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		var e error
		available_blocks, e = strconv.Atoi(line)

		if e != nil {
			return nil, e
		}

		break
	}

	if available_blocks < 1 {
		return &Puzzle{AvailableBlocks: available_blocks}, fmt.Errorf("illegal number of blocks %d", available_blocks)
	}

	return &Puzzle{
		AvailableBlocks: available_blocks,
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
