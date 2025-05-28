package search

import (
	"log"
	"os"
	"slices"

	"github.com/Greeshmanth1909/shadowfax/board"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

const isolatedPawnEval int = 10

var KnightTable = [64]int{
	0, -10, 0, 0, 0, 0, -10, 0,
	0, 0, 0, 5, 5, 0, 0, 0,
	0, 0, 10, 10, 10, 10, 0, 0,
	0, 0, 10, 20, 20, 10, 5, 0,
	5, 10, 15, 20, 20, 15, 10, 5,
	5, 10, 10, 20, 20, 10, 10, 5,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var PawnTable = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	10, 10, 0, -10, -10, 0, 10, 10,
	5, 0, 0, 5, 5, 0, 0, 5,
	0, 0, 10, 20, 20, 10, 0, 0,
	5, 5, 5, 10, 10, 5, 5, 5,
	10, 10, 10, 20, 20, 10, 10, 10,
	20, 20, 20, 30, 30, 20, 20, 20,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var BishopTable = [64]int{
	0, 0, -10, 0, 0, -10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var RookTable = [64]int{
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	25, 25, 25, 25, 25, 25, 25, 25,
	0, 0, 5, 10, 10, 5, 0, 0,
}

var KingE = [64]int{
	-50, -10, 0, 0, 0, 0, -10, -50,
	-10, 0, 10, 10, 10, 10, 0, -10,
	0, 10, 20, 20, 20, 20, 10, 0,
	0, 10, 20, 40, 40, 20, 10, 0,
	0, 10, 20, 40, 40, 20, 10, 0,
	0, 10, 20, 20, 20, 20, 10, 0,
	-10, 0, 10, 10, 10, 10, 0, -10,
	-50, -10, 0, 0, 0, 0, -10, -50,
}

var KingO = [64]int{
	0, 5, 5, -10, -10, 0, 10, 5,
	-30, -30, -30, -30, -30, -30, -30, -30,
	-50, -50, -50, -50, -50, -50, -50, -50,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
}

var Mirror64 = [64]int{
	56, 57, 58, 59, 60, 61, 62, 63,
	48, 49, 50, 51, 52, 53, 54, 55,
	40, 41, 42, 43, 44, 45, 46, 47,
	32, 33, 34, 35, 36, 37, 38, 39,
	24, 25, 26, 27, 28, 29, 30, 31,
	16, 17, 18, 19, 20, 21, 22, 23,
	8, 9, 10, 11, 12, 13, 14, 15,
	0, 1, 2, 3, 4, 5, 6, 7,
}

func EvalPosition(brd *board.S_Board) (score int) {
	var piece board.Piece
	var square board.Square
	score = brd.Material[board.WHITE] - brd.Material[board.BLACK]

	piece = board.Wp
	for pieceNum := 0; pieceNum < brd.PieceNum[piece]; pieceNum++ {
		square = board.Square(brd.PList[piece][pieceNum])
		if square == board.OFFBOARD || square == board.Square(board.EMPTY) {
			log.Fatalf("Invalid square %v", square)
		}
		score += PawnTable[board.Square120to64[square]]
	}

	piece = board.Bp
	for pieceNum := 0; pieceNum < brd.PieceNum[piece]; pieceNum++ {
		square = board.Square(brd.PList[piece][pieceNum])
		if square == board.OFFBOARD || square == board.Square(board.EMPTY) {
			log.Fatalf("Invalid square")
		}
		score -= PawnTable[Mirror64[board.Square120to64[square]]]
	}

	piece = board.Wn
	for pieceNum := 0; pieceNum < brd.PieceNum[piece]; pieceNum++ {
		square = board.Square(brd.PList[piece][pieceNum])
		if square == board.OFFBOARD || square == board.Square(board.EMPTY) {
			log.Fatalf("Invalid square")
		}
		score += KnightTable[board.Square120to64[square]]
	}

	piece = board.Bn
	for pieceNum := 0; pieceNum < brd.PieceNum[piece]; pieceNum++ {
		square = board.Square(brd.PList[piece][pieceNum])
		if square == board.OFFBOARD || square == board.Square(board.EMPTY) {
			log.Fatalf("Invalid square")
		}
		score -= KnightTable[Mirror64[board.Square120to64[square]]]
	}

	piece = board.Wb
	for pieceNum := 0; pieceNum < brd.PieceNum[piece]; pieceNum++ {
		square = board.Square(brd.PList[piece][pieceNum])
		if square == board.OFFBOARD || square == board.Square(board.EMPTY) {
			log.Fatalf("Invalid square")
		}
		score += BishopTable[board.Square120to64[square]]
	}

	piece = board.Bb
	for pieceNum := 0; pieceNum < brd.PieceNum[piece]; pieceNum++ {
		square = board.Square(brd.PList[piece][pieceNum])
		if square == board.OFFBOARD || square == board.Square(board.EMPTY) {
			log.Fatalf("Invalid square")
		}
		score -= BishopTable[Mirror64[board.Square120to64[square]]]
	}

	// piece = board.Wq
	// for pieceNum := 0; pieceNum < brd.PieceNum[piece]; pieceNum++ {
	//     square = brd.Pieces[pieceNum]
	//     if square == board.OFFBOARD || square == board.Square(board.EMPTY) {
	//         log.Fatalf("Invalid square")
	//     }
	//     score += QueenTable[board.Square120to64[square]]
	// }
	//
	// piece = board.Bq
	// for pieceNum := 0; pieceNum < brd.PieceNum[piece]; pieceNum++ {
	//     square = brd.Pieces[pieceNum]
	//     if square == board.OFFBOARD || square == board.Square(board.EMPTY) {
	//         log.Fatalf("Invalid square")
	//     }
	//     score -= QueenTable[Mirror64[board.Square120to64[square]]]
	// }

	piece = board.Wr
	for pieceNum := 0; pieceNum < brd.PieceNum[piece]; pieceNum++ {
		square = board.Square(brd.PList[piece][pieceNum])
		if square == board.OFFBOARD || square == board.Square(board.EMPTY) {
			log.Fatalf("Invalid square")
		}
		score += RookTable[board.Square120to64[square]]
	}

	piece = board.Br
	for pieceNum := 0; pieceNum < brd.PieceNum[piece]; pieceNum++ {
		square = board.Square(brd.PList[piece][pieceNum])
		if square == board.OFFBOARD || square == board.Square(board.EMPTY) {
			log.Fatalf("Invalid square")
		}
		score -= RookTable[Mirror64[board.Square120to64[square]]]
	}

	// // account for isolated Pawns
	// whiteIsolatedPawn, isolatedPawnNumber := isolatedPawn(brd, board.WHITE)
	// if whiteIsolatedPawn {
	//     score -= isolatedPawnNumber * isolatedPawnEval
	// }
	// blackIsolatedPawn, isolatedPawnNumber := isolatedPawn(brd, board.WHITE)
	// if blackIsolatedPawn {
	//     score += isolatedPawnNumber * isolatedPawnNumber
	// }

	if brd.Side == board.WHITE {
		return score
	} else {
		return -score
	}
}

/*
Returns a boolean stating the existence of an isolated pawn for side and
number of isolated paws.
*/
func isolatedPawn(brd *board.S_Board, side board.Color) (exists bool, number int) {
	pawnList := [2]board.Piece{board.Wp, board.Bp}
	pawn := pawnList[side]
	pawnsArray := brd.PList[pawn][:brd.PieceNum[pawn]]
	filesArray := []board.File{}
	for i := range pawnsArray {
		filesArray = append(filesArray, board.File(pawnsArray[i]/10))
	}
	for _, file := range filesArray {
		switch file {
		case board.FILE_A:
			if !slices.Contains(filesArray, board.FILE_B) {
				number++
			}
		case board.FILE_B:
			if !slices.Contains(filesArray, board.FILE_A) && !slices.Contains(filesArray, board.FILE_C) {
				number++
			}
		case board.FILE_C:
			if !slices.Contains(filesArray, board.FILE_B) && !slices.Contains(filesArray, board.FILE_D) {
				number++
			}
		case board.FILE_D:
			if !slices.Contains(filesArray, board.FILE_C) && !slices.Contains(filesArray, board.FILE_E) {
				number++
			}
		case board.FILE_E:
			if !slices.Contains(filesArray, board.FILE_D) && !slices.Contains(filesArray, board.FILE_F) {
				number++
			}
		case board.FILE_F:
			if !slices.Contains(filesArray, board.FILE_E) && !slices.Contains(filesArray, board.FILE_G) {
				number++
			}
		case board.FILE_G:
			if !slices.Contains(filesArray, board.FILE_F) && !slices.Contains(filesArray, board.FILE_H) {
				number++
			}
		case board.FILE_H:
			if !slices.Contains(filesArray, board.FILE_G) {
				number++
			}
		}
	}
	if number != 0 {
		exists = true
	}
	return
}
