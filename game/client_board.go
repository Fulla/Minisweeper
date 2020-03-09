package game

type ClientBoard struct {
	safePoints    map[Point]int
	flags         map[Point]bool
	mines         []Point
	activatedMine *Point
}

func NewClientBoard() *ClientBoard {
	return &ClientBoard{
		safePoints: make(map[Point]int, 0),
		flags:      make(map[Point]bool, 0),
	}
}

func (cl *ClientBoard) isDiscovered(point Point) bool {
	if _, ok := cl.safePoints[point]; ok {
		return true
	}
	return false
}

func (cl *ClientBoard) discoverSafePoints(safe map[Point]int) {
	for point, number := range safe {
		cl.safePoints[point] = number
	}
}

func (cl *ClientBoard) setMines(mines []Point, activated *Point) {
	cl.mines = mines
	cl.activatedMine = activated
}

func (cl *ClientBoard) SafePoints() map[Point]int {
	return cl.safePoints
}

func (cl *ClientBoard) Mines() []Point {
	return cl.mines
}

func (cl *ClientBoard) Flags() []Point {
	return cl.mines
}

func (cl *ClientBoard) Activated() *Point {
	return cl.activatedMine
}
