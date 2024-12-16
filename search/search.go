package search

import (
	"github.com/Greeshmanth1909/shadowfax/board"
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
