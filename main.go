package main

import (
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
)

func main() {
	s1 := board.InitSquares64()
	s2 := board.InitSquares120()
	fmt.Printf("%v\nStatus: running\n", board.Name)

	fmt.Println(s1)
	//board.PrintBitBoard()
	fmt.Println(s2)
	fmt.Println(board.FRtoSq120(board.FILE_E, board.RANK_4))
	fmt.Println("Printing bitboard")
	var bb uint64 = 0
	board.PrintBitBoard(bb)

	board.InitBitMasks()
	board.SetBit(63, &bb)
	board.PrintBitBoard(bb)
	board.SetBit(15, &bb)
	board.PrintBitBoard(bb)
	fmt.Println("********")
	board.ClearBit(15, &bb)
	board.PrintBitBoard(bb)

}
