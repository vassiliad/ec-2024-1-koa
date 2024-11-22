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

func calc_distance_to_any_point(start Point, all_points map[Point]int) map[Point]int {
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
		{x: 0, y: 1, z: 0},
		{x: 0, y: -1, z: 0},
		{x: 1, y: 0, z: 0},
		{x: -1, y: 0, z: 0},
		{x: 0, y: 0, z: 1},
		{x: 0, y: 0, z: -1},
	}

	// VV: Just copied the search implementation from last day
	reverse_path := map[Point]Point{}
	for open.Len() > 0 {

		heap_item := heap.Pop(&open)

		item := heap_item.(*utilities.HeapItem)

		cur := item.Value.(*Point)

		cur_g_score := g_score[*cur]

		for _, d := range deltas {
			next := cur.Add(d)

			if _, ok := all_points[next]; !ok {
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

	return g_score
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

	all_points := map[Point]int{}
	all_leaves := []Point{}

	// VV: Find all segments and leaves
	max_y := 0
	for _, branch := range input.Moves {
		pos := Point{}

		for _, move := range branch {

			d := deltas[move.Direction]

			for _ = range move.Repeat {
				pos = pos.Add(d)
				all_points[pos] = 1

				if pos.x == 0 && pos.z == 0 {
					max_y = max(max_y, pos.y)
				}
			}

		}

		all_leaves = append(all_leaves, pos)
	}

	// VV: Compute distances from any leaf to any point on the main trunk
	all_distances := []map[Point]int{}

	for _, leaf := range all_leaves {
		all_distances = append(all_distances, calc_distance_to_any_point(leaf, all_points))
	}

	// VV: Find a point on the main trunk which has the smallest distance from all leaves
	sol := int(^uint(0) >> 1)

	for y := range max_y + 1 {
		s := 0

		tap := Point{y: y}
		for _, distances := range all_distances {
			if dist, ok := distances[tap]; !ok {
				// VV: this leaf cannot reach the main trunk (is this even possible, all branches start at 0, 0, 0)
				s = sol
				break
			} else {
				s += dist
			}

		}

		sol = min(sol, s)
	}

	return sol
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
