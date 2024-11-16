package main

import (
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/Greeshmanth1909/shadowfax/eval"
	"github.com/Greeshmanth1909/shadowfax/position"
	"github.com/Greeshmanth1909/shadowfax/util"
)

func main() {
	util.InitAll()
	fmt.Printf("%v\nStatus: running\n", board.Name)

	startString := position.StartPosition
	var boardStructure board.S_Board
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)

	startString = "8/8/8/4P3/8/2B5/8/8 w - - 0 1"
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)

	fmt.Println(board.A7)

	var move eval.S_Move
	eval.SetFromSquare(&move, board.A7)
	eval.SetToSquare(&move, board.B8)
	eval.SetCapturedPiece(&move, board.Bk)
	eval.SetPromotedPiece(&move, board.Wr)
	eval.PrintMove(&move)

}
