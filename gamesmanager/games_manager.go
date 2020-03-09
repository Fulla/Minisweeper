package gamesmanager

import (
	"github.com/Fulla/Minisweeper/game"
)

type GamesManager struct {
	// single user, without persistence
	g *game.Game
}

func (gm *GamesManager) StartGame(f, c, m int) *game.Game {
	g := game.NewGame(f, c, m)
	// by now, just set the game as the only one game for the manager
	gm.g = g
	return g
}

func (gm *GamesManager) GetGame() *game.Game {
	return gm.g
}
