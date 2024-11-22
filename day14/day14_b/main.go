package main

import (
	"os"

	"puzzle/utilities"
)

type Point struct {
	x, y, z int
}

func (p Point) Add(o Point) Point {
	return Point{x: p.x + o.x, y: p.y + o.y, z: p.z + o.z}
}

func solution(input *utilities.Puzzle) int {
	deltas := map[int]Point{
		utilities.Up:       {x: 0, y: 1, z: 0},
		utilities.Down:     {x: 0, y: -1, z: 0},
		utilities.Right:    {x: 1, y: 0, z: 0},
		utilities.Left:     {x: -1, y: 0, z: 0},
		utilities.Forward:  {x: 0, y: 0, z: 1},
		utilities.Backward: {x: 0, y: 0, z: -1},
	}

	all_points := map[Point]int{}

	for _, branch := range input.Moves {
		pos := Point{}

		for _, move := range branch {

			d := deltas[move.Direction]

			for _ = range move.Repeat {
				pos = pos.Add(d)
				all_points[pos] = 1
			}

		}
	}

	return len(all_points)
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	// logger.Println("Input was", input)

	if err != nil {
		logger.Fatalln("Ran into problems while reading input. Problem", err)
	}

	sol := solution(input)

	logger.Println("Solution is", sol)
}
