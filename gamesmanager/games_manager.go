package gamesmanager

import (
	"context"

	"github.com/Fulla/Minisweeper/game"
	"golang.org/x/sync/semaphore"
)

type GamesManager struct {
	// single user, without persistence
	g   *game.Game
	sem *semaphore.Weighted
}

func NewManager() *GamesManager {
	return &GamesManager{
		sem: semaphore.NewWeighted(1),
	}
}

func (gm *GamesManager) StartGame(ctx context.Context, f, c, m int) *game.Game {
	g := game.NewGame(f, c, m)
	// by now, just set the game as the only one game for the manager
	gm.g = g
	gm.sem.Acquire(ctx, 1)
	return g
}

func (gm *GamesManager) GetGame(ctx context.Context) *game.Game {
	gm.sem.Acquire(ctx, 1)
	return gm.g
}

func (gm *GamesManager) FreeGame() {
	gm.sem.Release(1)
}
