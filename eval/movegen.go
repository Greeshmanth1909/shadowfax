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

// use arrays to infer to most valuable victim least valuable attacker
var vicArray = [13]int{0, 100, 200, 300, 400, 500, 600, 100, 200, 300, 400, 500, 600}
var MVVLVA [13][13]int

const FLAGENP = uint32(1) << 18
const FLAGPS = uint32(1) << 19
const FLAGC = uint32(1) << 24

const MAXPOSITIONMOVES = 256

// InitMvvLVa function initialises mvvlva array
func InitMvvLva() {
	for victim := board.Wp; victim < board.Bk; victim++ {
		for attacker := board.Wp; attacker < board.Bk; attacker++ {
			MVVLVA[victim][attacker] = vicArray[victim] + 6 - (vicArray[attacker] / 100)
		}
	}
}

var CastlePerm = [120]int{
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 13, 15, 15, 15, 12, 15, 15, 14, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 7, 15, 15, 15, 3, 15, 15, 11, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
}

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

func addQuietMove(brd *board.S_Board, move uint32, list *S_MoveList) {
	list.MoveList[list.Count].Move = move

    var mv S_Move
    mv.Move = move
    if brd.SearchKillers[0][brd.Ply] == move {
        list.MoveList[list.Count].Score = 900000
    } else if brd.SearchKillers[1][brd.Ply] == move {
        list.MoveList[list.Count].Score = 800000
    } else {
        list.MoveList[list.Count].Score = int(brd.SearchHistoryArray[brd.Pieces[GetFromSquare(&mv)]][GetToSquare(&mv)])
    }

	list.Count++
}

func addCaptureMove(brd *board.S_Board, move uint32, list *S_MoveList) {
	list.MoveList[list.Count].Move = move
	var m S_Move
	m.Move = move
	list.MoveList[list.Count].Score = MVVLVA[GetCapturedPiece(&m)][brd.Pieces[GetFromSquare(&m)]] + 1000000
	list.Count++
}

func addEnPassantMove(brd *board.S_Board, move uint32, list *S_MoveList) {
	list.MoveList[list.Count].Move = move
	list.MoveList[list.Count].Score = 105 + 1000000
	list.Count++
}

func addWhitePawnCapMove(brd *board.S_Board, from, to board.Square, capt board.Piece, list *S_MoveList) {
	if from == board.OFFBOARD || to == board.OFFBOARD {
		log.Fatalf("offboard square given\n")
	}
	if capt == board.EMPTY {
		log.Fatalf("invalid capt piece (%v)", capt)
	}

	if board.RankArr[from] == board.RANK_7 {
		addCaptureMove(brd, Move(from, to, capt, board.Wq, 0), list)
		addCaptureMove(brd, Move(from, to, capt, board.Wb, 0), list)
		addCaptureMove(brd, Move(from, to, capt, board.Wn, 0), list)
		addCaptureMove(brd, Move(from, to, capt, board.Wr, 0), list)
	} else {
		addCaptureMove(brd, Move(from, to, capt, board.EMPTY, 0), list)
	}
}

func addWhitePawnMove(brd *board.S_Board, from, to board.Square, list *S_MoveList) {
	if from == board.OFFBOARD || to == board.OFFBOARD {
		log.Fatalf("offboard square given\n")
	}

	if board.RankArr[from] == board.RANK_7 {
		addQuietMove(brd, Move(from, to, board.EMPTY, board.Wq, 0), list)
		addQuietMove(brd, Move(from, to, board.EMPTY, board.Wr, 0), list)
		addQuietMove(brd, Move(from, to, board.EMPTY, board.Wn, 0), list)
		addQuietMove(brd, Move(from, to, board.EMPTY, board.Wb, 0), list)
	} else {
		addQuietMove(brd, Move(from, to, board.EMPTY, board.EMPTY, 0), list)
	}
}

func addBlackPawnCapMove(brd *board.S_Board, from, to board.Square, capt board.Piece, list *S_MoveList) {
	if from == board.OFFBOARD || to == board.OFFBOARD {
		log.Fatalf("offboard square given\n")
	}
	if capt == board.EMPTY {
		log.Fatalf("invalid capt piece (%v)", capt)
	}

	if board.RankArr[from] == board.RANK_2 {
		addCaptureMove(brd, Move(from, to, capt, board.Bq, 0), list)
		addCaptureMove(brd, Move(from, to, capt, board.Bb, 0), list)
		addCaptureMove(brd, Move(from, to, capt, board.Bn, 0), list)
		addCaptureMove(brd, Move(from, to, capt, board.Br, 0), list)
	} else {
		addCaptureMove(brd, Move(from, to, capt, board.EMPTY, 0), list)
	}
}

