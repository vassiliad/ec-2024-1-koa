package main

import (
	"image"
	"os"

	"puzzle/utilities"
)

func rotate(p image.Point, dir rune, board [][]rune) {
	// VV: Not a rotation point

	if p.Y == 0 || p.Y == len(board)-1 || p.X == 0 && p.X == len(board[p.Y])-1 {
		return
	}

	letters := make([]rune, 8)

	offset := 1

	if dir == 'L' {
		// VV: Rotating once to the left is the same as rotating 7 to the right
		offset = (8 - 1) % 8
	}

	deltas := []image.Point{
		{X: -1, Y: -1},
		{X: 0, Y: -1},
		{X: 1, Y: -1},
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 0, Y: 1},
		{X: -1, Y: 1},
		{X: -1, Y: 0},
	}

	for idx, delta := range deltas {
		pos := p.Add(delta)
		letters[(offset+idx)%8] = board[pos.Y][pos.X]
	}

	for idx, delta := range deltas {
		pos := p.Add(delta)
		board[pos.Y][pos.X] = letters[idx]
	}

	// panic(string(letters))

}

func solution(input *utilities.Puzzle) string {

	rotations := 0

	for y := 1; y < len(input.Board)-1; y++ {
		for x := 1; x < len(input.Board[y])-1; x++ {
			rotate(image.Pt(x, y), input.Directions[rotations%len(input.Directions)], input.Board)
			rotations++

			// for _, row := range input.Board {
			// 	println(string(row))
			// }
		}
	}

	ret := []rune{}

	for y := 0; y < len(input.Board); y++ {
		for x := 0; x < len(input.Board[y]); x++ {
			if input.Board[y][x] == '>' {
				for i := x + 1; input.Board[y][i] != '<'; i++ {
					ret = append(ret, input.Board[y][i])
				}
			}
		}

		if len(ret) > 0 {
			return string(ret)
		}
	}

	panic("Oh no")
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
