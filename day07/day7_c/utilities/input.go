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

type Point struct {
	X, Y int
}

func ParseTerrain(terrain_buffer [][]rune) []string {
	h := len(terrain_buffer)

	terrain := []string{}

	start := Point{X: 0, Y: 0}
	visited := map[Point]int{}

	cur := start

	// VV: First, try to go right or down and if that fails then try moves towards the top left corner
	// i.e. up and left
	deltas := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}

	for {
		next := Point(start)

		for _, d := range deltas {
			x := cur.X + d[0]
			y := cur.Y + d[1]

			if y < 0 || y >= h {
				continue
			}

			w := len(terrain_buffer[y])

			if x < 0 || x >= w {
				continue
			}

			next.X = x
			next.Y = y

			if terrain_buffer[y][x] == rune('S') {
				break
			} else if terrain_buffer[y][x] == rune(' ') {
				continue
			}

			if _, been_there := visited[next]; !been_there {
				break
			}
		}

		if terrain_buffer[next.Y][next.X] == rune('S') {
			terrain = append(terrain, "S")
			break
		}

		visited[next] = 1

		// print(fmt.Sprintf("%d, %d : %c\n", next.Y, next.X, terrain_buffer[next.Y][next.X]))

		terrain = append(terrain, fmt.Sprintf("%c", terrain_buffer[next.Y][next.X]))

		cur = next

		// if len(terrain) == 20 {
		// 	break
		// }
	}

	return terrain
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

	terrain := ParseTerrain(terrain_buffer)

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
