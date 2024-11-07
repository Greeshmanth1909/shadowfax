package eval

import (
	"github.com/Greeshmanth1909/shadowfax/board"
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

	// Night
	setPiece(&attackerPiece, side, board.Wn, board.Bn)
	nightAttackSquares := [8]int{sq + 12, sq + 8, sq + 21, sq + 19, sq - 12, sq - 8, sq - 21, sq - 19}
	for _, val := range nightAttackSquares {
		if brd.Pieces[val] == attackerPiece {
			return true
		}
	}

	// Rook
	setPiece(&attackerPiece, side, board.Wr, board.Br) // pieces of same piece-type but opposite color MUST be used here
	rookAlongFile := checkFile(sq, attackerPiece, brd)
	rookAlongRank := checkRank(sq, attackerPiece, brd)
	if rookAlongFile || rookAlongRank {
		return true
	}

	// Bishop
	setPiece(&attackerPiece, side, board.Wb, board.Bb)
	bishopAlongDiagonal := checkDiagonals(sq, attackerPiece, brd)
	if bishopAlongDiagonal {
		return true
	}

	// Queen
	setPiece(&attackerPiece, side, board.Wq, board.Bq)
	queenAlongFile := checkFile(sq, attackerPiece, brd)
	queenAlongRank := checkRank(sq, attackerPiece, brd)
	queenAlongDiagonals := checkDiagonals(sq, attackerPiece, brd)
	if queenAlongFile || queenAlongRank || queenAlongDiagonals {
		return true
	}

	// King
	setPiece(&attackerPiece, side, board.Wk, board.Bk)
	kingAttackSquares := [8]int{sq - 10, sq - 11, sq - 9, sq + 1, sq - 1, sq + 10, sq + 9, sq + 11}
	for _, square := range kingAttackSquares {
		switch brd.Pieces[square] {
		case board.Piece(board.OFFBOARD):
		case board.EMPTY:
		case piece:
			return true
		}
	}

	// All pieces covered, return false
	return false
}

// setPiece is a helper function that sets piece stored in piece pointer to either white or black version of that piece based on color
func setPiece(piece *board.Piece, color board.Color, whitePiece, blackPiece board.Piece) {
	if color == board.WHITE {
		*piece = whitePiece
	} else {
		*piece = blackPiece
	}
}

// checkFile function takes a square and a piece and checks if that piece attacks the square along the square's file
func checkFile(sq int, piece board.Piece, brd *board.S_Board) bool {
	colSquare := sq + 10
	for {
		// ignore empty suares
		if brd.Pieces[colSquare] == board.EMPTY {
			colSquare += 10
			continue
		}
		if brd.Pieces[colSquare] == board.Piece(board.OFFBOARD) {
			break
		}
		if brd.Pieces[colSquare] != piece {
			// Another piece is 'blocking' the file, no point in searching further
			break
		}
		if brd.Pieces[colSquare] == piece {
			return true
		}
	}
	colSquare = sq - 10
	for {
		// ignore empty suares
		if brd.Pieces[colSquare] == board.EMPTY {
			colSquare -= 10
			continue
		}
		if brd.Pieces[colSquare] == board.Piece(board.OFFBOARD) {
			break
		}
		if brd.Pieces[colSquare] != piece {
			// Another piece is 'blocking' the file, no point in searching further
			break
		}
		if brd.Pieces[colSquare] == piece {
			return true
		}
	}
	return false
}

// checkRank function checks weather square `sq` is attacked by the piece `piece` along the square's rank
func checkRank(sq int, piece board.Piece, brd *board.S_Board) bool {
	rowSquare := sq + 1
	for {
		if brd.Pieces[rowSquare] == board.Piece(board.OFFBOARD) {
			break
		}
		if brd.Pieces[rowSquare] == board.EMPTY {
			rowSquare += 1
			continue
		}
		if brd.Pieces[rowSquare] == piece {
			return true
		}
		if brd.Pieces[rowSquare] != piece {
			break
		}
	}
	rowSquare = sq - 1
	for {
		if brd.Pieces[rowSquare] == board.Piece(board.OFFBOARD) {
			break
		}
		if brd.Pieces[rowSquare] == board.EMPTY {
			rowSquare -= 1
			continue
		}
		if brd.Pieces[rowSquare] == piece {
			return true
		}
		if brd.Pieces[rowSquare] != piece {
			break
		}
	}
	return false
}

// checkDiagonals function checks all squares along possible diagonals for the attacker piece that isn't blocked by another piece
func checkDiagonals(sq int, piece board.Piece, brd *board.S_Board) bool {
	// upper left
	dLeft := sq - 11
	for {
		if brd.Pieces[dLeft] == board.Piece(board.OFFBOARD) {
			break
		}
		if brd.Pieces[dLeft] == board.EMPTY {
			dLeft -= 11
			continue
		}
		if brd.Pieces[dLeft] == piece {
			return true
		}
		if brd.Pieces[dLeft] != piece {
			break
		}
	}
	dLeft = sq + 11
	for {
		if brd.Pieces[dLeft] == board.Piece(board.OFFBOARD) {
			break
		}
		if brd.Pieces[dLeft] == board.EMPTY {
			dLeft += 11
			continue
		}
		if brd.Pieces[dLeft] == piece {
			return true
		}
		if brd.Pieces[dLeft] != piece {
			break
		}
	}
	dLeft = sq + 9
	for {
		if brd.Pieces[dLeft] == board.Piece(board.OFFBOARD) {
			break
		}
		if brd.Pieces[dLeft] == board.EMPTY {
			dLeft += 9
			continue
		}
		if brd.Pieces[dLeft] == piece {
			return true
		}
		if brd.Pieces[dLeft] != piece {
			break
		}
	}
	dLeft = sq - 9
	for {
		if brd.Pieces[dLeft] == board.Piece(board.OFFBOARD) {
			break
		}
		if brd.Pieces[dLeft] == board.EMPTY {
			dLeft -= 9
			continue
		}
		if brd.Pieces[dLeft] == piece {
			return true
		}
		if brd.Pieces[dLeft] != piece {
			break
		}
	}

	return false
}
