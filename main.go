package main

import (
	"bufio"
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/Greeshmanth1909/shadowfax/eval"
	"github.com/Greeshmanth1909/shadowfax/position"
	"github.com/Greeshmanth1909/shadowfax/util"
	"os"
)

func main() {
	util.InitAll()
	fmt.Printf("%v\nStatus: running\n", board.Name)
	reader := bufio.NewReader(os.Stdin)

	startString := position.StartPosition
	var boardStructure board.S_Board
	position.Parse_FEN(&startString, &boardStructure)
	position.PrintBoard(&boardStructure)

	var list eval.S_MoveList
	eval.GenerateAllMoves(&boardStructure, &list)
	eval.PrintMoveList(&list)

	fmt.Println(boardStructure.EnP)

	for i := 0; i < list.Count; i++ {
		mv := list.MoveList[i]
		eval.MakeMove(&boardStructure, &mv)
		position.PrintBoard(&boardStructure)
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		fmt.Println("Taking move")
		eval.TakeMove(&boardStructure)
		position.PrintBoard(&boardStructure)
		newT, _ := reader.ReadString('\n')
		fmt.Println(newT)
	}
}
