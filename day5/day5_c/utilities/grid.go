package utilities

type Grid[T any] struct {
	Points []T
	Width  int
	Height int
}

func GridNew[T any](width int, height int) Grid[T] {
	return Grid[T]{
		Points: make([]T, width*height),
		Width:  width,
		Height: height,
	}
}

func (g *Grid[T]) Update(x, y int, v T) {
	g.Points[x+y*g.Width] = v
}

func (g *Grid[T]) Get(x, y int) T {
	return g.Points[x+y*g.Width]
}

// Shifts a column up and puts the top-most value at the bottom point
func (g *Grid[T]) ShiftUp(x int) {
	removed := g.Get(x, 0)

	for i := range g.Height - 1 {
		g.Update(x, i, g.Get(x, i+1))
	}

	g.Update(x, g.Width-1, removed)
}

// Inserts a new value at position x, y, shifting the old Point one slot down
// the very last value in the column is discarded
func (g *Grid[T]) InsertInColumn(x, y int, v T) {
	for i := range g.Height - y - 1 {
		g.Update(x, y+i+1, g.Get(x, i+y))
	}

	g.Update(x, y, v)
}
