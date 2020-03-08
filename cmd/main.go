package main

import (
	"context"

	"github.com/Fulla/Minisweeper/server"
)

func main() {
	c := context.Background()
	server.Serve(c)
}
