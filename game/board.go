package game

import (
	"math/rand"

	"github.com/pkg/errors"
)

type Board struct {
	files      int
	columns    int
	minesCount int
	// mines is a map just storing true at the points there is a mine
	mines map[Point]bool
}

func NewBoard(f, c, m int) *Board {
	return &Board{
		files:      f,
		columns:    c,
		minesCount: m,
		mines:      make(map[Point]bool),
	}
}

func (b *Board) randomMineExcept(safe Point) {
	for {
		x := rand.Intn(b.files)
		y := rand.Intn(b.columns)
		p := Point{x, y}
		if p == safe {
			continue
		}
		if b.mines[p] {
			continue
		}
		b.mines[p] = true
		break
	}
}

func (b *Board) generateMines(safe Point, number int) error {
	if number >= b.columns*b.files {
		return errors.Errorf("Cannot generate more mines than available tiles")
	}
	for i := 0; i < number; i++ {
		b.randomMineExcept(safe)
	}
	return nil
}

func (b *Board) surroundingTiles(target Point) []Point {
	surroundings := []Point{}
	for x := target.file - 1; x <= target.file+1; x++ {
		for y := target.column - 1; y <= target.column+1; y++ {
			if x < 0 || x >= b.files {
				continue
			}
			if y < 0 || y >= b.columns {
				continue
			}
			if x == target.file && y == target.column {
				continue
			}
			surroundings = append(surroundings, Point{x, y})
		}
	}
	return surroundings
}

func (b *Board) numberOfMines(points []Point) (number int) {
	for _, p := range points {
		if b.mines[p] {
			number = number + 1
		}
	}
	return
}
