package game

import (
	"math/rand"
	"sync"

	"github.com/pkg/errors"
)

type Board struct {
	files      int
	columns    int
	minesCount int
	// mines is a map just storing true at the points there is a mine
	mines map[Point]bool
	l     sync.Mutex
}

func NewBoard(f, c, m int) *Board {
	return &Board{
		files:      f,
		columns:    c,
		minesCount: m,
		mines:      make(map[Point]bool),
	}
}

func (b *Board) Dimensions() (int, int) {
	return b.files, b.columns
}

func (b *Board) randomMineExcept(safe Point) {
	for {
		x := rand.Intn(b.files)
		y := rand.Intn(b.columns)
		p := Point{x, y}
		if p == safe {
			continue
		}
		if b.isMine(p) {
			continue
		}

		break
	}
}

func (b *Board) generateMines(safe Point) error {
	if b.minesCount >= b.columns*b.files {
		return errors.Errorf("Cannot generate more mines than available tiles")
	}
	for i := 0; i < b.minesCount; i++ {
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
		if b.isMine(p) {
			number = number + 1
		}
	}
	return
}

// discovers a point of the board
// if the point is a mine, returns true as the second return value
// if the point is not surrounded by any mine, then it iteratively discover the surroundings
// Returns a map with the number of surrounding mines for each discovered point
func (b *Board) discover(point Point, omit func(Point) bool) (map[Point]int, bool) {
	discovered := make(map[Point]int)
	if b.isMine(point) {
		return discovered, true
	}
	remaining := make(chan Point, 32)
	remaining <- point

	for p := range remaining {
		surr := b.surroundingTiles(point)
		number := b.numberOfMines(surr)

		if _, ok := discovered[p]; ok || omit(p) {
			continue
		}
		discovered[p] = number
		if number != 0 {
			continue
		}
		for _, s := range surr {
			remaining <- s
		}
	}
	return discovered, false
}

func (b *Board) isMine(p Point) bool {
	b.l.Lock()
	defer b.l.Unlock()
	return b.mines[p]
}

func (b *Board) setMine(p Point) {
	b.l.Lock()
	defer b.l.Unlock()
	b.mines[p] = true
}

func (b *Board) minesList() (mines []Point) {
	b.l.Lock()
	defer b.l.Unlock()
	for m := range b.mines {
		mines = append(mines, m)
	}
	return
}
