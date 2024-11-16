package utilities

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

type Puzzle struct {
	Words []string
	Text  []string
}

// From  https://github.com/golang/example/blob/master/hello/reverse/reverse.go
func reverse_string(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	words := []string{}
	first_line := true

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimLeft(line, " ")

		if len(line) == 0 {
			continue
		}

		if first_line {
			line = line[6:]
			first_line = false
		}

		words = append(words, strings.Split(line, ",")...)
		break
	}

	text := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		text = append(text, line)
	}

	for _, w := range words {
		rev_w := reverse_string(w)

		if !slices.Contains(words, rev_w) {
			words = append(words, rev_w)
		}
	}

	return &Puzzle{
		Words: words,
		Text:  text,
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
