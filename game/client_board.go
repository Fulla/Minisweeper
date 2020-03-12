package game

import "sync"

type ClientBoard struct {
	safePoints    map[Point]int
	flags         map[Point]bool
	mines         []Point
	activatedMine *Point
	l             sync.Mutex
}

func NewClientBoard() *ClientBoard {
	return &ClientBoard{
		safePoints: make(map[Point]int, 0),
		flags:      make(map[Point]bool, 0),
		mines:      make([]Point, 0),
	}
}

func (cl *ClientBoard) isDiscovered(point Point) bool {
	cl.l.Lock()
	defer cl.l.Unlock()
	if _, ok := cl.safePoints[point]; ok {
		return true
	}
	return false
}

func (cl *ClientBoard) discoverSafePoints(safe map[Point]int) {
	cl.l.Lock()
	defer cl.l.Unlock()
	for point, number := range safe {
		cl.safePoints[point] = number
	}
}

func (cl *ClientBoard) setMines(mines []Point, activated *Point) {
	cl.l.Lock()
	defer cl.l.Unlock()
	cl.mines = mines
	cl.activatedMine = activated
}

func (cl *ClientBoard) SafePoints() map[Point]int {
	cl.l.Lock()
	defer cl.l.Unlock()
	return cl.safePoints
}

func (cl *ClientBoard) Mines() []Point {
	cl.l.Lock()
	defer cl.l.Unlock()
	return cl.mines
}

func (cl *ClientBoard) Flags() []Point {
	cl.l.Lock()
	defer cl.l.Unlock()
	flags := make([]Point, 0)
	for fl := range cl.flags {
		flags = append(flags, fl)
	}
	return flags
}

func (cl *ClientBoard) Activated() *Point {
	cl.l.Lock()
	defer cl.l.Unlock()
	return cl.activatedMine
}
