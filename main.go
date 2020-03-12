package main

import (
	"context"

	"github.com/Fulla/Minisweeper/gamesmanager"

	"github.com/Fulla/Minisweeper/server"
)

func main() {
	c := context.Background()
	gm := gamesmanager.NewManager()
	s := server.NewServer(gm)
	s.Serve(c)
}
