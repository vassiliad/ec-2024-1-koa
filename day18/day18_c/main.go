package main

import (
	"container/heap"
	"os"

	"puzzle/utilities"
)

type Point struct {
	x, y int
}

func (p Point) Add(o Point) Point {
	return Point{x: p.x + o.x, y: p.y + o.y}
}

func flood(start Point, puzzle *utilities.Puzzle) map[Point]int {
	g_score := map[Point]int{start: 0}

	reverse_path := map[Point]Point{}
	open := make(utilities.PriorityQueue, 1)

	open[0] = &utilities.HeapItem{
		Value:    &start,
		Priority: 0,
	}
	heap.Init(&open)

	deltas := []Point{
		{x: -1, y: 0},
		{x: +1, y: 0},
		{x: 0, y: -1},
		{x: 0, y: +1},
	}

	for open.Len() > 0 {

		heap_item := heap.Pop(&open)

		item := heap_item.(*utilities.HeapItem)

		cur := item.Value.(*Point)
		cur_g_score := g_score[*cur]
		next_cost := cur_g_score + 1

		for _, d := range deltas {
			next := d.Add(*cur)

			if next.x < 0 || next.x >= len(puzzle.Board[0]) || next.y < 0 || next.y >= len(puzzle.Board) || puzzle.Board[next.y][next.x] == '#' {
				continue
			}

			if known_g, ok := g_score[next]; ok && known_g < next_cost {
				continue
			} else {

				g_score[next] = next_cost
				reverse_path[next] = *cur

				updated := false

				for idx, q := range open {
					if *q.Value.(*Point) == next {
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

	return g_score
}

func solution(input *utilities.Puzzle) int {
	// VV: so ... we do need a flood. There's probabably a way to reuse the deltas between trees to only do 1 flood
	// but this was easier to implement.
	// Just do 1 flood per tree and find the point which has the minimum distance from all trees

	tree_distances := []map[Point]int{}

	for y, row := range input.Board {
		for x, c := range row {
			if c == 'P' {

				tree_distances = append(tree_distances, flood(Point{x: x, y: y}, input))
			}
		}
	}

	min_sum_distance := int(^uint(0) >> 1)

	for y, row := range input.Board {
		for x, c := range row {
			if c == '.' {
				t := 0

				for _, dis := range tree_distances {
					if d, ok := dis[Point{x: x, y: y}]; ok {
						t += d
					} else {
						panic("oh no")
					}
				}

				min_sum_distance = min(min_sum_distance, t)
			}
		}
	}

	return min_sum_distance
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
