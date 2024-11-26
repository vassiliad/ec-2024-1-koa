package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	Spins  []int
	Wheels [][][]rune
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	wheels := [][][]rune{}
	spins := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, " ")

		if len(line) == 0 {
			continue
		}

		if len(spins) == 0 {

			for _, tok := range strings.Split(line, ",") {
				spin, err := strconv.Atoi(tok)

				if err != nil {
					return &Puzzle{
						Wheels: wheels,
						Spins:  spins,
					}, fmt.Errorf("invalid spin number, underlying error %+v", err)
				}

				spins = append(spins, spin)
				wheels = append(wheels, [][]rune{})
			}
		} else {
			chars := []rune(line)

			for i := 0; i < len(chars); i += 4 {
				if chars[i] == rune(' ') {
					continue
				}
				wheels[i/4] = append(wheels[i/4], chars[i:i+3])
			}
		}

	}

	return &Puzzle{
		Wheels: wheels,
		Spins:  spins,
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
