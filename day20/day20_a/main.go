package main

import (
	"container/list"
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
}

type Explore struct {
	height int
	State
}

func solution(input *utilities.Puzzle) int {
	const max_time = 100

	max_height := 1000

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
			State: State{dir: dir, pos: start, secs: 0}, height: max_height,
		}
		best[s.State] = max_height

		pending.PushBack(s)
	}

	for pending.Front() != nil {
		f := pending.Front()
		pending.Remove(f)

		cur := f.Value.(Explore)

		cur.secs++

		for _, next_dir := range Transitions[cur.dir] {
			delta := Deltas[next_dir]
			pos := cur.pos.Add(delta)

			if !can_move(pos, input) {
				continue
			}

			next := cur

			next.dir = next_dir
			next.pos = pos

			switch input.Board[pos.Y][pos.X] {
			case '.':
				next.height--
			case '-':
				next.height -= 2
			case '+':
				next.height += 1
			}

			if best_height, ok := best[next.State]; ok && best_height >= next.height {
				continue
			}

			best[next.State] = next.height

			if next.secs == max_time {
				max_height = max(max_height, next.height)
				continue
			}

			pending.PushBack(next)
		}
	}

	return max_height
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
