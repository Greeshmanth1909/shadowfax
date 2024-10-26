package main

import (
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/Greeshmanth1909/shadowfax/position"
	"github.com/Greeshmanth1909/shadowfax/util"
)

func main() {
	util.InitAll()
	fmt.Printf("%v\nStatus: running\n", board.Name)
	fmt.Println("Printing bitboard")
	var bb uint64 = 0
	board.PrintBitBoard(bb)
	board.SetBit(63, &bb)
	board.PrintBitBoard(bb)
	board.SetBit(15, &bb)
	board.PrintBitBoard(bb)
	fmt.Println("********")
	board.ClearBit(15, &bb)
	board.PrintBitBoard(bb)

	startString := position.StartPosition
	var boardStructure board.S_Board
	position.Parse_FEN(&startString, &boardStructure)

	fmt.Println("******")
	fmt.Println(boardStructure.Pieces[board.Square64to120[4]])
	fmt.Println(boardStructure.Pieces[board.Square64to120[4]] == board.Bk)
	fmt.Println(boardStructure.Pieces[board.Square64to120[28]] == board.EMPTY)

	position.PrintBoard(&boardStructure)
}
