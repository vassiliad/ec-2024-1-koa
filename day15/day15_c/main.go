package main

import (
	"container/heap"
	"fmt"
	"image"
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

func calc_distance(start, end image.Point, g_score map[image.Point]int, puzzle *utilities.Puzzle) int {
	g_score[start] = 0

	open := make(utilities.PriorityQueue, 1)

	open[0] = &utilities.HeapItem{
		Value:    &start,
		Priority: 0,
	}
	heap.Init(&open)

	deltas := []image.Point{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
	}

	// VV: Just copied the search implementation from last day
	reverse_path := map[image.Point]image.Point{}

	push_new_point := func(point, source image.Point, score int) {
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

		cur := item.Value.(*image.Point)

		if cur.X == end.X && cur.Y == end.Y {

			return g_score[*cur]
		}

		cur_g_score := g_score[*cur]
		tentative_gscore := cur_g_score + 1

		for _, d := range deltas {
			next := cur.Add(d)

			if next.X < 0 || next.Y < 0 || next.X >= len(puzzle.Board[0]) || next.Y >= len(puzzle.Board) || puzzle.Board[cur.Y][cur.X] == rune('#') || puzzle.Board[cur.Y][cur.X] == rune('~') {
				continue
			}

			if known_g, ok := g_score[next]; ok && known_g < tentative_gscore {
				continue
			} else {
				push_new_point(next, *cur, tentative_gscore)
			}
		}
	}

	fmt.Printf("Cannot go from %+v to %+v\n", start, end)
	fmt.Println(string(puzzle.Board[start.Y][start.X]), "to", string(puzzle.Board[end.Y][end.X]))
	panic("this looks like a bug, good luck")
}

func solution(input *utilities.Puzzle) int {
	/* VV: so the idea is to compress the graph a bit and group "fruit nodes" together.

	So, starting from a collection of nodes (e.g. start, fruit A, fruit B, etc) visit all
	non-visited fruits. Once you collect all fruits, head back to the starting position.

	Assumes there're fewer than N fruits where N is the number of bits in int.
	Solution uses an integer to keep track of collected fruit types.

	Unfortunately, I couldn't figure out how to take into account the special structure of
	the input so this thing took 6 minutes on my machine. On the plus side it's a general solution!
	*/
	var start image.Point

	for x := 0; x < len(input.Board[0]); x++ {
		if input.Board[0][x] == rune('.') {
			start = image.Point{X: x}
		}
	}

	fruits := map[int][]image.Point{}

	// VV: keys are source points, values are the g_score maps
	all_distances := map[image.Point]map[image.Point]int{}

	start_distances := map[image.Point]int{}

	for y, row := range input.Board {
		for x, c := range row {
			if c < rune('A') || c > rune('Z') {
				continue
			}
			idx := int(c - rune('A'))
			pos := image.Pt(x, y)
			fruits[idx] = append(fruits[idx], pos)
		}
	}

	// VV: This just computes distance between the starting point and all fruits
	for _, all_i := range fruits {
		for _, pos := range all_i {
			start_gscore := map[image.Point]int{}
			dist := calc_distance(start, pos, start_gscore, input)

			start_distances[pos] = dist
			if other, ok := all_distances[pos]; ok {
				other[start] = dist
			} else {
				all_distances[pos] = map[image.Point]int{start: dist}
			}
		}
	}

	// VV: Now compute distances between each fruit that's not the same kind of fruit
	skipped := 0
	computed := 0
	for f_i, all_i := range fruits {
		for _, pos_i := range all_i {
			pos_i_distances := all_distances[pos_i]

			// VV: for 1 of the fruits in group X we can use the same g_score map to track the
			// distance to all fruits whose type is other than X
			i_gscore := map[image.Point]int{}

			for f_j, all_j := range fruits {

				if f_j <= f_i {
					continue
				}

				for _, pos_j := range all_j {

					computed++
					dist, ok := i_gscore[pos_j]

					if !ok {
						dist = calc_distance(pos_i, pos_j, i_gscore, input)
					} else {
						skipped++
					}

					pos_i_distances[pos_j] = dist
					if other, ok := all_distances[pos_j]; ok {
						other[pos_i] = dist
					} else {
						panic("Start should have already initialized this")
					}
				}
			}
		}

		fmt.Printf("%.2d --- skipped %d, computed %d\n", f_i, skipped, computed)
	}

	all_distances[start] = start_distances

	fmt.Printf("skipped %d computed %d\n", skipped, computed)

	start_point := Point{
		x: start.X, y: start.Y, going_back: false, fruits: 0,
	}
	return calc_distance_to_any_fruit(start_point, input.TargetFruits, all_distances, fruits)
}

func calc_distance_to_any_fruit(start Point, target_fruits int, all_distances map[image.Point]map[image.Point]int, fruits map[int][]image.Point) int {
	start.going_back = false

	fruit_types := map[image.Point]int{}

	for f_type, positions := range fruits {

		for _, pos := range positions {
			fruit_types[pos] = f_type
		}
	}

	g_score := map[Point]int{start: 0}

	open := make(utilities.PriorityQueue, 1)

	open[0] = &utilities.HeapItem{
		Value:    &start,
		Priority: 0,
	}
	heap.Init(&open)

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
			if cur.fruits == target_fruits {

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
		if !cur.going_back {
			neighbours := all_distances[image.Pt(cur.x, cur.y)]
			for next_pos, dist := range neighbours {
				next := *cur

				next.x = next_pos.X
				next.y = next_pos.Y

				if _, ok := fruit_types[next_pos]; !ok {
					if next_pos.X != start.x && next_pos.Y != start.y {
						panic("connects to non fruit")
					}
					continue
				}

				f_type := fruit_types[image.Pt(next.x, next.y)]

				if cur.fruits&(1<<f_type) == 1 {
					// VV: already collected this type of fruit
					continue
				}

				next.fruits |= 1 << f_type

				tentative_gscore := cur_g_score + dist

				if known_g, ok := g_score[next]; ok && known_g < tentative_gscore {
					continue
				} else {
					push_new_point(next, *cur, tentative_gscore)
				}
			}
		} else {
			// VV: we've already pre-computed the return trip!
			dist, ok := all_distances[image.Pt(cur.x, cur.y)][image.Pt(start.x, start.y)]

			if !ok {
				panic("Oh no there's no way back")
			}

			if cur.fruits != target_fruits {
				panic("oh no, we're returning but haven't collected every fruit yet :(")
			}

			tentative_gscore := cur_g_score + dist

			next := *cur
			next.x = start.x
			next.y = start.y

			if known_g, ok := g_score[next]; ok && known_g < tentative_gscore {
				continue
			} else {
				push_new_point(next, *cur, tentative_gscore)
			}

		}
	}

	panic("good luck")
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
