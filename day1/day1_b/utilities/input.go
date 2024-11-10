package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Enemy int

const (
	X = -1
	A = 0
	B = 1
	C = 3
	D = 5
)

func ReadScanner(scanner *bufio.Scanner) ([]Enemy, error) {
	ret := []Enemy{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		for _, ch := range line {
			switch ch := ch; ch {
			case 'x':
				ret = append(ret, X)
			case 'A':
				ret = append(ret, A)
			case 'B':
				ret = append(ret, B)
			case 'C':
				ret = append(ret, C)
			case 'D':
				ret = append(ret, D)
			default:
				return ret, fmt.Errorf("unexpected character %c", ch)
			}
		}

	}

	return ret, scanner.Err()
}

func ReadString(text string) ([]Enemy, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]Enemy, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
