package game

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSurroundings(t *testing.T) {
	b := NewBoard(4, 4, 0)

	s := b.surroundingTiles(Point{0, 0})
	assert.Equal(t, []Point{Point{0, 1}, Point{1, 0}, Point{1, 1}}, s)

	s = b.surroundingTiles(Point{3, 3})
	assert.Equal(t, []Point{Point{2, 2}, Point{2, 3}, Point{3, 2}}, s)

	s = b.surroundingTiles(Point{2, 1})
	assert.Equal(t, []Point{
		Point{1, 0}, Point{1, 1}, Point{1, 2},
		Point{2, 0}, Point{2, 2},
		Point{3, 0}, Point{3, 1}, Point{3, 2},
	}, s)

	s = b.surroundingTiles(Point{5, 5})
	assert.Equal(t, []Point{}, s)
}

func TestNumberOfMines(t *testing.T) {
	b := NewBoard(4, 4, 0)
	b.mines = map[Point]bool{Point{2, 2}: true, Point{2, 3}: true, Point{1, 2}: true}

	// test surroundings of Point{3, 3}
	surr := []Point{Point{2, 2}, Point{2, 3}, Point{3, 2}}
	assert.Equal(t, 2, b.numberOfMines(surr))
	// test surroundings of Point{0, 0}
	surr = []Point{Point{0, 1}, Point{1, 0}, Point{1, 1}}
	assert.Equal(t, 0, b.numberOfMines(surr))
	// test surroundings of Point{1, 3}
	surr = []Point{Point{0, 2}, Point{0, 3}, Point{1, 2}, Point{2, 2}, Point{2, 3}}
	assert.Equal(t, 3, b.numberOfMines(surr))

	// I don't use the surroundingTiles function to get the surrounding points because
	// that's not the function I'm testing here
}
