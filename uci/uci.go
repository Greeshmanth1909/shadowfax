package uci

import (
	"bufio"
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
	"os"
)

func UciLoop() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("id name %v\n", board.Name)
	fmt.Printf("id author Greeshmanth\n")
	fmt.Println("uciok")

	var brd *board.S_Board
	var info *board.S_SearchInfo

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
		} else if input[:8] == "position" {
			ParsePosition(input, brd)
		} else if input[:10] == "ucinewgame" {
			ParsePosition("position startpos\n", brd)
		} else if input[:2] == "go" {
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
	return
}

func ParseGo(line string, info *board.S_SearchInfo, brd *board.S_Board) {
	return
}
