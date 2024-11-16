package utilities

import (
	"bufio"
	"os"
	"strings"
)

type Puzzle struct {
	Words []string
	Text  string
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

	text := ""

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimLeft(line, " ")

		if len(line) == 0 {
			continue
		}

		text += line
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