func addBlackPawnMove(brd *board.S_Board, from, to board.Square, list *S_MoveList) {
	if from == board.OFFBOARD || to == board.OFFBOARD {
		log.Fatalf("offboard square given\n")
	}

	if board.RankArr[from] == board.RANK_2 {
		addQuietMove(brd, Move(from, to, board.EMPTY, board.Bq, 0), list)
		addQuietMove(brd, Move(from, to, board.EMPTY, board.Br, 0), list)
		addQuietMove(brd, Move(from, to, board.EMPTY, board.Bn, 0), list)
		addQuietMove(brd, Move(from, to, board.EMPTY, board.Bb, 0), list)
	} else {
		addQuietMove(brd, Move(from, to, board.EMPTY, board.EMPTY, 0), list)
	}
}

func GenerateAllMoves(brd *board.S_Board, list *S_MoveList) {
	// board.CheckBoard(brd)
	side := brd.Side

	if side == board.WHITE {
		for pnum := 0; pnum < brd.PieceNum[board.Wp]; pnum++ {
			sq := brd.PList[board.Wp][pnum]
			if brd.Pieces[sq+9] != board.Piece(board.OFFBOARD) {
				if board.PieceCol[brd.Pieces[sq+9]] == board.BLACK {
					addWhitePawnCapMove(brd, board.Square(sq), board.Square(sq+9), brd.Pieces[sq+9], list)
				}
			}
			if brd.Pieces[sq+11] != board.Piece(board.OFFBOARD) {
				if board.PieceCol[brd.Pieces[sq+11]] == board.BLACK {
					addWhitePawnCapMove(brd, board.Square(sq), board.Square(sq+11), brd.Pieces[sq+11], list)
				}
			}
			if brd.Pieces[sq+10] == board.EMPTY {
				addWhitePawnMove(brd, board.Square(sq), board.Square(sq+10), list)
			}
			if board.RankArr[sq] == board.RANK_2 && brd.Pieces[sq+10] == board.EMPTY && brd.Pieces[sq+20] == board.EMPTY {
				addQuietMove(brd, Move(board.Square(sq), board.Square(sq+20), board.EMPTY, board.EMPTY, FLAGPS), list)
			}
			if brd.EnP != board.NO_SQ {
				if board.Square(sq+11) == (brd.EnP) {
					addEnPassantMove(brd, Move(board.Square(sq), board.Square(sq+11), board.EMPTY, board.EMPTY, FLAGENP), list)
				}
				if board.Square(sq+9) == (brd.EnP) {
					addEnPassantMove(brd, Move(board.Square(sq), board.Square(sq+9), board.EMPTY, board.EMPTY, FLAGENP), list)
				}
			}
		}

		if brd.CastlePerm&int(board.WKCT) != 0 {
			if brd.Pieces[board.G1] == board.EMPTY && brd.Pieces[board.F1] == board.EMPTY {
				if !SquareAttacked(board.E1, board.BLACK, brd) && !SquareAttacked(board.F1, board.BLACK, brd) {
					addQuietMove(brd, Move(board.E1, board.G1, board.EMPTY, board.EMPTY, FLAGC), list)
				}
			}
		}

		if brd.CastlePerm&int(board.WQCT) != 0 {
			if brd.Pieces[board.D1] == board.EMPTY && brd.Pieces[board.C1] == board.EMPTY && brd.Pieces[board.B1] == board.EMPTY {
				if !SquareAttacked(board.E1, board.BLACK, brd) && !SquareAttacked(board.D1, board.BLACK, brd) {
					addQuietMove(brd, Move(board.E1, board.C1, board.EMPTY, board.EMPTY, FLAGC), list)
				}
			}
		}

	}
	if side == board.BLACK {
		for pnum := 0; pnum < brd.PieceNum[board.Bp]; pnum++ {
			sq := brd.PList[board.Bp][pnum]
			if sq <= 0 {
				break
			}
			if brd.Pieces[sq-9] != board.Piece(board.OFFBOARD) {
				if board.PieceCol[brd.Pieces[sq-9]] == board.WHITE {
					addBlackPawnCapMove(brd, board.Square(sq), board.Square(sq-9), brd.Pieces[sq-9], list)
				}
			}
			if brd.Pieces[sq-11] != board.Piece(board.OFFBOARD) {
				if board.PieceCol[brd.Pieces[sq-11]] == board.WHITE {
					addBlackPawnCapMove(brd, board.Square(sq), board.Square(sq-11), brd.Pieces[sq-11], list)
				}
			}
			if brd.Pieces[sq-10] == board.EMPTY {
				addBlackPawnMove(brd, board.Square(sq), board.Square(sq-10), list)
			}
			if board.RankArr[sq] == board.RANK_7 && brd.Pieces[sq-10] == board.EMPTY && brd.Pieces[sq-20] == board.EMPTY {
				addQuietMove(brd, Move(board.Square(sq), board.Square(sq-20), board.EMPTY, board.EMPTY, FLAGPS), list)
			}
			if brd.EnP != board.NO_SQ {
				if board.Square(sq-11) == (brd.EnP) {
					addEnPassantMove(brd, Move(board.Square(sq), board.Square(sq-11), board.EMPTY, board.EMPTY, FLAGENP), list)
				}
				if board.Square(sq-9) == (brd.EnP) {
					addEnPassantMove(brd, Move(board.Square(sq), board.Square(sq-9), board.EMPTY, board.EMPTY, FLAGENP), list)
				}
			}
		}
		if brd.CastlePerm&int(board.BKCT) != 0 {
			if brd.Pieces[board.G8] == board.EMPTY && brd.Pieces[board.F8] == board.EMPTY {
				if !SquareAttacked(board.E8, board.WHITE, brd) && !SquareAttacked(board.F8, board.WHITE, brd) {
					addQuietMove(brd, Move(board.E8, board.G8, board.EMPTY, board.EMPTY, FLAGC), list)
				}
			}
		}

		if brd.CastlePerm&int(board.BQCT) != 0 {
			if brd.Pieces[board.D8] == board.EMPTY && brd.Pieces[board.C8] == board.EMPTY && brd.Pieces[board.B8] == board.EMPTY {
				if !SquareAttacked(board.E8, board.WHITE, brd) && !SquareAttacked(board.D8, board.WHITE, brd) {
					addQuietMove(brd, Move(board.E8, board.C8, board.EMPTY, board.EMPTY, FLAGC), list)
				}
			}
		}
	}

	// Sliding Pieces
	startIndex := LoopSlidingPiecesIndex[side]
	piece := LoopSlidingPieces[startIndex]
	for piece != board.EMPTY {
		for pnum := 0; pnum < brd.PieceNum[piece]; pnum++ {
			sq := brd.PList[piece][pnum]
			if brd.Pieces[sq] == board.EMPTY {
				continue
			}
			if sq == 0 {
				break
			}
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
						addQuietMove(brd, Move(board.Square(sq), board.Square(toSq), board.EMPTY, board.EMPTY, 0), list)
						toSq += offset
						continue
					}
					addCaptureMove(brd, Move(board.Square(sq), board.Square(toSq), brd.Pieces[toSq], board.EMPTY, 0), list)
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
		for pnum := 0; pnum < brd.PieceNum[piece]; pnum++ {
			sq := brd.PList[piece][pnum]
			if brd.Pieces[sq] == board.EMPTY {
				continue
			}
			if sq == 0 {
				break
			}
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
					addQuietMove(brd, Move(board.Square(sq), board.Square(toSq), board.EMPTY, board.EMPTY, 0), list)
					continue
				}
				addCaptureMove(brd, Move(board.Square(sq), board.Square(toSq), brd.Pieces[toSq], board.EMPTY, 0), list)
			}
		}
		startIndex++
		piece = NonSlidingPieces[startIndex]
	}
}

