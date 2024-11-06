package eval

import (
	"github.com/Greeshmath1909/shadowfax/board"
	"log"
)

// SquareAttacked fuction checks weater a certain square is attacked by a particular side's pieces
func SquareAttacked(sq120 board.Square, side board.Color, brd *board.S_Board) bool {
	if sq120 == board.OFFBOARD {
		log.Fatalf("SquareAttacked: offboard input given %v", sq120)
	}
	sq := int(sq120)

	// check for pawn attacks
	if side == board.WHITE {
		if brd.Pieces[sq-9] == board.Wp || brd.Pieces[sq-11] == board.Wp {
			return true
		}
	}
	if side == board.BLACK {
		if brd.Pieces[sq+9] == board.Bp || brd.Pieces[sq+11] == board.Bp {
			return true
		}
	}

	var attackerPiece board.Piece
	// night
	setPiece(&attackerPiece, side, board.Wn, board.Bn)
	nightAttackSquares := [8]int{sq + 12, sq + 8, sq + 21, sq + 19, sq - 12, sq - 8, sq - 21, sq - 19}
	for _, val := range nightAttackSquares {
		if brd.Pieces[val] == attackerPiece {
			return true
		}
	}

	// Rook
	setPiece(&attackerPiece, side, board.Wr, board.Br) // pieces of same piece-type but opposite color MUST be used here
	columnSquare := sq + 10
	for {
		if brd.Pieces[columnSquare] == board.OFFBOARD {
			break
		}
		if brd.Pieces[columnSquare] == attakerPiece {
			return true
		}
		columnSquare += 10
	}
	columnSquare = sq - 10
	for {
		if brd.Pieces[columnSquare] == board.OFFBOARD {
			break
		}
		if brd.Pieces[columnSquare] == attakerPiece {
			return true
		}
		columnSquare -= 10
	}
}

// setPiece is a helper function that sets piece stored in piece pointer to either white or black version of that piece based on color
func setPiece(piece *board.Piece, color board.Color, whitePiece, blackPiece board.Piece) {
	if color == board.WHITE {
		*piece = whitePiece
	} else {
		*piece = blackPiece
	}
}
