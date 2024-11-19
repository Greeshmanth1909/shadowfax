package eval

import (
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMove(t *testing.T) {
	// set move: pawn capture promption to queen
	var move S_Move
	SetFromSquare(&move, board.A7)
	SetToSquare(&move, board.B8)
	SetCapturedPiece(&move, board.Bk)
	SetPromotedPiece(&move, board.Wq)

    newMv := Move(board.A7, board.B8, board.Bk, board.Wq, 0)
    assert.Equal(t, newMv, move.Move, "error with Move func")
	frmSq := GetFromSquare(&move)
	toSq := GetToSquare(&move)
	capPiece := GetCapturedPiece(&move)
	promPiece := GetPromotedPiece(&move)

	assert.Equal(t, frmSq, board.A7, "from square err")
	assert.Equal(t, toSq, board.B8, "to square err")
	assert.Equal(t, capPiece, board.Bk, "cap Piece err")
	assert.Equal(t, promPiece, board.Wq, "from square err")
}