func clearPiece(sq board.Square, brd *board.S_Board) {
	piece := brd.Pieces[sq]
	if piece == board.EMPTY || piece == board.Piece(board.OFFBOARD) {
		log.Fatalf("trying to move empty piece frm (%v)", sq)
	}
	col := board.PieceCol[piece]
	//index := 0
	t_pieceNum := -1
	hashPiece(brd, piece, sq)

	brd.Pieces[sq] = board.EMPTY
	brd.Material[col] -= board.PieceVal[piece]

	if board.BigPiece[piece] {
		brd.BigPiece[col]--
		if board.MajPiece[piece] {
			brd.MajPiece[col]--
		} else {
			brd.MinPiece[col]--
		}
	} else {
		board.ClearBit(board.Square120to64[sq], &brd.Pawns[col])
		board.ClearBit(board.Square120to64[sq], &brd.Pawns[board.BOTH])
	}

	// Clear piece from PList
	for i := 0; i < brd.PieceNum[piece]; i++ {
		val := brd.PList[piece][i]
		if val == int(sq) {
			t_pieceNum = i
			break
		}
	}
	if t_pieceNum == -1 {
		log.Fatalf("t_piecenum is -1 (%v)", t_pieceNum)
	}

	brd.PieceNum[piece]--
	brd.PList[piece][t_pieceNum] = brd.PList[piece][brd.PieceNum[piece]]
}

