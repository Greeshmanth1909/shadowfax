package board

// empty, pawn, bishop, night, rook, queen, king, pawn, bishop, night, rook, queen, king
var BigPiece = [13]bool{false, false, true, true, true, true, true, false, true, true, true, true, true}
var MajPiece = [13]bool{false, false, false, false, true, true, true, false, false, false, true, true, true}
var MinPiece = [13]bool{false, false, true, true, false, false, false, false, true, true, false, false, false}
var PieceVal = [13]int{0, 100, 325, 325, 550, 1000, 50000, 100, 325, 325, 550, 1000, 50000}
var PieceCol = [13]Color{BOTH, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, BLACK, BLACK, BLACK, BLACK, BLACK, BLACK}

// UpdatePieceList iterates over the entire board and adds updates the existing pieces to the board struct accordingly
func UpdatePieceList(brd *S_Board) {
	// set the piece list values to zero to avoid recounting everything
	for i := 0; i < 2; i++ {
		brd.BigPiece[i] = 0
		brd.MinPiece[i] = 0
		brd.MajPiece[i] = 0
		brd.Pawns[i] = uint64(0)
	}

	for i := 0; i < 13; i++ {
		brd.PieceNum[i] = 0
		for j := 0; j < 10; j++ {
			brd.PList[i][j] = 0
		}
	}

	for sq, val := range brd.Pieces {
		if val != EMPTY {
			piece := val
			color := PieceCol[piece]
			if BigPiece[piece] {
				brd.BigPiece[color]++
			} else if MajPiece[piece] {
				brd.MajPiece[color]++
			} else if MinPiece[piece] {
				brd.MinPiece[color]++
			}

			brd.PList[piece][brd.PieceNum[piece]] = sq
			brd.PieceNum[piece]++

			if piece == Wk {
				brd.KingSquare[WHITE] = sq
			}
			if piece == Bk {
				brd.KingSquare[BLACK] = sq
			}

			if piece == Wp {
				SetBit(Square120to64[sq], &brd.Pawns[WHITE])
				SetBit(Square120to64[sq], &brd.Pawns[BOTH])
			}
			if piece == Bp {
				SetBit(Square120to64[sq], &brd.Pawns[BLACK])
				SetBit(Square120to64[sq], &brd.Pawns[BOTH])
			}
		}
	}
}
