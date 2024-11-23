package main

import (
	"container/heap"
	"os"

	"puzzle/utilities"
)

type Point struct {
	x, y, z int
}

func (p Point) Add(o Point) Point {
	return Point{x: p.x + o.x, y: p.y + o.y, z: p.z + o.z}
}

func calc_distance_to_any_fruit(start Point, puzzle *utilities.Puzzle) int {
	// VV: Technically we need only find distances to the main trunk
	// for my input this ran fast enough so that I didn't mind computing distances to all points

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
	for open.Len() > 0 {

		heap_item := heap.Pop(&open)

		item := heap_item.(*utilities.HeapItem)

		cur := item.Value.(*Point)

		if puzzle.Board[cur.y][cur.x] != '.' && puzzle.Board[cur.y][cur.x] != '#' {

			for t := *cur; t != start; t = reverse_path[t] {
				puzzle.Board[t.y][t.x] = rune('*')
			}

			for _, row := range puzzle.Board {
				println(string(row))
			}

			return g_score[*cur]
		}

		cur_g_score := g_score[*cur]

		for _, d := range deltas {
			next := cur.Add(d)

			if next.x < 0 || next.y < 0 || next.x >= len(puzzle.Board[0]) || next.y >= len(puzzle.Board) || puzzle.Board[cur.y][cur.x] == rune('#') {
				continue
			}

			tentative_gscore := cur_g_score + 1

			if known_g, ok := g_score[next]; ok && known_g < tentative_gscore {
				continue
			} else {

				g_score[next] = tentative_gscore

				if _, ok := reverse_path[next]; ok {
					reverse_path[next] = *cur
					continue
				}
				reverse_path[next] = *cur

				t := &utilities.HeapItem{
					Value:    &next,
					Priority: tentative_gscore,
				}

				heap.Push(&open, t)
			}
		}
	}

	panic("oh no")
}

func solution(input *utilities.Puzzle) int {

	for x := 0; x < len(input.Board[0]); x++ {
		if input.Board[0][x] == rune('.') {
			return calc_distance_to_any_fruit(Point{x: x}, input) * 2
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
