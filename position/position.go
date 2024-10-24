package position

import (
	"github.com/Greeshmanth1909/shadowfax/board"
)

func ResetBoard(brd *board.S_Board) {
	for i := range brd.Pieces {
		brd.Pieces[i] = int(board.OFFBOARD)
	}

	for i := range brd.Pieces {
		brd.Pieces[board.Square64to120[i]] = int(board.EMPTY)
	}

	for i := 0; i < 3; i++ {
		brd.BigPiece[i] = 0
		brd.MinPiece[i] = 0
		brd.MajPiece[i] = 0
		brd.Pawns[i] = uint64(0)
	}

	brd.Side = board.BOTH
	brd.EnP = board.NO_SQ
	brd.FiftyMove = 0
	brd.CastlePerm = 0
	brd.Ply = 0
	brd.HisPly = 0
	brd.PosKey = uint64(0)
	brd.KingSquare[int(board.WHITE)] = 0
	brd.KingSquare[int(board.BLACK)] = 0

	for i := 0; i < 13; i++ {
		brd.PieceNum[i] = 0
	}
}
