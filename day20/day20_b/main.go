package main

import (
	"container/list"
	"fmt"
	"image"
	"os"

	"puzzle/utilities"
)

const (
	Up    = 0
	Down  = 1
	Left  = 2
	Right = 3
)

var Transitions = map[Direction][]Direction{
	Up:    {Left, Up, Right},
	Down:  {Right, Down, Left},
	Left:  {Down, Left, Up},
	Right: {Up, Right, Down},
}

var Deltas = map[Direction]image.Point{
	Up:    image.Pt(+0, -1),
	Down:  image.Pt(+0, +1),
	Left:  image.Pt(-1, +0),
	Right: image.Pt(+1, +0),
}

func can_move(pt image.Point, puzzle *utilities.Puzzle) bool {
	return pt.Y >= 0 && pt.Y < len(puzzle.Board) &&
		pt.X >= 0 && pt.X < len(puzzle.Board[pt.Y]) &&
		puzzle.Board[pt.Y][pt.X] != '#'
}

type Direction int

type State struct {
	dir  Direction
	pos  image.Point
	secs int
	// VV: Bitfield for the flags (checkpoints). A=1, B=2, C=4
	flags int
}

type Explore struct {
	height int
	State
}

func solution(input *utilities.Puzzle) int {
	const starting_height = 10000

	start := image.Point{}

	for y, row := range input.Board {
		for x, c := range row {
			if c == 'S' {
				start = image.Pt(x, y)
				// VV: Not sure whether this is correct or not, it likely doesn't matter
				// As the Glider would not end up entering this tile multiple times (worse than travelling elsewhere)
				input.Board[y][x] = '.'
				break
			}
		}
	}

	pending := list.New()
	best := map[State]int{}

	for dir := range Direction(4) {
		s := Explore{
			State: State{dir: dir, pos: start, secs: 0, flags: -1}, height: starting_height,
		}
		best[s.State] = starting_height

		pending.PushBack(s)
	}

	min_time := int(^uint(0) >> 1)

	steps := 0

	for pending.Front() != nil {
		f := pending.Front()
		pending.Remove(f)

		cur := f.Value.(Explore)

		cur.secs++

		steps++

		// VV: This took aaaaaaaaaaaaaaaaaaages to finish
		if steps%1_000_000 == 0 {
			fmt.Printf("%+v\n", cur)
		}

		for _, next_dir := range Transitions[cur.dir] {
			delta := Deltas[next_dir]
			pos := cur.pos.Add(delta)

			if !can_move(pos, input) {
				continue
			}

			next := cur

			next.dir = next_dir
			next.pos = pos

			if (input.Board[pos.Y][pos.X] >= 'A') && (input.Board[pos.Y][pos.X] <= 'C') {
				next_flag := int(input.Board[pos.Y][pos.X] - 'A')
				if next_flag != next.flags+1 {
					// VV: Got here either multiple times OR in the wrong order
					continue
				}

				next.flags++
			}

			switch input.Board[pos.Y][pos.X] {
			case '.', 'A', 'B', 'C':
				next.height -= 1
			case '-':
				next.height -= 2
			case '+':
				next.height += 1
			}

			if next.pos == start {
				if next.flags == 2 && next.height >= starting_height {
					return next.secs
				}

				if next.height < starting_height {
					// VV: Cliff edge is too high, acts like a #
					continue
				}
			}

			if best_height, ok := best[next.State]; ok && best_height >= next.height {
				continue
			}

			best[next.State] = next.height
			pending.PushBack(next)
		}
	}

	return min_time
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
