package main

import (
	"container/heap"
	"os"

	"puzzle/utilities"
)

type Point struct {
	x, y       int
	fruits     int
	going_back bool
}

func (p *Point) Add(o Point) Point {

	return Point{x: p.x + o.x, y: p.y + o.y, fruits: p.fruits | o.fruits, going_back: p.going_back}
}

func calc_distance_to_any_fruit(start Point, puzzle *utilities.Puzzle) int {
	start.going_back = false
	g_score := map[Point]int{start: 0}

	open := make(utilities.PriorityQueue, 1)

	open[0] = &utilities.HeapItem{
		Value:    &start,
		Priority: 0,
	}
	heap.Init(&open)

	deltas := []Point{
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 0},
	}

	// VV: Just copied the search implementation from last day
	reverse_path := map[Point]Point{}

	push_new_point := func(point, source Point, score int) {
		g_score[point] = score

		if _, ok := reverse_path[point]; ok {
			reverse_path[point] = source
			return
		}

		reverse_path[point] = source

		t := &utilities.HeapItem{
			Value:    &point,
			Priority: score,
		}

		heap.Push(&open, t)
	}

	for open.Len() > 0 {

		heap_item := heap.Pop(&open)

		item := heap_item.(*utilities.HeapItem)

		cur := item.Value.(*Point)

		if !cur.going_back {
			// VV: This is a fruit run, stop when we've collected all kinds of fruits
			if cur.fruits == (1<<puzzle.NumFruits)-1 {
				next := *cur
				next.going_back = true

				push_new_point(next, *cur, g_score[*cur])
			}
		} else {
			// VV: This is us returning back to the entrance
			if cur.x == start.x && cur.y == start.y {
				return g_score[*cur]
			}
		}

		cur_g_score := g_score[*cur]

		for _, d := range deltas {
			next := cur.Add(d)

			if next.x < 0 || next.y < 0 || next.x >= len(puzzle.Board[0]) || next.y >= len(puzzle.Board) || puzzle.Board[cur.y][cur.x] == rune('#') || puzzle.Board[cur.y][cur.x] == rune('~') {
				continue
			}

			if !cur.going_back && puzzle.Board[cur.y][cur.x] != '.' {
				next.fruits |= 1 << (puzzle.Board[cur.y][cur.x] - rune('A'))
			}

			tentative_gscore := cur_g_score + 1

			if known_g, ok := g_score[next]; ok && known_g < tentative_gscore {
				continue
			} else {
				push_new_point(next, *cur, tentative_gscore)
			}
		}
	}

	panic("no way back!")
}

func solution(input *utilities.Puzzle) int {

	for x := 0; x < len(input.Board[0]); x++ {
		if input.Board[0][x] == rune('.') {
			start := Point{x: x}
			return calc_distance_to_any_fruit(start, input)
		}
	}

	panic("oh no")
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
