package game

type ClientBoard struct {
	safePoints map[Point]int
	flags      map[Point]bool
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
// 	"bombs": []
// }
//
// as is the most direct way we can encode out data given our internal representation
