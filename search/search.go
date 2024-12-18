package search

import (
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/Greeshmanth1909/shadowfax/eval"
	"math"
	"time"
)

func CheckUp() {
	// check if time up or interrupt from GUI
}
func IsRepetition(brd *board.S_Board) bool {
	pKey := brd.PosKey

	for i := brd.HisPly - brd.FiftyMove; i < brd.HisPly; i++ {
		if brd.History[i].PosKey == pKey {
			return true
		}
	}
	return false
}

func SearchPositions(brd *board.S_Board, info *board.S_SearchInfo) {
	var bestMove uint32
	var bestScore int
	var currentDepth int
	// var pvMoves int
	// var pvNum int

	bestScore = int(math.Inf(-1))
	ClearForSearch(brd, info)

	for currentDepth = 0; currentDepth < info.Depth; currentDepth++ {
		bestScore = AlphaBeta(int(math.Inf(-1)), int(math.Inf(1)), currentDepth, 1, brd, info)
		// pvMoves = GetPvLine(currentDepth, brd)
		bestMove = brd.PvArray[0]
		fmt.Printf("score %v depth %v nodes %v\n", bestScore, currentDepth, info.Nodes)
		fmt.Printf("move ")
		var m eval.S_Move
		m.Move = bestMove
		eval.PrintMove(&m)
		for i := 0; i < brd.PvTable.NumEntries; i++ {
			val := brd.PvArray[i]
			var mv eval.S_Move
			mv.Move = val
			eval.PrintMove(&mv)
		}
	}

}

func ClearForSearch(brd *board.S_Board, info *board.S_SearchInfo) {
	for i := 0; i < 13; i++ {
		for j := 0; j < board.BrdSqrNum; j++ {
			brd.SearchHistoryArray[i][j] = 0
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < board.MAXDEPTH; j++ {
			brd.SearchKillers[i][j] = 0
		}
	}

	// Clear pvtable
	for key := range brd.PvTable.PvTableEntries {
		delete(brd.PvTable.PvTableEntries, key)
	}
	brd.PvTable.NumEntries = 0
	brd.Ply = 0

	info.StartTime = time.Now()
	info.Stopped = 0
	info.Nodes = 0
}

func Quiescence(alpha, beta int, brd *board.S_Board, info *board.S_SearchInfo) (score int) {
	return
}
func AlphaBeta(alpha, beta, depth, doNull int, brd *board.S_Board, info *board.S_SearchInfo) (score int) {
	return
}
