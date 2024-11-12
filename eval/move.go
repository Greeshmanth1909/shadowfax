package eval

import (
	"github.com/Greeshmanth1909/shadowfax/board"
)

type S_Move struct {
	Move  uint32
	Score int
}

// setFromSquare function sets the given square to the corresponding bits in the move int
func setFromsquare(mv *S_Move, square board.Piece) {
	mv.Move |= uint32(square)
}

// setToSquare
func setToSquare(mv *S_Move, square board.Square) {
	mv.Move |= (uint32(square) << 7)
}

// setCapturedPiece
func setCapturedPiece(mv *S_Move, piece board.Piece) {
	mv.Move |= (uint32(piece) << 14)
}

// setEnP sets the enpassant flag
func setEnP(mv *S_Move) {
	mv.Move |= uint32(1) << 18
}

// setPawnStart sets the pawn start move flag
func setPawnStart(mv *S_Move) {
	mv.Move |= uint32(1) << 19
}

func setPromotedPiece(mv *S_Move, piece board.Piece) {
	mv.Move |= (uint32(piece) << 20)
}

func setCastleFlag(mv *S_Move) {
	mv.Move |= uint32(1) << 24
}

// returns from square
func getFromsquare(mv *S_Move) uint32 {
	return mv.Move & 0x3f
}

// returns to square
func getToSquare(mv *S_Move) uint32 {
	return ((mv.Move >> 7) & 0xf)
}

// returns captured piece
func getCapturedPiece(mv *S_Move) uint32 {
	return (mv.Move >> 14) & 0xf
}

// returns true if enp flag is set
func getEnP(mv *S_Move) bool {
	if mv.Move&0x40000 == 1 {
		return true
	}
	return false
}

// returns true if pawnStart flag is set
func getPawnStart(mv *S_Move) bool {
	if mv.Move&80000 == 1 {
		return true
	}
	return false
}

// returns promoted piece, if any
func getPromotedPiece(mv *S_Move) bool {
	if ((mv.Move >> 20) & 0xf) == 1 {
		return true
	}
	return false
}

// returns true if castle flag is set
func getCastleFlag(mv *S_Move) bool {
	if mv.Move&0x1000000 == 1 {
		return true
	}
	return false
}
