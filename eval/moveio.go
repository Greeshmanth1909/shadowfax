package eval

import (
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
	"log"
	"strconv"
)

func PrintMove(mv *S_Move) {
	frmSq := GetFromSquare(mv)
	toSq := GetToSquare(mv)

	if GetPromotedPiece(mv) == board.Piece(0) {
		fmt.Printf("%v%v\n", ConvSq120ToAlge(frmSq), ConvSq120ToAlge(toSq))
	} else {
		fmt.Printf("%v%v%v\n", ConvSq120ToAlge(frmSq), ConvSq120ToAlge(toSq), GetPromotedPieceAlg(GetPromotedPiece(mv)))
	}
}

// ConvSq120ToAlge function converts sq120 to corresponding algebraic notation i.e. 87 -> A7
func ConvSq120ToAlge(sq board.Square) string {
	sqInt := int(sq)
	file := sqInt % 10
	rank := (sqInt / 10) - 1
	var algFile string
	switch file {
	case 1:
		algFile = "a"
	case 2:
		algFile = "b"
	case 3:
		algFile = "c"
	case 4:
		algFile = "d"
	case 5:
		algFile = "e"
	case 6:
		algFile = "f"
	case 7:
		algFile = "g"
	case 8:
		algFile = "h"
	default:
		log.Fatalf("something went wrong (%v)", file)
	}

	strRnk := strconv.Itoa(rank)
	return algFile + strRnk
}

// GetPromotedPieceAlg function returns algebraic notation of a given prompted, returns q (queen) by default
func GetPromotedPieceAlg(p board.Piece) string {
	switch p {
	case board.Wk, board.Bk:
		return "n"
	case board.Wb, board.Bb:
		return "b"
	case board.Wq, board.Bq:
		return "q"
	case board.Wr, board.Br:
		return "r"
	default:
		return "q"
	}
}

// PrintMoveList function prints all the moves in a move list to the screen
func PrintMoveList(list *S_MoveList) {
	for i, val := range list.MoveList {
		if val.Move == 0 {
			break
		}
		fmt.Printf("%v. Move: ", i+1)
		PrintMove(&val)
//		fmt.Println("Score: ", val.Score)
	}
	fmt.Println("Total move list: ", list.Count)
}
