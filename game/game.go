package game

type Point struct {
	file   int
	column int
}

type Game struct {
	state       string
	board       *Board
	clientBoard *ClientBoard
}

// NewGame createas a new Board, but without setting the mines location,
// as we need to ensure that the first move is a safe move
// It also creates a pristine ClientBoard
func NewGame(f, c, m int) *Game {
	gm := &Game{
		state:       "initial",
		board:       NewBoard(f, c, m),
		clientBoard: NewClientBoard(),
	}
	return gm
}

func (gm *Game) State() string {
	return gm.state
}

func (gm *Game) ClientBoard() *ClientBoard {
	return gm.clientBoard
}

func (gm *Game) Discover(point Point) map[Point]int {
	gm.startGame(point)
	discovered := make(map[Point]int)
	if gm.clientBoard.isDiscovered(point) {
		return discovered
	}
	discovered, isMine := gm.board.discover(point, gm.clientBoard.isDiscovered)
	if isMine {
		gm.setGameOver(point)
		return discovered
	}
	gm.clientBoard.discoverSafePoints(discovered)
	return discovered
}

func (gm *Game) startGame(point Point) {
	if gm.state != "initial" {
		return
	}
	gm.board.generateMines(point)
	gm.state = "playing"
}

func (gm *Game) setGameOver(mine Point) {
	mines := gm.board.minesList()
	gm.clientBoard.setMines(mines, &mine)
	gm.state = "game over"
}
