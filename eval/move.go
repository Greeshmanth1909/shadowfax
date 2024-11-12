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
