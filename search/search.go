package search

import (
	"github.com/Greeshmanth1909/shadowfax/board"
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

}

func Quiescence(alpha, beta int, brd *board.S_Board, info *board.S_SearchInfo) (score int) {
	return
}
func AlphaBeta(alpha, beta, depth, doNull int, brd *board.S_Board, info *board.S_SearchInfo) (score int) {
	return
}
