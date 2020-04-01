package life

import (
	"image"
	"image/color"
	"math/rand"
)

const DIMENSION = 400

type cell struct {
	X int
	Y int
}

type stateCell struct {
	X     int
	Y     int
	state bool
}


type Cells map[string]cell

func (cells *Cells) Add(c cell) {
	(*cells)[string(c.X) + " " + string(c.Y)] = c
}

type life struct {
	Cells      Cells
	neighbours map[stateCell]int
}

/*
Any live cell with fewer than two live neighbours dies, as if by underpopulation.
Any live cell with more than three live neighbours dies, as if by overpopulation.

Any live cell with two or three live neighbours lives on to the next generation.
Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction

for every alive cell note their neighbours and their state

*/



func (l *life) findNeighbours() {
	l.neighbours = make(map[stateCell]int)
	for _, c := range l.Cells {
		for _, n := range c.neighbours() {
			_, ok := l.Cells[string(n.X) + " " + string(n.Y)]
			state := ok
			l.neighbours[stateCell{X: n.X, Y: n.Y, state: state}]++
		}
	}
}

func (l *life) Tick() life {
	l.findNeighbours()

	result := life{
		Cells:      make(map[string]cell),
		neighbours: make(map[stateCell]int),
	}

	for c, neighbourCount := range l.neighbours {
		if c.state == false && neighbourCount == 3 {
			result.Cells.Add(cell{X: c.X, Y: c.Y})
		} else if c.state == true && neighbourCount == 2 || neighbourCount == 3 {
			result.Cells.Add(cell{X: c.X, Y: c.Y})
		}
	}

	return result
}

func (l *life) Draw() image.Image {
	width := DIMENSION
	height := DIMENSION

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for _, c := range l.Cells {
		img.Set(c.X, c.Y, color.Black)
	}

	return img.SubImage(img.Rect)
}

func (c cell) left() cell {
	return cell{X: c.X - 1, Y: c.Y}
}

func (c cell) right() cell {
	return cell{X: c.X + 1, Y: c.Y}
}

func (c cell) up() cell {
	return cell{X: c.X, Y: c.Y - 1}
}

func (c cell) down() cell {
	return cell{X: c.X, Y: c.Y + 1}
}

func (c *cell) neighbours() []cell {
	return []cell{
		c.left(),
		c.left().up(),
		c.up(),
		c.up().right(),
		c.right(),
		c.right().down(),
		c.down(),
		c.down().left(),
	}
}

func getRandomNum() int {
	min := 1
	max := 6
	return rand.Intn(max-min) + min
}

func Initialise() life {
	l := life{
		Cells:      make(map[string]cell),
		neighbours: make(map[stateCell]int),
	}

	for i := 0; i < DIMENSION; i++ {
		for j := 0; j < DIMENSION; j++ {
			if getRandomNum() > 4 {
				l.Cells.Add(cell{
					X: i,
					Y: j,
				})
			}
		}
	}

	return l
}


