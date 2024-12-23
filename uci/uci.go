package uci

import (
	"bufio"
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/Greeshmanth1909/shadowfax/eval"
	"github.com/Greeshmanth1909/shadowfax/position"
	"os"
	"strings"
)

func UciLoop() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("id name %v\n", board.Name)
	fmt.Printf("id author Greeshmanth\n")
	fmt.Println("uciok")

	var brd = &board.S_Board{}
	var info = &board.S_SearchInfo{}

	board.InitPvTable(brd)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		if input == "\n" {
			continue
		}

		if input == "isready\n" {
			fmt.Println("readyok")
		} else if strings.HasPrefix(input, "position") {
			ParsePosition(input, brd)
		} else if strings.HasPrefix(input, "ucinewgame") {
			ParsePosition("position startpos\n", brd)
		} else if strings.HasPrefix(input, "go") {
			ParseGo(input, info, brd)
		} else if input == "quit\n" {
			info.Quit = true
			break
		} else if input == "uci\n" {
			fmt.Printf("id name %v\n", board.Name)
			fmt.Printf("id author Greeshmanth\n")
			fmt.Printf("uciok\n")
		}

		if info.Quit {
			break
		}

	}

}

func ParsePosition(line string, brd *board.S_Board) {
	lineList := strings.Split(line, " ")
	startPos := position.StartPosition

	if lineList[1] == "startpos" {
		position.Parse_FEN(&startPos, brd)
	} else if lineList[1] == "fen" {
		fen := strings.Join(lineList[2:8], " ")
		position.Parse_FEN(&fen, brd)
	} else {
		position.Parse_FEN(&startPos, brd)
	}

	if strings.Contains(line, "moves") {
		var index int
		for i, val := range lineList {
			if val == "moves" {
				index = i
			}
		}
		moveList := lineList[index+1:]
		fmt.Println(moveList)
		for _, move := range moveList {
			mv := eval.ParseMove(move, brd)
			if mv == 0 {
				break
			}
			var m eval.S_Move
			m.Move = mv
			eval.MakeMove(brd, &m)
			brd.Ply = 0
		}
	}
	position.PrintBoard(brd)
}

func ParseGo(line string, info *board.S_SearchInfo, brd *board.S_Board) {
	return
}
