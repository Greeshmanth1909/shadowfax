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

	fmt.Println("*******")
	fmt.Println("Add pawn to d2")
	bb = bb | uint64(1)<<board.Square120to64[board.D2]
	board.PrintBitBoard(bb)
	fmt.Println("Add pawn to g2")
	bb = bb | uint64(1)<<board.Square120to64[board.G2]
	board.PrintBitBoard(bb)
	fmt.Println("Add pawn to b4")
	bb |= uint64(1) << board.Square120to64[board.B4]
	board.PrintBitBoard(bb)
	bb |= uint64(1) << board.Square120to64[board.H8]
	board.PrintBitBoard(bb)
	fmt.Printf("bitcount %v\n", board.CountBits(bb))
	fmt.Println("Popping")
	ind := board.PopBits(&bb)
	board.PrintBitBoard(bb)
	fmt.Println(ind)
}
