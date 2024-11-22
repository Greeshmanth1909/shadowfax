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

	startString = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)

	var list eval.S_MoveList
	eval.GenerateAllMoves(&boardStructure, &list)
	eval.PrintMoveList(&list)
	// fmt.Println(boardStructure.EnP)
}
