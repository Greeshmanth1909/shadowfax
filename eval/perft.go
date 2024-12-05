package eval

import (
	"github.com/Greeshmanth1909/shadowfax/board"
)

func PerftTest(depth int, brd *board.S_Board) int {
	leafNodes := 0
	var list S_MoveList

	if depth == 0 {
		leafNodes++
		return leafNodes
	}
	depth--
	GenerateAllMoves(brd, &list)
	for i := 0; i < list.Count; i++ {
		mv := list.MoveList[i]
		if MakeMove(brd, &mv) {
			leafNodes += PerftTest(depth, brd)
			TakeMove(brd)
		} else {
			continue
		}
	}
	return leafNodes
}
