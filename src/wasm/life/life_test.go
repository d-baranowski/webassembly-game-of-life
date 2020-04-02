package life

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Still lifes

func TestBox(t *testing.T) {
	initialState := [][]int{
		{1,1,1,1,1},
		{1,1,1,1,1},
		{1,1,0,0,1},
		{1,1,0,0,1},
		{1,1,1,1,1},
	}

	l := &Life{}
	l.Initialise(5, 5, initialState)


	s1 := l.Print()
	l.Tick()
	s2 :=l.Print()
	l.Tick()
	s3 :=l.Print()

	assert.Equal(t, s1, s2, "The two states should be the same.")
	assert.Equal(t, s2, s3, "The two states should be the same.")
}

func TestBlinker(t *testing.T) {
	initialState := [][]int{
		{1,1,1,1,1},
		{1,1,1,1,1},
		{1,0,0,0,1},
		{1,1,1,1,1},
		{1,1,1,1,1},
	}

	blinkA :=
		"000000\n" +
		"001000\n" +
		"001000\n" +
		"001000\n" +
		"000000\n"

	blinkB :=
		"000000\n" +
		"000000\n" +
		"011100\n" +
		"000000\n" +
		"000000\n"

	l := &Life{}
	l.Initialise(5, 5, initialState)


	l.Tick()
	s1 :=l.Print()
	l.Tick()
	s2 :=l.Print()

	assert.Equal(t, s1, blinkB, "The two states should be the same.")
	assert.Equal(t, s2, blinkA, "The two states should be the same.")
}

func TestBeeHive(t *testing.T) {
	//l := Life{
	//	Cells:      make(map[string]cell),
	//	neighbours: make(map[stateCell]int),
	//}
	//
	//reference := cell{X: 5, Y: 5}
	//l.Cells.Add(reference)
	//l.Cells.Add(reference.right().up())
	//l.Cells.Add(reference.right().down())
	//l.Cells.Add(reference.right().right().up())
	//l.Cells.Add(reference.right().right().down())
	//l.Cells.Add(reference.right().right().right())
	//
	//l2 := l.Tick()
	//l3 := l2.Tick()
	//
	//assert.Equal(t, l.Cells, l2.Cells, "The two states should be the same.")
	//assert.Equal(t, l.Cells, l3.Cells, "The two states should be the same.")
}

func TestLink(t *testing.T) {
	cells := make([]*Cell, 0)
	for i := 0; i < 10; i++ {
		cells = append(cells, &Cell{X: uint16(i), Y: 0})
	}

	l := &Link{}
	for _, cell := range cells {
		l = l.add(cell)
	}

	var i *Link = l.next

	contains := func(val *Cell) bool {
		for _, cell := range cells {
			if val == cell {
				return true
			}
		}

		return false
	}

	fail := false
	for i != nil {
		if !contains(i.value) {
			fail = true
		}
		i = i.next
	}

	if fail {
		assert.Fail(t, "Linked list should contain all the values")
	}
}

func BenchmarkInitialise(b *testing.B) {
	l := Life{}
	l.Initialise(1000, 1000, RandomMap(1000, 1000))
}

func BenchmarkTick(b *testing.B) {
	l := Life{}
	l.Initialise(1000, 1000, RandomMap(1000, 1000))
	b.StartTimer()
	for i := 0; i < 10; i++ {
		l.Tick()
	}
}
