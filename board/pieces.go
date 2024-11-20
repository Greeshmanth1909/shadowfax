package board

import (
	"log"
	"os"
	"reflect"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

// empty, pawn, bishop, night, rook, queen, king, pawn, bishop, night, rook, queen, king
var BigPiece = [13]bool{false, false, true, true, true, true, true, false, true, true, true, true, true}
var MajPiece = [13]bool{false, false, false, false, true, true, true, false, false, false, true, true, true}
var MinPiece = [13]bool{false, false, true, true, false, false, false, false, true, true, false, false, false}
var PieceVal = [13]int{0, 100, 325, 325, 550, 1000, 50000, 100, 325, 325, 550, 1000, 50000}
var PieceCol = [13]Color{BOTH, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, BLACK, BLACK, BLACK, BLACK, BLACK, BLACK}
var PieceSlides = [13]bool{false, false, false, true, true, true, false, false, false, true, true, true, false}

// UpdatePieceList iterates over the entire board and adds updates the existing pieces to the board struct accordingly
func UpdatePieceList(brd *S_Board) {
	// set the piece list values to zero to avoid recounting everything
	for i := 0; i < 2; i++ {
		brd.BigPiece[i] = 0
		brd.MinPiece[i] = 0
		brd.MajPiece[i] = 0
		brd.Pawns[i] = uint64(0)
		brd.Material[i] = 0
	}
	// brd.Pawns has three entries; white, black and both
	brd.Pawns[2] = uint64(0)

	for i := 0; i < 13; i++ {
		brd.PieceNum[i] = 0
		for j := 0; j < 10; j++ {
			brd.PList[i][j] = 0
		}
	}

	for sq, val := range brd.Pieces {
		if int(val) == 100 {
			continue
		}
		if val != EMPTY {
			piece := val
			color := PieceCol[piece]
			if BigPiece[piece] {
				brd.BigPiece[color]++
			}
			if MajPiece[piece] {
				brd.MajPiece[color]++
			}
			if MinPiece[piece] {
				brd.MinPiece[color]++
			}
			brd.Material[color] += PieceVal[piece]

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

func CheckBoard(brd *S_Board) {
	var T_PieceNum [13]int
	var T_BigPiece, T_MajPiece, T_Material, T_MinPiece [2]int

	for piece := Wp; piece <= Bk; piece++ {
		for pieceIndex := 0; pieceIndex < brd.PieceNum[piece]; pieceIndex++ {
			sq := brd.PList[piece][pieceIndex]
			// check weather the piece exists on the board
			if brd.Pieces[sq] != piece {
				log.Fatalf("CheckBoard: Invalid Piece Placement")
			}
		}
	}

	// Material count, piece count etc
	for _, sq120 := range Square64to120 {
		piece := brd.Pieces[sq120]
		if piece != EMPTY {
			color := PieceCol[piece]
			if BigPiece[piece] {
				T_BigPiece[color]++
			}
			if MajPiece[piece] {
				T_MajPiece[color]++
			}
			if MinPiece[piece] {
				T_MinPiece[color]++
			}

			T_Material[color] += PieceVal[piece]
			T_PieceNum[piece]++
		}
	}

	// Ensure PieceNums are the same
	if !reflect.DeepEqual(T_PieceNum, brd.PieceNum) {
		log.Fatalf("CheckBoard: PieceNums not matching")
	}

	// Count pawn bitboards
	wpCount := CountBits(brd.Pawns[WHITE])
	bpCount := CountBits(brd.Pawns[BLACK])
	if wpCount != brd.PieceNum[Wp] {
		log.Fatalf("CheckBoard: White pawn count not matching")
	}
	if bpCount != brd.PieceNum[Bp] {
		log.Fatalf("CheckBoard: Black pawn count not matching")
	}

	// Check pawn bitboards
	t_pawns := [3]uint64{brd.Pawns[WHITE], brd.Pawns[BLACK], brd.Pawns[BOTH]}
	for {
		sq64 := PopBits(&t_pawns[WHITE])
		if sq64 == 64 {
			break
		}
		if brd.Pieces[Square64to120[sq64]] != Wp {
			log.Fatalf("CheckBoard: White pawn on the wrong square")
		}
	}
	for {
		sq64 := PopBits(&t_pawns[BLACK])
		if sq64 == 64 {
			break
		}
		if brd.Pieces[Square64to120[sq64]] != Bp {
			log.Fatalf("CheckBoard: Black pawn on the wrong square")
		}
	}
	for {
		sq64 := PopBits(&t_pawns[BOTH])
		if sq64 == 64 {
			break
		}
		if !(brd.Pieces[Square64to120[sq64]] == Wp || brd.Pieces[Square64to120[sq64]] == Bp) {
			log.Fatalf("CheckBoard: pawn on the wrong square")
		}
	}

	// Other assertions
	if !reflect.DeepEqual(T_Material, brd.Material) {
		log.Fatalf("CheckBoard: Material not equal temp != brd (%v != %v)", T_Material, brd.Material)
	}
	if !reflect.DeepEqual(T_BigPiece, brd.BigPiece) {
		log.Fatalf("CheckBoard: BigPiece not equal temp != brd (%v != %v)", T_BigPiece, brd.BigPiece)
	}
	if !reflect.DeepEqual(T_MinPiece, brd.MinPiece) {
		log.Fatalf("CheckBoard: MinPiece not equal temp != brd (%v != %v)", T_MinPiece, brd.MinPiece)
	}
	if !reflect.DeepEqual(T_MajPiece, brd.MajPiece) {
		log.Fatalf("CheckBoard: MajPiece not equal temp != brd (%v != %v)", T_MajPiece, brd.MajPiece)
	}

	if !(brd.Side == WHITE || brd.Side == BLACK) {
		log.Fatalf("CheckBoard: Invalid Side %v", brd.Side)
	}

	if GenerateHash(brd) != brd.PosKey {
		log.Fatalf("Checkboard: hash generation mismatch")
	}

	// King position
	if brd.Pieces[brd.KingSquare[WHITE]] != Wk {
		log.Fatalf("CheckBoard: Misplaced White King %v", brd.Pieces[brd.KingSquare[WHITE]])
	}
	if brd.Pieces[brd.KingSquare[BLACK]] != Bk {
		log.Fatalf("CheckBoard: Misplaced Black King %v", brd.Pieces[brd.KingSquare[BLACK]])
	}

	// TODO: Add enp square checking
}
