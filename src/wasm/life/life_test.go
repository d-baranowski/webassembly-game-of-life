package life

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Still lifes

func TestBox(t *testing.T) {
	l := Life{
		Cells:      make(map[string]cell),
		neighbours: make(map[stateCell]int),
	}

	reference := cell{X: 5, Y: 5}
	l.Cells.Add(reference)
	l.Cells.Add(reference.right())
	l.Cells.Add(reference.right().down())
	l.Cells.Add(reference.down())

	l2 := l.Tick()
	l3 := l2.Tick()

	assert.Equal(t, l.Cells, l2.Cells, "The two states should be the same.")
	assert.Equal(t, l.Cells, l3.Cells, "The two states should be the same.")
}

func TestBeeHive(t *testing.T) {
	l := Life{
		Cells:      make(map[string]cell),
		neighbours: make(map[stateCell]int),
	}

	reference := cell{X: 5, Y: 5}
	l.Cells.Add(reference)
	l.Cells.Add(reference.right().up())
	l.Cells.Add(reference.right().down())
	l.Cells.Add(reference.right().right().up())
	l.Cells.Add(reference.right().right().down())
	l.Cells.Add(reference.right().right().right())

	l2 := l.Tick()
	l3 := l2.Tick()

	assert.Equal(t, l.Cells, l2.Cells, "The two states should be the same.")
	assert.Equal(t, l.Cells, l3.Cells, "The two states should be the same.")
}

func TestBlinker(t *testing.T) {
	l := Life{
		Cells:      make(map[string]cell),
		neighbours: make(map[stateCell]int),
	}

	reference := cell{X: 5, Y: 5}
	l.Cells.Add(reference)
	l.Cells.Add(reference.left())
	l.Cells.Add(reference.right())

	l2 := l.Tick()
	l3 := l2.Tick()

	stateTwo := Cells{}
	stateTwo.Add(reference)
	stateTwo.Add(reference.up())
	stateTwo.Add(reference.down())

	assert.Equal(t, l2.Cells, stateTwo, "The two states should be the same.")
	assert.Equal(t, l3.Cells, l.Cells, "The two states should be the same.")
}

// 7879084769 ns
func BenchmarkTick(b *testing.B) {
	life := Initialise(0, 0)
	for i := 0; i < 100; i++ {
		life.Tick()
	}
}
