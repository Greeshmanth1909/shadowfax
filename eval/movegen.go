package eval

import (
	"fmt"
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

var LoopSlidingPieces = [8]board.Piece{board.Wb, board.Wr, board.Wq, board.EMPTY, board.Bb, board.Br, board.Bq, board.EMPTY}
var LoopSlidingPiecesIndex = [2]int{0, 4}
var NonSlidingPieces = [6]board.Piece{board.Wn, board.Wk, board.EMPTY, board.Bn, board.Bk, board.EMPTY}
var NonSlidingPiecesIndex = [2]int{0, 3}

var PieceDir = [13][8]int{{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{12, 8, 21, 19, -12, -8, -21, -19},
	{11, 9, -11, -9, 0, 0, 0, 0},
	{-1, -10, 1, 10, 0, 0, 0, 0},
	{10, 1, 11, 9, -10, -1, -11, -9},
	{10, 1, 11, 9, -10, -1, -11, -9},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{12, 8, 21, 19, -12, -8, -21, -19},
	{11, 9, -11, -9, 0, 0, 0, 0},
	{-1, -10, 1, 10, 0, 0, 0, 0},
	{10, 1, 11, 9, -10, -1, -11, -9},
	{10, 1, 11, 9, -10, -1, -11, -9},
}

var PieceDirNum = [13]int{0, 0, 8, 4, 4, 8, 8, 0, 8, 4, 4, 8, 8}

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
	if from == board.OFFBOARD || to == board.OFFBOARD {
		log.Fatalf("offboard square given\n")
	}
	if capt == board.EMPTY {
		log.Fatalf("invalid capt piece (%v)", capt)
	}

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
	if from == board.OFFBOARD || to == board.OFFBOARD {
		log.Fatalf("offboard square given\n")
	}

	if board.RankArr[from] == board.RANK_7 {
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Wq, 0), list)
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Wr, 0), list)
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Wn, 0), list)
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Wb, 0), list)
	} else {
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.EMPTY, 0), list)
	}
}

func AddBlackPawnCapMove(brd *board.S_Board, from, to board.Square, capt board.Piece, list *S_MoveList) {
	if from == board.OFFBOARD || to == board.OFFBOARD {
		log.Fatalf("offboard square given\n")
	}
	if capt == board.EMPTY {
		log.Fatalf("invalid capt piece (%v)", capt)
	}

	if board.RankArr[from] == board.RANK_2 {
		AddCaptureMove(brd, Move(from, to, capt, board.Bq, 0), list)
		AddCaptureMove(brd, Move(from, to, capt, board.Bb, 0), list)
		AddCaptureMove(brd, Move(from, to, capt, board.Bn, 0), list)
		AddCaptureMove(brd, Move(from, to, capt, board.Br, 0), list)
	} else {
		AddCaptureMove(brd, Move(from, to, capt, board.EMPTY, 0), list)
	}
}

func AddBlackPawnMove(brd *board.S_Board, from, to board.Square, list *S_MoveList) {
	if from == board.OFFBOARD || to == board.OFFBOARD {
		log.Fatalf("offboard square given\n")
	}

	if board.RankArr[from] == board.RANK_2 {
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Bq, 0), list)
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Br, 0), list)
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Bn, 0), list)
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.Bb, 0), list)
	} else {
		AddQuietMove(brd, Move(from, to, board.EMPTY, board.EMPTY, 0), list)
	}
}

