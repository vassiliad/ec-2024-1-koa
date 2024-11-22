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

	pos := Point{}

	max_y := 0

	for _, move := range input.Moves {

		d := deltas[move.Direction]

		delta := Point{x: move.Repeat * d.x, y: move.Repeat * d.y, z: move.Repeat * d.z}
		new_pos := pos.Add(delta)

		pos = new_pos
		max_y = max(max_y, new_pos.y)
	}

	return max_y
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
