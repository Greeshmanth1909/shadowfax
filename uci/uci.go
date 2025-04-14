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
	// Split line only once and reuse the result
	lineList := strings.Split(line, " ")
	startPos := position.StartPosition

	if lineList[1] == "startpos" {
		position.Parse_FEN(&startPos, brd)
	} else if lineList[1] == "fen" {
		// Only perform the join operation when needed
		fen := strings.Join(lineList[2:8], " ")
		position.Parse_FEN(&fen, brd)
	} else {
		position.Parse_FEN(&startPos, brd)
	}

	// Find the "moves" index only once
	movesIndex := -1
	for i, val := range lineList {
		if val == "moves" {
			movesIndex = i
			break
		}
	}

	// Process moves if found
	if movesIndex != -1 && len(lineList) > movesIndex+1 {
		moveList := lineList[movesIndex+1:]
		fmt.Println(moveList)
		for _, move := range moveList {
			// Trim suffix only if needed
			if strings.HasSuffix(move, "\n") {
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

	// Split the line only once
	lineList := strings.Split(line, " ")

	// Cache string contains checks for reuse
	hasInfinite := strings.Contains(line, "infinite")
	hasBinc := strings.Contains(line, "binc")
	hasWinc := strings.Contains(line, "winc")
	hasWtime := strings.Contains(line, "wtime")
	hasBtime := strings.Contains(line, "btime")
	hasMovesToGo := strings.Contains(line, "movestogo")
	hasMoveTime := strings.Contains(line, "movetime")
	hasDepth := strings.Contains(line, "depth")

	if hasInfinite {
		// Handle infinite search
	}

	if hasBinc && brd.Side == board.BLACK {
		index := findIndex("binc", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}
		inc, _ = strconv.Atoi(lineList[index+1])
	}

	if hasWinc && brd.Side == board.WHITE {
		index := findIndex("winc", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}
		inc, _ = strconv.Atoi(lineList[index+1])
	}

	if hasWtime && brd.Side == board.WHITE {
		index := findIndex("wtime", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}
		Time, _ = strconv.Atoi(lineList[index+1])
	}

	if hasBtime && brd.Side == board.BLACK {
		index := findIndex("btime", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}
		Time, _ = strconv.Atoi(lineList[index+1])
	}

	if hasMovesToGo {
		index := findIndex("movestogo", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}

		valueStr := lineList[index+1]
		if strings.HasSuffix(valueStr, "\n") {
			valueStr = strings.TrimSuffix(valueStr, "\n")
		}
		movesToGo, _ = strconv.Atoi(valueStr)
	}

	if hasMoveTime {
		index := findIndex("movetime", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}

		valueStr := lineList[index+1]
		if strings.HasSuffix(valueStr, "\n") {
			valueStr = strings.TrimSuffix(valueStr, "\n")
		}
		mt, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Println(err)
		}
		moveTime = mt
	}

	if hasDepth {
		index := findIndex("depth", lineList)
		if index == -1 {
			fmt.Println("invalid index returned")
			return
		}

		valueStr := lineList[index+1]
		if strings.HasSuffix(valueStr, "\n") {
			valueStr = strings.TrimSuffix(valueStr, "\n")
		}
		d, err := strconv.Atoi(valueStr)
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
