package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Plan struct {
	Name    string
	Actions []string
}
type Puzzle struct {
	Plans   []Plan
	Terrain []string
}

func ReadScanner(scanner *bufio.Scanner) (*Puzzle, error) {
	plans := []Plan{}
	terrain_buffer := [][]rune{}

	parsing_terrain := false
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "S") {
			parsing_terrain = true
		}

		if !parsing_terrain {
			tokens := strings.Split(line, ":")
			name := tokens[0]
			actions := strings.Split(tokens[1], ",")

			plans = append(plans, Plan{Name: name, Actions: actions})
		} else {
			terrain_buffer = append(terrain_buffer, []rune(line))
		}

	}

	w := len(terrain_buffer[0])
	lines := len(terrain_buffer)

	// VV: First and last line contain W segments, the rest just 2
	terrain_size := (lines-2)*2 + w*2

	terrain := make([]string, terrain_size)

	for i := range w {
		terrain[i] = fmt.Sprintf("%c", terrain_buffer[0][i])
		// VV: The last line is reversed
		terrain[w+(lines-2)+i] = fmt.Sprintf("%c", terrain_buffer[lines-1][w-i-1])
	}

	for l := range lines - 2 {
		terrain[w+l] = fmt.Sprintf("%c", terrain_buffer[l+1][w-1])
		terrain[w+(lines-2)+w+(lines-2-l-1)] = fmt.Sprintf("%c", terrain_buffer[l+1][0])
	}

	return &Puzzle{
		Plans:   plans,
		Terrain: terrain,
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
