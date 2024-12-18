package main

import (
	"bufio"
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/Greeshmanth1909/shadowfax/eval"
	"github.com/Greeshmanth1909/shadowfax/position"
	"github.com/Greeshmanth1909/shadowfax/search"
	"github.com/Greeshmanth1909/shadowfax/util"
	"os"
	// "time"
)

func main() {
	util.InitAll()
	fmt.Printf("%v\nStatus: running\n", board.Name)
	reader := bufio.NewReader(os.Stdin)

	startString := position.StartPosition
	// startString = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"
	startString = "2rr3k/pp3pp1/1nnqbN1p/3pN3/2pP4/2P3Q1/PPB4P/R4RK1 w - - 0 1"
	var boardStructure board.S_Board
	board.InitPvTable(&boardStructure)
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)

	var list eval.S_MoveList
	eval.GenerateAllMoves(&boardStructure, &list)
	// eval.PrintMoveList(&list)

	// fmt.Println(boardStructure.EnP)

	// for i := 0; i < list.Count; i++ {
	// 	mv := list.MoveList[i]
	// 	eval.MakeMove(&boardStructure, &mv)
	// 	position.PrintBoard(&boardStructure)
	// 	text, _ := reader.ReadString('\n')
	// 	fmt.Println(text)
	// 	fmt.Println("Taking move")
	// 	eval.TakeMove(&boardStructure)
	// 	position.PrintBoard(&boardStructure)
	// 	newT, _ := reader.ReadString('\n')
	// 	fmt.Println(newT)
	// }
	var info board.S_SearchInfo

	for {
		position.PrintBoard(&boardStructure)
		val, _ := reader.ReadString('\n')
		isCheckmate, mv := eval.ParseMove(val, &boardStructure)
		if isCheckmate {
			fmt.Println("Checkmate!")
			break
		}
		if val == "quit\n" {
			break
		}
		if val == "t\n" {
			eval.TakeMove(&boardStructure)
		}
		if val == "s\n" {
			// start := time.Now()
			// num := eval.PerftTest(4, &boardStructure)
			// fmt.Println(num)
			// end := time.Since(start).Milliseconds()
			// fmt.Printf("TIME IN MS: %v\n", end)
			info.Depth = 5
			search.SearchPositions(&boardStructure, &info)

		}
		if mv != 0 {
			var m eval.S_Move
			m.Move = mv
			if eval.MakeMove(&boardStructure, &m) {
				var p board.PvEntry
				p.PosKey = boardStructure.PosKey
				p.Move = mv
				boardStructure.PvTable.PvTableEntries[boardStructure.PosKey] = p
				boardStructure.PvArray[boardStructure.PvTable.NumEntries] = mv
				boardStructure.PvTable.NumEntries++
			}
		} else {
			fmt.Println("Invalid move")
		}
		if search.IsRepetition(&boardStructure) {
			fmt.Println("position repeated")
		}
	}

}