func GenerateAllMoves(brd *board.S_Board, list *S_MoveList) {
	board.CheckBoard(brd)
	side := brd.Side

	if side == board.WHITE {
		for _, sq := range brd.PList[board.Wp] {
			if brd.Pieces[sq+9] != board.Piece(board.OFFBOARD) {
				if board.PieceCol[brd.Pieces[sq+9]] == board.BLACK {
					AddWhitePawnCapMove(brd, board.Square(sq), board.Square(sq+9), brd.Pieces[sq+9], list)
				}
			}
			if brd.Pieces[sq+11] != board.Piece(board.OFFBOARD) {
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
			if board.Square(sq+11) == (brd.EnP) {
				AddQuietMove(brd, Move(board.Square(sq), board.Square(sq+11), board.EMPTY, board.EMPTY, FLAGENP), list)
			}
			if board.Square(sq+9) == (brd.EnP) {
				AddQuietMove(brd, Move(board.Square(sq), board.Square(sq+9), board.EMPTY, board.EMPTY, FLAGENP), list)
			}
		}

		if brd.CastlePerm&int(board.WKCT) != 0 {
			if brd.Pieces[board.G1] == board.EMPTY && brd.Pieces[board.F1] == board.EMPTY {
				if !SquareAttacked(board.E1, board.BLACK, brd) && !SquareAttacked(board.F1, board.BLACK, brd) {
					AddQuietMove(brd, Move(board.E1, board.G1, board.EMPTY, board.EMPTY, FLAGC), list)
				}
			}
		}

		if brd.CastlePerm&int(board.WQCT) != 0 {
			if brd.Pieces[board.D1] == board.EMPTY && brd.Pieces[board.C1] == board.EMPTY && brd.Pieces[board.B1] == board.EMPTY {
				if !SquareAttacked(board.E1, board.BLACK, brd) && !SquareAttacked(board.D1, board.BLACK, brd) {
					AddQuietMove(brd, Move(board.E1, board.C1, board.EMPTY, board.EMPTY, FLAGC), list)
				}
			}
		}

	}
	if side == board.BLACK {
		for _, sq := range brd.PList[board.Bp] {
			if sq <= 0 {
				break
			}
			if brd.Pieces[sq-9] != board.Piece(board.OFFBOARD) {
				if board.PieceCol[brd.Pieces[sq-9]] == board.WHITE {
					AddBlackPawnCapMove(brd, board.Square(sq), board.Square(sq-9), brd.Pieces[sq-9], list)
				}
			}
			if brd.Pieces[sq-11] != board.Piece(board.OFFBOARD) {
				if board.PieceCol[brd.Pieces[sq-11]] == board.WHITE {
					AddBlackPawnCapMove(brd, board.Square(sq), board.Square(sq-11), brd.Pieces[sq-11], list)
				}
			}
			if brd.Pieces[sq-10] == board.EMPTY {
				AddBlackPawnMove(brd, board.Square(sq), board.Square(sq-10), list)
			}
			if board.RankArr[sq] == board.RANK_7 && brd.Pieces[sq-10] == board.EMPTY && brd.Pieces[sq-20] == board.EMPTY {
				AddQuietMove(brd, Move(board.Square(sq), board.Square(sq-20), board.EMPTY, board.EMPTY, FLAGPS), list)
			}
			if board.Square(sq-11) == (brd.EnP) {
				AddQuietMove(brd, Move(board.Square(sq), board.Square(sq-11), board.EMPTY, board.EMPTY, FLAGENP), list)
			}
			if board.Square(sq-9) == (brd.EnP) {
				AddQuietMove(brd, Move(board.Square(sq), board.Square(sq-9), board.EMPTY, board.EMPTY, FLAGENP), list)
			}
		}
		if brd.CastlePerm&int(board.BKCT) != 0 {
			if brd.Pieces[board.G8] == board.EMPTY && brd.Pieces[board.F8] == board.EMPTY {
				if !SquareAttacked(board.E8, board.WHITE, brd) && !SquareAttacked(board.F8, board.WHITE, brd) {
					fmt.Println("moveGen: BKCT")
					AddQuietMove(brd, Move(board.E8, board.G8, board.EMPTY, board.EMPTY, FLAGC), list)
				}
			}
		}

		if brd.CastlePerm&int(board.BQCT) != 0 {
			if brd.Pieces[board.D8] == board.EMPTY && brd.Pieces[board.C8] == board.EMPTY && brd.Pieces[board.B8] == board.EMPTY {
				if !SquareAttacked(board.E8, board.WHITE, brd) && !SquareAttacked(board.D8, board.WHITE, brd) {
					fmt.Println("moveGen: BQCT")
					AddQuietMove(brd, Move(board.E8, board.C8, board.EMPTY, board.EMPTY, FLAGC), list)
				}
			}
		}
	}

	// Sliding Pieces
	startIndex := LoopSlidingPiecesIndex[side]
	piece := LoopSlidingPieces[startIndex]
	for piece != board.EMPTY {
		for _, sq := range brd.PList[piece] {
			if sq == 0 {
				break
			}
			fmt.Printf("slider pieceInd: %v, piece: %v\n", sq, piece)
			if !board.ValidatePiece(brd.Pieces[sq]) {
				log.Fatalf("Invalid slider piece (%v)", brd.Pieces[sq])
			}
			pieceDirArray := PieceDir[brd.Pieces[sq]]
			for _, offset := range pieceDirArray {
				if offset == 0 {
					continue
				}
				toSq := sq + offset
				for brd.Pieces[toSq] != board.Piece(board.OFFBOARD) {
					if board.PieceCol[brd.Pieces[toSq]] == side {
						break
					}
					if brd.Pieces[toSq] == board.EMPTY {
						AddQuietMove(brd, Move(board.Square(sq), board.Square(toSq), board.EMPTY, board.EMPTY, 0), list)
						toSq += offset
						continue
					}
					AddCaptureMove(brd, Move(board.Square(sq), board.Square(toSq), brd.Pieces[toSq], board.EMPTY, 0), list)
					break
				}
			}
		}
		startIndex++
		piece = LoopSlidingPieces[startIndex]
	}

	// NonSlider Pieces
	startIndex = NonSlidingPiecesIndex[side]
	piece = NonSlidingPieces[startIndex]
	for piece != board.EMPTY {
		for _, sq := range brd.PList[piece] {
			if sq == 0 {
				break
			}
			fmt.Printf("non slide pieceInd: %v, piece: %v\n", sq, piece)
			if !board.ValidatePiece(brd.Pieces[sq]) {
				log.Fatalf("Invalid slider piece (%v)", brd.Pieces[sq])
			}

			pieceDirArray := PieceDir[brd.Pieces[sq]]
			for _, offset := range pieceDirArray {
				toSq := sq + offset
				if brd.Pieces[toSq] == board.Piece(board.OFFBOARD) {
					continue
				}
				if board.PieceCol[brd.Pieces[toSq]] == side {
					continue
				}
				if brd.Pieces[toSq] == board.EMPTY {
					AddQuietMove(brd, Move(board.Square(sq), board.Square(toSq), board.EMPTY, board.EMPTY, 0), list)
					continue
				}
				AddCaptureMove(brd, Move(board.Square(sq), board.Square(toSq), brd.Pieces[toSq], board.EMPTY, 0), list)
			}
		}
		startIndex++
		piece = NonSlidingPieces[startIndex]
	}
}

func MakeMove(brd *board.S_Board, mv *S_Move) {
	frm, to, capt, pro, flag := getMove(mv)
	side := brd.Side
	piece := brd.Pieces[frm]
	pieceAtTo := brd.Pieces[to]
	if piece == board.EMPTY || piece == board.Piece(board.OFFBOARD) {
		log.Fatalf("trying to move empty piece frm (%v)", frm)
	}
	if pieceAtTo == board.Piece(board.OFFBOARD) {
		log.Fatalf("Invalid to Sq %v", to)
	}
	if board.PieceCol[piece] != side {
		log.Fatalf("Invalid Piece movement")
	}

	// Remove piece from frm square and place it at to square
	brd.Pieces[frm] = board.EMPTY
	brd.Pieces[to] = piece

	// Check for capture
	if capt != 0 {
		if pieceAtTo != capt {
			log.Fatalf("Invalid Capture")
		}
		// Update piece list, this might cause problems
		for i, sq := range brd.PList[pieceAtTo] {
			if sq == int(to) {
				brd.PList[pieceAtTo][i] = int(board.EMPTY)
				brd.PieceNum[pieceAtTo]--
				if board.BigPiece[pieceAtTo] {
					brd.MajPiece[side^1]--
				} else if board.MinPiece[pieceAtTo] {
					brd.MinPiece[side^1]--
				}
				brd.Material[side] += board.PieceVal[pieceAtTo]
				break
			}
		}
	}

	// Promotion, if any
	if pro != 0 {
		brd.Pieces[to] = pro
		for i, val := range brd.PList[piece] {
			if val == int(frm) {
				brd.PList[piece][i] = int(board.EMPTY)
			}
		}
		brd.Material[side] -= board.PieceVal[piece]
		brd.Material[side] += board.PieceVal[pro]
		brd.PieceNum[pro]++
		brd.PList[pro][brd.PieceNum[pro]] = int(to)
	}

	if flag == 0 {
		fmt.Println("heh")
	}
}
