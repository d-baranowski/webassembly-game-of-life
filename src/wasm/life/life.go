package life

import (
	"image"
	"image/color"
	"math/rand"
)

// Keep track of changed cells
// Create a grid of tightly connected cells by generating a 2d array first
// For every changed cell and their neighbours check if they should change
// Change all cells that should change and make the should change list new changed

// Flip the ones that have changed
// For every cell in changed list check if it or its neighbours should change and create a new changed list

const TOP = 0
const TOP_RIGHT = 1
const RIGHT = 2
const BOTTOM_RIGHT = 3
const BOTTOM = 4
const BOTTOM_LEFT = 5
const LEFT = 6
const TOP_LEFT = 7

type Link struct {
	value *Cell
	next  *Link
}

type void struct{}
var member void

func (l *Link) add(c *Cell) *Link {
	l.value = c
	return &Link{
		value: nil,
		next:  l,
	}
}

type Cell struct {
	neighbours [8]*Cell
	X          uint16
	Y          uint16
	alive      bool
}

func (c *Cell) ShouldFlip() bool {
	var liveNeighbourCount uint8 = 0
	for _, n := range c.neighbours {
		if n.alive {
			liveNeighbourCount++
		}
	}

	//Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	//Any live cell with more than three live neighbours dies, as if by overpopulation.
	if c.alive && (liveNeighbourCount < 2 || liveNeighbourCount > 3) {
		return true
	}

	//Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	if !c.alive && liveNeighbourCount == 3 {
		return true
	}

	//Any live cell with two or three live neighbours lives on to the next generation.
	return false
}

type Life struct {
	grid    *Cell
	w       int
	h       int
	changed *Link
}

func (l *Life) getNeighbourCoordinates(x int, y int) (uint16, uint16) {
	x += l.w
	x %= l.w
	y += l.h
	y %= l.h

	return uint16(x), uint16(y)
}

func RandomMap(w int, h int) [][]int {
	reference := make([][]int, w)
	for x := 0; x < w; x++ {
		reference[x] = make([]int, h)

		for y := 0; y < h; y++ {
			if rand.Intn(32) % 3 == 0 {
				reference[x][y] = 1
			}
		}
	}

	return reference
}

func (l *Life) Initialise(w int, h int, initialState [][]int) {
	l.w = w
	l.h = h
	reference := make([][]*Cell, w)
	l.changed = &Link{}

	for x := 0; x < w; x++ {
		reference[x] = make([]*Cell, h)

		for y := 0; y < h; y++ {
			c := Cell{
				neighbours: [8]*Cell{},
				X:          uint16(x),
				Y:          uint16(y),
				alive:      initialState[x][y] == 1,
			}
			reference[x][y] = &c
		}
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := reference[x][y]
			nx, ny := l.getNeighbourCoordinates(x, y-1)
			c.neighbours[TOP] = reference[nx][ny]
			nx, ny = l.getNeighbourCoordinates(x+1, y-1)
			c.neighbours[TOP_RIGHT] = reference[nx][ny]
			nx, ny = l.getNeighbourCoordinates(x+1, y)
			c.neighbours[RIGHT] = reference[nx][ny]
			nx, ny = l.getNeighbourCoordinates(x+1, y+1)
			c.neighbours[BOTTOM_RIGHT] = reference[nx][ny]
			nx, ny = l.getNeighbourCoordinates(x, y+1)
			c.neighbours[BOTTOM] = reference[nx][ny]
			nx, ny = l.getNeighbourCoordinates(x-1, y+1)
			c.neighbours[BOTTOM_LEFT] = reference[nx][ny]
			nx, ny = l.getNeighbourCoordinates(x-1, y)
			c.neighbours[LEFT] = reference[nx][ny]
			nx, ny = l.getNeighbourCoordinates(x-1, y-1)
			c.neighbours[TOP_LEFT] = reference[nx][ny]
			l.changed = l.changed.add(c)
		}
	}

	l.grid = reference[0][0]

	l.Tick()
}

func (l *Life) Tick() {
	// Flip the ones that have changed
	i := l.changed.next
	for i != nil {
		cell := i.value
		cell.alive = !cell.alive
		i = i.next
	}

	newChanged := &Link{}

	// For every cell in changed list check if it or its neighbours should change and create a new changed list
	i = l.changed.next
	set := make(map[*Cell]void)
	for i != nil {
		cell := i.value
		set[cell] = member
		for _, neighbour := range cell.neighbours {
			set[neighbour] = member
		}

		i = i.next
	}

	for cell, _ := range set {
		if cell.ShouldFlip() {
			newChanged = newChanged.add(cell)
		}
	}

	l.changed = newChanged
}

func(l *Life) Draw(img *image.RGBA) *image.RGBA {
	/*upLeft := image.Point{0, 0}
	lowRight := image.Point{l.w, l.h}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight}) */

	topLeft := l.grid
	if topLeft.alive {
		img.Set(int(topLeft.X), int(topLeft.Y), color.Black)
	}

	row := topLeft.neighbours[BOTTOM]
	for i := 0; i < l.h; i++ {
		if row.alive {
			img.Set(int(row.X), int(row.Y), color.Black)
		}

		cell := row.neighbours[RIGHT]
		for j := 0; j < l.w; j++ {
			if cell.alive {
				img.Set(int(cell.X), int(cell.Y), color.Black)
			}

			cell = cell.neighbours[RIGHT]
		}

		row = row.neighbours[BOTTOM]
	}

	return img
}

func(l *Life) ForEachAlive(do func(*Cell)) {
	topLeft := l.grid
	if topLeft.alive {
		do(topLeft)
	}

	row := topLeft.neighbours[BOTTOM]
	for i := 0; i < l.h; i++ {
		if row.alive {
			do(topLeft)
		}

		cell := row.neighbours[RIGHT]
		for j := 0; j < l.w; j++ {
			if cell.alive {
				do(topLeft)
			}

			cell = cell.neighbours[RIGHT]
		}

		row = row.neighbours[BOTTOM]
	}
}

func(l *Life) Print() string {
	result := ""
	row := l.grid
	for i := 0; i < l.h; i++ {
		stringRow := ""

		if row.alive {
			stringRow += "1"
		} else {
			stringRow += "0"
		}

		column := row.neighbours[RIGHT]
		for j := 0; j < l.w; j++ {
			if column.alive {
				stringRow += "1"
			} else {
				stringRow += "0"
			}

			column = column.neighbours[RIGHT]
		}

		result += stringRow + "\n"
		row = row.neighbours[BOTTOM]
	}

	return result
}