package game

type Point struct {
	file   int
	column int
}

type GameManager struct {
	state       string
	board       *Board
	clientBoard *ClientBoard
}

// NewGame createas a new Board, but without setting the mines location,
// as we need to ensure that the first move is a safe move
// It also creates a pristine ClientBoard
func (gm *GameManager) NewGame(f, c, m int) {
	gm.board = NewBoard(f, c, m)
	gm.state = "initial"
	gm.clientBoard = NewClientBoard()
}

func (gm *GameManager) discover(point Point) {
	if gm.clientBoard.isDiscovered(point) {
		return
	}
}
