package main

import (
	"container/heap"
	"os"

	"image"
	"puzzle/utilities"
)

var deltas = []image.Point{
	{X: 1, Y: 0},
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
}

func Neighbours(p image.Point, puzzle *utilities.Puzzle) []image.Point {
	ret := []image.Point{}

	for _, off := range deltas {
		pos := p.Add(off)

		if (pos.X < 0 || pos.X >= puzzle.Width || pos.Y < 0 || pos.Y >= puzzle.Height) || puzzle.Board[pos.Y][pos.X] == rune('#') {
			continue
		}

		ret = append(ret, pos)

	}

	return ret
}

func solution(input *utilities.Puzzle) int {
	all_start := []image.Point{}
	end := image.Point{}

	// VV: We're basically searching from multiple starting points at the same time
	// the first one that reaches the end wins
	for y, row := range input.Board {
		for x, c := range row {
			if c == rune('S') {
				row[x] = rune('0')
				all_start = append(all_start, image.Point{X: x, Y: y})
			} else if c == rune('E') {
				row[x] = rune('0')
				end = image.Point{X: x, Y: y}
			}
		}
	}

	return find_path(all_start, end, input)
}

func trueCost(start, end image.Point, puzzle *utilities.Puzzle) int {
	val_start := int(puzzle.Board[start.Y][start.X] - rune('0'))
	val_end := int(puzzle.Board[end.Y][end.X] - rune('0'))

	// VV: You can either :

	// VV: 1. move directly towards the platform (i.e. if its lower than current platform then push lever down)
	lower := val_start - val_end

	if lower < 0 {
		lower = -lower
	}

	// VV: 2. go the opposite way (i.e. if it's lower than current platform then push lever up) and wrap around
	raise := 10 - lower

	val := 1 + min(lower, raise)

	return val
}

func find_path(start []image.Point, end image.Point, puzzle *utilities.Puzzle) int {
	g_score := map[image.Point]int{}
	reverse_path := map[image.Point]image.Point{}
	open := make(utilities.PriorityQueue, len(start))

	for idx, s := range start {
		g_score[s] = 0
		open[idx] = &utilities.HeapItem{
			Value:    &s,
			Priority: 0,
		}
	}

	heap.Init(&open)

	for open.Len() > 0 {

		heap_item := heap.Pop(&open)

		item := heap_item.(*utilities.HeapItem)

		cur := item.Value.(*image.Point)

		if *cur == end {
			return g_score[*cur]
		}

		neighbours := Neighbours(*cur, puzzle)
		cur_g_score := g_score[*cur]

		for _, next := range neighbours {
			next_cost := cur_g_score + trueCost(*cur, next, puzzle)

			if known_g, ok := g_score[next]; ok && known_g < next_cost {
				continue
			} else {

				g_score[next] = next_cost
				reverse_path[next] = *cur

				updated := false

				for idx, q := range open {
					if *q.Value.(*image.Point) == next {
						updated = true

						q.Priority = next_cost
						heap.Fix(&open, idx)
					}
				}

				if !updated {
					t := &utilities.HeapItem{
						Value:    &next,
						Priority: next_cost,
					}

					heap.Push(&open, t)
				}
			}
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