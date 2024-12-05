package board

import (
	"math/rand"
)

var PieceKeys [13][120]uint64

var SideKey uint64

var CastleKeys [16]uint64

// InitHashKey initialises position keys with pseudo-random uint64 numbers
func InitHashKeys() {
	for i := 0; i < 13; i++ {
		for j := 0; j < 120; j++ {
			PieceKeys[i][j] = rand.Uint64()
		}
	}
	SideKey = rand.Uint64()
	for i := 0; i < 16; i++ {
		CastleKeys[i] = rand.Uint64()
	}
}

// GenerateHash generates a Zobrist Hash for a given position
func GenerateHash(pos *S_Board) uint64 {
	var finalKey uint64 = 0
	for i, piece := range pos.Pieces {
		if piece != Piece(NO_SQ) && piece != Piece(EMPTY) && piece != Piece(OFFBOARD) {
			finalKey ^= PieceKeys[piece][i]
		}
	}
	// TODO: Add logging and telemetry
	if pos.Side == WHITE {
		finalKey ^= SideKey
	}

	if pos.EnP != NO_SQ {
		finalKey ^= PieceKeys[EMPTY][pos.EnP]
	}

	if pos.CastlePerm != 0 {
		finalKey ^= CastleKeys[pos.CastlePerm]
	}

	return finalKey
}