func addPiece(sq board.Square, brd *board.S_Board, pc board.Piece) {
	piece := brd.Pieces[sq]
	col := board.PieceCol[pc]

	if piece == board.Piece(board.OFFBOARD) {
		log.Fatalf("Invalid to square (%v)", piece)
	}
	if piece != board.EMPTY {
		log.Fatalf("Square not empty (%v)", piece)
	}

	hashPiece(brd, pc, sq)
	brd.Pieces[sq] = pc

	if board.BigPiece[pc] {
		brd.BigPiece[col]++
		if board.MajPiece[pc] {
			brd.MajPiece[col]++
		} else {
			brd.MinPiece[col]++
		}
	} else {
		board.SetBit(board.Square120to64[sq], &brd.Pawns[col])
		board.SetBit(board.Square120to64[sq], &brd.Pawns[board.BOTH])
	}

	brd.Material[col] += board.PieceVal[pc]
	brd.PList[pc][brd.PieceNum[pc]] = int(sq)
	brd.PieceNum[pc]++
}

func movePiece(from, to board.Square, brd *board.S_Board) {
	if brd.Pieces[from] == board.Piece(board.OFFBOARD) || brd.Pieces[to] == board.Piece(board.OFFBOARD) {
		log.Fatalf("movePiece: invalid from/to square (from: %v) (to: %v)", brd.Pieces[from], brd.Pieces[to])
	}
	piece := brd.Pieces[from]
	col := board.PieceCol[piece]

	hashPiece(brd, piece, from)
	brd.Pieces[from] = board.EMPTY

	hashPiece(brd, piece, to)
	brd.Pieces[to] = piece

	flag := false

	if !board.BigPiece[piece] {
		board.ClearBit(board.Square120to64[from], &brd.Pawns[col])
		board.ClearBit(board.Square120to64[from], &brd.Pawns[board.BOTH])
		board.SetBit(board.Square120to64[to], &brd.Pawns[col])
		board.SetBit(board.Square120to64[to], &brd.Pawns[board.BOTH])
	}

	for i := 0; i < brd.PieceNum[piece]; i++ {
		val := brd.PList[piece][i]
		if val == int(from) {
			brd.PList[piece][i] = int(to)
			flag = true
			break
		}
	}
	if !flag {
		log.Fatalf("movePiece: Piece not found in PList")
	}
}

