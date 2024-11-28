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

func find_path(start, end Point, puzzle *utilities.Puzzle) int {
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

		if *cur == end {
			return g_score[*cur]
		}

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

	panic("oh no")
}

func solution(input *utilities.Puzzle) int {
	trees := []Point{}

	all_start := []Point{}

	for y, row := range input.Board {
		for x, c := range row {
			if c == '.' && (x == 0 || x == len(row)-1 || y == 0 || y == len(input.Board)-1) {
				all_start = append(all_start, Point{x: x, y: y})
			}

			if c == 'P' {
				trees = append(trees, Point{x: x, y: y})
			}
		}
	}

	// VV: Similar to the task before, only now for each tree we care about the minimum distance to the 2 starts
	// I should have implemented flood and be done with it quicker, but I was too lazy :)
	max_time := -1

	for _, tree := range trees {
		tree_min := int(^uint(0) >> 1)

		for _, start := range all_start {
			t := find_path(start, tree, input)
			tree_min = min(tree_min, t)
		}
		max_time = max(max_time, tree_min)
	}

	return max_time
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
