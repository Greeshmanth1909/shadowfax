package uci

import (
	"bufio"
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/Greeshmanth1909/shadowfax/eval"
	"github.com/Greeshmanth1909/shadowfax/position"
	"github.com/Greeshmanth1909/shadowfax/search"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var startInterruptPeek = false
var m = sync.Mutex{}

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
			if strings.Contains(move, "\n") {
				move = strings.TrimSuffix(move, "\n")
			}
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
	depth := -1
	movesToGo := 30
	moveTime := -1
	Time := -1
	inc := 0

	info.TimeSet = false

	lineList := strings.Split(line, " ")

	if strings.Contains(line, "infinite") {
	}

	if strings.Contains(line, "binc") && brd.Side == board.BLACK {
		index := findIndex("binc", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}
		inc, _ = strconv.Atoi(lineList[index+1])
	}

	if strings.Contains(line, "winc") && brd.Side == board.WHITE {
		index := findIndex("winc", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}
		inc, _ = strconv.Atoi(lineList[index+1])
	}

	if strings.Contains(line, "wtime") && brd.Side == board.WHITE {
		index := findIndex("wtime", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}
		Time, _ = strconv.Atoi(lineList[index+1])
	}

	if strings.Contains(line, "btime") && brd.Side == board.BLACK {
		index := findIndex("btime", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}
		Time, _ = strconv.Atoi(lineList[index+1])
	}

	if strings.Contains(line, "movestogo") {
		index := findIndex("movestogo", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}
		movesToGo, _ = strconv.Atoi(strings.TrimSuffix(lineList[index+1], "\n"))
	}

	if strings.Contains(line, "movetime") {
		index := findIndex("movetime", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}
		mt, err := strconv.Atoi(strings.TrimSuffix(lineList[index+1], "\n"))
		if err != nil {
			fmt.Println(err)
		}
		moveTime = mt
	}

	if strings.Contains(line, "depth") {
		index := findIndex("depth", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}
		d, err := strconv.Atoi(strings.TrimSuffix(lineList[index+1], "\n"))
		if err != nil {
			fmt.Println(err)
		}
		depth = d + 1
		fmt.Println("PRINTING DEPTH ", depth)
	}

	if moveTime != -1 {
		Time = moveTime
		movesToGo = 1
	}

	info.StartTime = time.Now()
	info.Depth = depth

	if Time != -1 {
		info.TimeSet = true
		Time /= movesToGo
		Time -= 50
		info.StopTime = int64(Time + inc)
	}

	if depth == -1 {
		info.Depth = board.MAXDEPTH
	}
	fmt.Printf("time:%v start:%v stop:%v depth:%v timeset:%v\n", Time, info.StartTime, info.StopTime, info.Depth, info.TimeSet)

	search.SearchPositions(brd, info) // This function takes a while to finish
}

func findIndex(str string, arr []string) int {
	for i, val := range arr {
		if val == str {
			return i
		}
	}
	return -1
}

// func InputWaiting() bool {
// 	fd := int(os.Stdin.Fd())
// 	var readfds syscall.FdSet
// 	readfds.Set(fd)
//
// 	tv := syscall.Timeval{}
// 	err := syscall.Select(fd+1, &readfds, nil, nil, &tv)
// 	if err != nil {
// 		return false
// 	}
// 	return readfds.IsSet(fd)
// }