func MakeMove(brd *board.S_Board, mv *S_Move) bool {
	from, to, capt, pro, flag := getMove(mv)

	if brd.Pieces[from] == board.EMPTY || brd.Pieces[from] == board.Piece(board.OFFBOARD) {
		log.Fatalf("MakeMove: invalid from square")
	}
	if brd.Pieces[to] == board.Piece(board.OFFBOARD) {
		log.Fatalf("MakeMove: invalid to square")
	}

	side := brd.Side
	brd.History[brd.HisPly].PosKey = brd.PosKey

	if flag == FLAGENP {
		if side == board.WHITE {
			clearPiece(to-board.Square(10), brd)
		} else {
			clearPiece(to+board.Square(10), brd)
		}
	}
	/* Im assuming castling will be represented with king move and castle flag. That is the case so make
	   move will move the king later on so the rook needs to move now
	*/

	if flag == FLAGC {
		switch to {
		case board.C1:
			movePiece(board.A1, board.D1, brd)
		case board.C8:
			movePiece(board.A8, board.D8, brd)
		case board.G1:
			movePiece(board.H1, board.F1, brd)
		case board.G8:
			movePiece(board.H8, board.F8, brd)
		default:
			log.Fatalf("Invalid to square in castling (%v)", to)
		}
	}

	if brd.EnP != board.NO_SQ {
		hashEnP(brd)
	}

	hashC(brd)
	brd.History[brd.HisPly].Move = mv.Move
	brd.History[brd.HisPly].CastlePerm = brd.CastlePerm
	brd.History[brd.HisPly].EnP = brd.EnP
	brd.History[brd.HisPly].FiftyMove = brd.FiftyMove

	brd.CastlePerm &= CastlePerm[from]
	brd.CastlePerm &= CastlePerm[to]
	brd.EnP = board.NO_SQ

	hashC(brd)
	brd.FiftyMove++

	if capt != board.EMPTY {
		if capt == board.Piece(board.OFFBOARD) {
			log.Fatalf("MakeMove: Invalid capture piece")
		}
		clearPiece(to, brd)
		brd.FiftyMove = 0
	}
	brd.Ply++
	brd.HisPly++

	if board.PiecePawn[brd.Pieces[from]] {
		brd.FiftyMove = 0
		if flag == FLAGPS {
			if side == board.WHITE {
				brd.EnP = from + board.Square(10)
				if board.RankArr[brd.EnP] != board.RANK_3 {
					log.Fatalf("Invalid enp rank")
				}
			} else {
				brd.EnP = from - board.Square(10)
				if board.RankArr[brd.EnP] != board.RANK_6 {
					log.Fatalf("Invalid enp rank")
				}
			}
			hashEnP(brd)
		}
	}

	movePiece(from, to, brd)

	if pro != board.EMPTY {
		if pro == board.Piece(board.OFFBOARD) || board.PiecePawn[pro] {
			log.Fatalf("Invalid promotion")
		}
		clearPiece(to, brd)
		addPiece(to, brd, pro)
	}

	brd.Side ^= 1
	hashSide(brd)

	if board.PieceKing[brd.Pieces[to]] {
		brd.KingSquare[side] = int(to)
	}

	// board.CheckBoard(brd)
	if SquareAttacked(board.Square(brd.KingSquare[side]), brd.Side, brd) {
		TakeMove(brd)
		return false
	}

	return true
}

func TakeMove(brd *board.S_Board) {
	// board.CheckBoard(brd)
	var mv S_Move
	brd.HisPly--
	brd.Ply--

	mv.Move = brd.History[brd.HisPly].Move

	from, to, capt, pro, flag := getMove(&mv)

	brd.Side ^= 1
	hashSide(brd)
	side := brd.Side

	if flag == FLAGENP {
		if side == board.WHITE {
			addPiece(to-board.Square(10), brd, board.Bp)
		} else {
			addPiece(to+board.Square(10), brd, board.Wp)
		}
	}

	if flag == FLAGC {
		switch to {
		case board.C1:
			movePiece(board.D1, board.A1, brd)
		case board.C8:
			movePiece(board.D8, board.A8, brd)
		case board.G1:
			movePiece(board.F1, board.H1, brd)
		case board.G8:
			movePiece(board.F8, board.H8, brd)
		default:
			log.Fatalf("Invalid to square in castling (%v)", to)
		}
	}

	// Remove existing castling perms and re-calculate them
	hashC(brd)
	if brd.EnP != board.NO_SQ {
		hashEnP(brd)
	}
	brd.CastlePerm = brd.History[brd.HisPly].CastlePerm
	brd.EnP = brd.History[brd.HisPly].EnP
	brd.FiftyMove = brd.History[brd.HisPly].FiftyMove

	hashC(brd)
	if brd.EnP != board.NO_SQ {
		hashEnP(brd)
	}

	movePiece(to, from, brd)

	if capt != board.EMPTY {
		if capt == board.Piece(board.OFFBOARD) {
			log.Fatalf("MakeMove: Invalid capture piece")
		}
		addPiece(to, brd, capt)
	}

	if pro != board.EMPTY {
		if pro == board.Piece(board.OFFBOARD) || board.PiecePawn[pro] {
			log.Fatalf("Invalid promotion")
		}
		clearPiece(from, brd)
		if side == board.WHITE {
			addPiece(from, brd, board.Wp)
		} else {
			addPiece(from, brd, board.Bp)
		}
	}
	if board.PieceKing[brd.Pieces[from]] {
		brd.KingSquare[side] = int(from)
	}
	//board.CheckBoard(brd)
}

func hashEnP(brd *board.S_Board) {
	brd.PosKey ^= board.PieceKeys[board.EMPTY][brd.EnP]
}

func hashPiece(brd *board.S_Board, piece board.Piece, square board.Square) {
	brd.PosKey ^= board.PieceKeys[piece][square]
}

func hashC(brd *board.S_Board) {
	if brd.CastlePerm != 0 {
		brd.PosKey ^= board.CastleKeys[brd.CastlePerm]
	}
}

func hashSide(brd *board.S_Board) {
	brd.PosKey ^= board.SideKey
}
