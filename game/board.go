package game

import (
	"math/rand"
	"sync"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
		b.setMine(p)
		return
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
	for x := target.File - 1; x <= target.File+1; x++ {
		for y := target.Column - 1; y <= target.Column+1; y++ {
			if x < 0 || x >= b.files {
				continue
			}
			if y < 0 || y >= b.columns {
				continue
			}
			if x == target.File && y == target.Column {
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
		logrus.WithField("point", point).Infof("Pushed mine")
		return discovered, true
	}
	remaining := make(chan Point, 32)
	remaining <- point

	var current Point
LOOP:
	for {
		select {
		case current = <-remaining:
		default:
			break LOOP
		}
		logrus.WithField("point", current).Infof("Discovering")
		surr := b.surroundingTiles(current)
		number := b.numberOfMines(surr)

		if _, ok := discovered[current]; ok || omit(current) {
			continue
		}
		discovered[current] = number
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
