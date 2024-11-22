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

	startString = "rnbqkbnr/p1p1p3/3p3p/1p1p4/2P1Pp2/8/PP1P1PpP/RNBQKB1R w KQkq e3 0 1"
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)

	startString = "5k2/1n6/4n3/2p3N1/8/3N4/8/5K2 b - - 0 1"
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)

	var list eval.S_MoveList
	eval.GenerateAllMoves(&boardStructure, &list)
	eval.PrintMoveList(&list)
	// fmt.Println(boardStructure.EnP)
	// fmt.Println(eval.ConvSq120ToAlge(boardStructure.EnP))
}
