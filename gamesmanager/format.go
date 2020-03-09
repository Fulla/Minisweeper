package gamesmanager

import (
	"strconv"

	"github.com/Fulla/Minisweeper/game"
	"github.com/pkg/errors"
)

// The client_board is stored in json as:
//
// {
// 	"state": "playing",
// 	"safePoints": {
// 		1: [
// 			{"x": 1, "y": 4},
// 			{"x": 2, "y": 2},
// 			...
// 			{"x": 8, "y": 2}
// 		],
// 		2: [
// 			{"x": 3, "y": 4},
// 			{"x": 7, "y": 2},
// 			...
// 			{"x": 5, "y": 9}
// 		],
// 		...
// 		8: [
// 			{"x": 11, "y": 3},
// 		]
// 	},
// 	"flags": [
// 		{"x": 11, "y": 2},
// 		{"x": 10, "y": 4},
// 		...
// 		{"x": 5, "y": 8},
// 	],
// 	"mines": [],
//  "activatedMine": null
// }
//
// as is the most direct way we can encode out data given our internal representation

type ExportedClientBoard struct {
	FullExport bool                    `json:"fullExport"`
	State      string                  `json:"state"`
	SafePoints map[string][]game.Point `json:"safePoints"`
	Flags      []game.Point            `json:"flags"`
	Mines      []game.Point            `json:"mines"`
	Activated  *game.Point             `json:"activatedMine"`
}

func (gm *GamesManager) safePointsByNumber(points map[game.Point]int) map[string][]game.Point {
	byNumber := make(map[string][]game.Point)
	for p, n := range points {
		num := strconv.Itoa(n)
		byNumber[num] = append(byNumber[num], p)
	}
	return byNumber
}

func (gm *GamesManager) ExportClientBoard(g *game.Game) (*ExportedClientBoard, error) {
	if g == nil {
		return nil, errors.Errorf("Nil game")
	}
	cl := g.ClientBoard()
	if cl == nil {
		return nil, errors.Errorf("Client board not found")
	}
	safe := gm.safePointsByNumber(cl.SafePoints())
	exported := &ExportedClientBoard{
		FullExport: true,
		SafePoints: safe,
		State:      g.State(),
		Flags:      cl.Flags(),
		Mines:      cl.Mines(),
		Activated:  cl.Activated(),
	}
	return exported, nil
}
