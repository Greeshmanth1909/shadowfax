package eval

import (
	"github.com/Greeshmanth1909/shadowfax/board"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

const FLAGENP = uint32(1) << 18
const FLAGPS = uint32(1) << 19
const FLAGC = uint32(1) << 24

const MAXPOSITIONMOVES = 256

type S_MoveList struct {
	MoveList [MAXPOSITIONMOVES]S_Move
	Count    int
}

func AddQuietMove(brd *board.S_Board, move uint32, list *S_MoveList) {
	list.MoveList[list.Count].Move = move
	list.MoveList[list.Count].Score = 0
	list.Count++
}

func AddCaptureMove(brd *board.S_Board, move uint32, list *S_MoveList) {
	list.MoveList[list.Count].Move = move
	list.MoveList[list.Count].Score = 0
	list.Count++
}
func AddEnPassantMove(brd *board.S_Board, move uint32, list *S_MoveList) {
	list.MoveList[list.Count].Move = move
	list.MoveList[list.Count].Score = 0
	list.Count++
}

func AddWhitePawnCapMove(brd *board.S_Board, from, to board.Square, capt board.Piece, list *S_MoveList) {
	if board.RankArr[from] == board.RANK_7 {
		AddCaptureMove(brd, Move(from, to, capt, board.Wq, 0), list)
		AddCaptureMove(brd, Move(from, to, capt, board.Wb, 0), list)
		AddCaptureMove(brd, Move(from, to, capt, board.Wn, 0), list)
		AddCaptureMove(brd, Move(from, to, capt, board.Wr, 0), list)
	} else {
		AddCaptureMove(brd, Move(from, to, capt, board.EMPTY, 0), list)
	}
}

func AddWhitePawnMove(brd *board.S_Board, from, to board.Square, list *S_MoveList) {
	if board.RankArr[from] == board.RANK_7 {
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Wq, 0), list)
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Wr, 0), list)
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Wk, 0), list)
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Wb, 0), list)
	} else {
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.EMPTY, 0), list)
	}
}

func GenerateAllMoves(brd *board.S_Board, list *S_MoveList) {
	side := brd.Side

	if side == board.WHITE {
		for _, sq := range brd.PList[board.Wp] {
			if sq+9 != int(board.OFFBOARD) {
				if board.PieceCol[brd.Pieces[sq+9]] == board.BLACK {
					AddWhitePawnCapMove(brd, board.Square(sq), board.Square(sq+9), brd.Pieces[sq+9], list)
				}
			}
			if sq+11 != int(board.OFFBOARD) {
				if board.PieceCol[brd.Pieces[sq+11]] == board.BLACK {
					AddWhitePawnCapMove(brd, board.Square(sq), board.Square(sq+11), brd.Pieces[sq+11], list)
				}
			}
			if brd.Pieces[sq+10] == board.EMPTY {
				AddWhitePawnMove(brd, board.Square(sq), board.Square(sq+10), list)
			}
			if board.RankArr[sq] == board.RANK_2 && brd.Pieces[sq+10] == board.EMPTY && brd.Pieces[sq+20] == board.EMPTY {
				AddQuietMove(brd, Move(board.Square(sq), board.Square(sq+20), board.EMPTY, board.EMPTY, FLAGPS), list)
			}

			if sq+11 == int(brd.EnP) {
				AddQuietMove(brd, Move(board.Square(sq), board.Square(sq+11), board.EMPTY, board.EMPTY, FLAGENP), list)
			}
		}
	}
}
