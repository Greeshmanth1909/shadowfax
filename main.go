package main

import (
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
)

func main() {
	InitAll()
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

}

func InitAll() {
	board.InitSquares64()
	board.InitSquares120()
	board.InitBitMasks()
	board.InitHashKeys()
}
