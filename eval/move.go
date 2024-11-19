package eval

import (
	"github.com/Greeshmanth1909/shadowfax/board"
)

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

    if f1 == uint32(1) << 18 {
        // set enp
        SetEnP(&mv)
    }
    if f1 == uint32(1) << 24 {
        SetCastleFlag(&mv)
    }
    if f1 == uint32(1) << 19 {
        SetPawnStart(&mv)
    }
    return mv.Move
}
