package game

import "sync"

type Point struct {
	file   int
	column int
}

type Game struct {
	state       string
	board       *Board
	clientBoard *ClientBoard
	l           sync.Mutex
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
	gm.l.Lock()
	defer gm.l.Unlock()
	return gm.state
}

func (gm *Game) setState(state string) {
	gm.l.Lock()
	defer gm.l.Unlock()
	gm.state = state
}

func (gm *Game) ClientBoard() *ClientBoard {
	return gm.clientBoard
}

func (gm *Game) BoardDimensions() (int, int) {
	return gm.board.Dimensions()
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
	if gm.State() != "initial" {
		return
	}
	gm.board.generateMines(point)
	gm.setState("playing")
}

func (gm *Game) setGameOver(mine Point) {
	mines := gm.board.minesList()
	gm.clientBoard.setMines(mines, &mine)
	gm.setState("game over")
}

func (gm *Game) checkEndCondition() {
	fil, col := gm.BoardDimensions()
	tiles := fil * col
	if len(gm.clientBoard.safePoints)+len(gm.board.mines) < tiles {
		return
	}
	mines := gm.board.minesList()
	gm.clientBoard.setMines(mines, nil)
	gm.setState("win")
}
