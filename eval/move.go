package eval

import (
	"github.com/Greeshmanth1909/shadowfax/board"
	"strconv"
)

const NOMOVE uint32 = 0

type S_Move struct {
	Move  uint32
	Score int
}

// setFromSquare function sets the given square to the corresponding bits in the move int
func SetFromSquare(mv *S_Move, square board.Square) {
	mv.Move |= uint32(square)
}

// setToSquare
func SetToSquare(mv *S_Move, square board.Square) {
	mv.Move |= (uint32(square) << 7)
}

// setCapturedPiece
func SetCapturedPiece(mv *S_Move, piece board.Piece) {
	mv.Move |= (uint32(piece) << 14)
}

// setEnP sets the enpassant flag
func SetEnP(mv *S_Move) {
	mv.Move |= uint32(1) << 18
}

// setPawnStart sets the pawn start move flag
func SetPawnStart(mv *S_Move) {
	mv.Move |= uint32(1) << 19
}

func SetPromotedPiece(mv *S_Move, piece board.Piece) {
	mv.Move |= (uint32(piece) << 20)
}

func SetCastleFlag(mv *S_Move) {
	mv.Move |= uint32(1) << 24
}

// returns from square
func GetFromSquare(mv *S_Move) board.Square {
	return board.Square(mv.Move & 0x7f)
}

// returns to square
func GetToSquare(mv *S_Move) board.Square {
	return board.Square((mv.Move >> 7) & 0x7f)
}

// returns captured piece
func GetCapturedPiece(mv *S_Move) board.Piece {
	return board.Piece((mv.Move >> 14) & 0xf)
}

// returns true if enp flag is set
func GetEnP(mv *S_Move) bool {
	if mv.Move&0x40000 != 0 {
		return true
	}
	return false
}

// returns true if pawnStart flag is set
func GetPawnStart(mv *S_Move) bool {
	if mv.Move&0x80000 != 0 {
		return true
	}
	return false
}

// returns promoted piece, if any
func GetPromotedPiece(mv *S_Move) board.Piece {
	return board.Piece((mv.Move >> 20) & 0xf)
}

// returns true if castle flag is set
func GetCastleFlag(mv *S_Move) bool {
	if mv.Move&0x1000000 != 0 {
		return true
	}
	return false
}

// Move function uses the above helpers to generate a move integer
func Move(frm, to board.Square, capt, pro board.Piece, f1 uint32) uint32 {
	mv := S_Move{}
	SetFromSquare(&mv, frm)
	SetToSquare(&mv, to)
	SetCapturedPiece(&mv, capt)
	SetPromotedPiece(&mv, pro)

	if f1 == FLAGENP {
		SetEnP(&mv)
	} else if f1 == FLAGC {
		SetCastleFlag(&mv)
	} else if f1 == FLAGPS {
		SetPawnStart(&mv)
	}
	return mv.Move
}

func getMove(mv *S_Move) (frm, to board.Square, capt, pro board.Piece, flag uint32) {
	frm = GetFromSquare(mv)
	to = GetToSquare(mv)
	capt = GetCapturedPiece(mv)
	pro = GetPromotedPiece(mv)
	flag = uint32(0)
	if GetEnP(mv) {
		flag = FLAGENP
	} else if GetPawnStart(mv) {
		flag = FLAGPS
	} else if GetCastleFlag(mv) {
		flag = FLAGC
	}
	return frm, to, capt, pro, flag
}

func ParseMove(move string, brd *board.S_Board) (bool, uint32) {
	if len(move) > 5 {
		return false, NOMOVE
	}
	if !(move[0] >= 'a' && move[0] <= 'h') {
		return false, NOMOVE
	}
	if !(move[1] >= '1' && move[1] <= '8') {
		return false, NOMOVE
	}
	if !(move[2] >= 'a' && move[2] <= 'h') {
		return false, NOMOVE
	}
	if !(move[3] >= '1' && move[3] <= '8') {
		return false, NOMOVE
	}

	side := brd.Side

	from := convertSquareStringToSquare(move[:2])
	to := convertSquareStringToSquare(move[2:4])

	var mvList S_MoveList
	GenerateAllMoves(brd, &mvList)
	if mvList.MoveList[0].Move == 0 {
		return true, NOMOVE
	}
	for _, val := range mvList.MoveList {
		legalFrom := GetFromSquare(&val)
		legalTo := GetToSquare(&val)
		if legalFrom == from && legalTo == to {
			if len(move) == 5 {
				pro := checkPromPiece(move, side)
				legalPromPiece := GetPromotedPiece(&val)
				if pro == legalPromPiece {
					return false, val.Move
				}
			}
			return false, val.Move
		}
	}
	return false, NOMOVE
}

func convertSquareStringToSquare(square string) board.Square {
	file := square[0]
	rank, _ := strconv.Atoi(string(square[1]))
	r := 0
	switch file {
	case 'a':
		r = 1
	case 'b':
		r = 2
	case 'c':
		r = 3
	case 'd':
		r = 4
	case 'e':
		r = 5
	case 'f':
		r = 6
	case 'g':
		r = 7
	case 'h':
		r = 8
	}
	return board.Square((rank+1)*10 + r)
}

func checkPromPiece(move string, side board.Color) board.Piece {
	str := rune(move[4])
	pieces := map[rune][2]board.Piece{
		'q': {board.Wq, board.Bq},
		'r': {board.Wr, board.Br},
		'b': {board.Wb, board.Bb},
		'n': {board.Wn, board.Bn},
	}
	ls, ok := pieces[str]
	if !ok {
		return pieces['q'][side]
	}
	return ls[side]
}

// adds valid moves from pv table to pv array
func GetPvLine(depth int, brd *board.S_Board) int {
	move := board.ProbePvTable(brd, brd.PosKey)
	count := 0
	for move != 0 && count < depth {
		var m S_Move
		m.Move = move
		if MakeMove(brd, &m) {
			brd.PvArray[count] = move
			count++
		} else {
			break
		}
		move = board.ProbePvTable(brd, brd.PosKey)
	}
	for brd.Ply > 0 {
		TakeMove(brd)
	}
	return count
}
