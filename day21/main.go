package main

import (
	"image"
	"os"

	"puzzle/utilities"
)

func rotate[K any](p image.Point, dir rune, board [][]K) {
	if p.Y == 0 || p.Y == len(board)-1 || p.X == 0 && p.X == len(board[p.Y])-1 {
		// VV: Not a rotation point
		return
	}

	letters := make([]K, 8)

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
}

func clone(src [][]image.Point) [][]image.Point {
	the_clone := make([][]image.Point, len(src))

	for y := 0; y < len(src); y++ {
		the_clone[y] = make([]image.Point, len(src[y]))
		copy(the_clone[y], src[y])
	}

	return the_clone
}

func solution(input *utilities.Puzzle, max_rounds int) string {

	// VV: Run 1 full rotation and record the final location of the starting indices
	one := make([][]image.Point, len(input.Board))

	for y := 0; y < len(input.Board); y++ {
		one[y] = make([]image.Point, len(input.Board[y]))

		for x := 0; x < len(input.Board[y]); x++ {
			one[y][x] = image.Pt(x, y)
		}
	}

	rotations := 0
	for y := 1; y < len(one)-1; y++ {
		for x := 1; x < len(one[y])-1; x++ {
			rotate(image.Pt(x, y), input.Directions[rotations%len(input.Directions)], one)
			rotations++
		}
	}

	// VV: If we apply "one" on itself, we'll get 2-rounds worth of permutations -> "two"
	// apply "two" on "one" and you get "three", apply "two" on "two" and you get "four" ...
	// All you have to do is compute all the "powers of 2" that make up the number @max_rounds
	// and then "add them" i.e. apply one on the other to get your filan permutation table.
	// Apply the permutation on the original board and that's your decrypted message.
	p2_permutations := [][][]image.Point{one}

	last_power := one

	for max_power := 2; max_power <= max_rounds; max_power *= 2 {
		this_power := clone(last_power)

		for y := 0; y < len(this_power); y++ {
			for x := 0; x < len(this_power[y]); x++ {
				idx := last_power[y][x]
				this_power[y][x] = last_power[idx.Y][idx.X]
			}
		}
		p2_permutations = append(p2_permutations, this_power)

		last_power = this_power
	}

	current := make([][]image.Point, len(input.Board))

	for y := 0; y < len(input.Board); y++ {
		current[y] = make([]image.Point, len(input.Board[y]))

		for x := 0; x < len(input.Board[y]); x++ {
			current[y][x] = image.Pt(x, y)
		}
	}

	new_board := make([][]rune, len(input.Board))

	for y := 0; y < len(input.Board); y++ {
		new_board[y] = make([]rune, len(input.Board[y]))
	}

	// VV: once you have all the "powers of two" permutations then just
	// add them up to get "max_rounds" permutations and you're done.
	for i := range p2_permutations {
		rev_i := len(p2_permutations) - i - 1
		power_i := 1 << rev_i

		if power_i&max_rounds != 0 {
			swaps := p2_permutations[rev_i]

			for y := 0; y < len(input.Board); y++ {
				for x := 0; x < len(input.Board[y]); x++ {
					idx := swaps[y][x]
					new_board[y][x] = input.Board[idx.Y][idx.X]
				}
			}
			new_board, input.Board = input.Board, new_board

			// VV: quack
			for _, row := range input.Board {
				println(string(row))
			}
		}
	}

	ret := []rune{}

	for y := 0; y < len(input.Board); y++ {
		for x := 0; x < len(input.Board[y]); x++ {
			if input.Board[y][x] == '>' {
				for i := x + 1; input.Board[y][i] != '<'; i++ {
					ret = append(ret, input.Board[y][i])
				}

				return string(ret)
			}
		}
	}

	return "the end"
}

func main() {
	// https://www.reddit.com/r/everybodycodes/comments/1h2yc5e/2024_q21_message_to_the_knights/
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	// logger.Println("Input was", input)

	if err != nil {
		logger.Fatalln("Ran into problems while reading input. Problem", err)
	}

	sol := solution(input, input.Rounds)

	logger.Println("Solution is", sol)
}
