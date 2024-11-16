package main

import (
	"container/list"
	"fmt"
	"os"

	"ec-2024-1-koa/day6_b/utilities"
)

type Point struct {
	CurrentNode string
	Distance    int
	Path        string
}

func solution(input *utilities.Puzzle) string {
	remaining := list.New()

	fruit_distances := map[int][]Point{}

	remaining.PushFront(Point{CurrentNode: "RR", Distance: 0, Path: "R"})

	cur := remaining.Front()

	total_fruits := 0

	for {
		if cur == nil {
			break
		}

		branch := cur.Value.(Point)

		if _, ok := input.Branches[branch.CurrentNode]; !ok {
			cur = cur.Next()
			continue
		}

		for _, n := range input.Branches[branch.CurrentNode] {
			path := fmt.Sprintf("%s%c", branch.Path, n[0])
			child := Point{
				Distance:    branch.Distance + 1,
				Path:        path,
				CurrentNode: n,
			}

			if n == "@" {
				total_fruits += 1
				if others, ok := fruit_distances[child.Distance]; ok {
					fruit_distances[child.Distance] = append(others, child)
				} else {
					fruit_distances[child.Distance] = []Point{child}
				}
			} else {
				remaining.PushBack(child)
			}

		}

		cur = cur.Next()
	}

	for _, fruits := range fruit_distances {
		if len(fruits) == 1 {
			return fruits[0].Path
		}
	}

	panic("Should never happen")
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
