package gamesmanager

import (
	"context"
	"sync"
	"time"

	"github.com/Fulla/Minisweeper/game"
	"golang.org/x/sync/semaphore"
)

type playerInstance struct {
	Sem        *semaphore.Weighted
	LastAccess time.Time
	Game       *game.Game
}

type GamesManager struct {
	// support for 64 active games
	cache map[string]*playerInstance
	l     sync.Mutex
}

func NewManager() *GamesManager {
	return &GamesManager{
		cache: make(map[string]*playerInstance, 64),
	}
}

func (gm *GamesManager) StartGame(ctx context.Context, ip string, f, c, m int) *game.Game {
	gm.l.Lock()
	defer gm.l.Unlock()
	g := game.NewGame(f, c, m)

	// if player exists
	if pl := gm.cache[ip]; pl != nil {
		pl.LastAccess = time.Now()
		pl.Sem.Acquire(ctx, 1)
		pl.Game = g
		return g
	}

	// if player does not exist

	// if the cache is full
	if len(gm.cache) >= 64 {
		last := time.Now()
		toDelete := ""
		for ip, pl := range gm.cache {
			if pl.LastAccess.Before(last) {
				last = pl.LastAccess
				toDelete = ip
			}
		}
		delete(gm.cache, toDelete)
	}

	// create player instance and save to cache
	pl := &playerInstance{
		Sem:        semaphore.NewWeighted(1),
		LastAccess: time.Now(),
		Game:       g,
	}
	pl.Sem.Acquire(ctx, 1)
	gm.cache[ip] = pl

	return g
}

func (gm *GamesManager) GetGame(ctx context.Context, ip string) *game.Game {
	pl := gm.cache[ip]
	if pl == nil {
		return nil
	}
	pl.LastAccess = time.Now()
	pl.Sem.Acquire(ctx, 1)
	return pl.Game
}

func (gm *GamesManager) FreeGame(ip string) {
	pl := gm.cache[ip]
	if pl == nil {
		return
	}
	pl.LastAccess = time.Now()
	pl.Sem.Release(1)
}
