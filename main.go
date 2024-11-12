package main

import (
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
	_ "github.com/Greeshmanth1909/shadowfax/eval"
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
	for i := range boardStructure.Pieces {
		boardStructure.Pieces[i] = board.Piece(board.OFFBOARD)
	}
	position.Parse_FEN(&startString, &boardStructure)

	fmt.Println("******")
	fmt.Println(boardStructure.Pieces[board.Square64to120[4]])
	fmt.Println(boardStructure.Pieces[board.Square64to120[4]] == board.Bk)
	fmt.Println(boardStructure.Pieces[board.Square64to120[28]] == board.EMPTY)

	position.PrintBoard(&boardStructure)
	fmt.Println("******")
	startString = "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)
	fmt.Println("******")
	startString = "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)
	startString = "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2"
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)

	startString = "rnbq1rk1/ppp2ppp/3bpn2/3p4/3P4/2P2N2/PP2PPPP/RNBQKB1R w KQ - 2 7"
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)

	//board.CheckBoard(&boardStructure)

	startString = "8/8/8/4P3/8/2B5/8/8 w - - 0 1"
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)
}
