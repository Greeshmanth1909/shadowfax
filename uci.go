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
var searchRunning = false
var currentInfo *board.S_SearchInfo

var benchmarkPositions = []string{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
	"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 10",
	"8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 11",
	"4rrk1/pp1n3p/3q2pQ/2p1pb2/2PP4/2P3N1/P2B2PP/4RRK1 b - - 7 19",
	"rq3rk1/ppp2ppp/1bnpb3/3N2B1/3NP3/7P/PPPQ1PP1/2KR3R w - - 7 14 moves d4e6",
	"r1bq1r1k/1pp1n1pp/1p1p4/4p2Q/4Pp2/1BNP4/PPP2PPP/3R1RK1 w - - 2 14 moves g2g4",
	"r3r1k1/2p2ppp/p1p1bn2/8/1q2P3/2NPQN2/PPP3PP/R4RK1 b - - 2 15",
	"r1bbk1nr/pp3p1p/2n5/1N4p1/2Np1B2/8/PPP2PPP/2KR1B1R w kq - 0 13",
	"r1bq1rk1/ppp1nppp/4n3/3p3Q/3P4/1BP1B3/PP1N2PP/R4RK1 w - - 1 16",
	"4r1k1/r1q2ppp/ppp2n2/4P3/5Rb1/1N1BQ3/PPP3PP/R5K1 w - - 1 17",
	"2rqkb1r/ppp2p2/2npb1p1/1N1Nn2p/2P1PP2/8/PP2B1PP/R1BQK2R b KQ - 0 11",
	"r1bq1r1k/b1p1npp1/p2p3p/1p6/3PP3/1B2NN2/PP3PPP/R2Q1RK1 w - - 1 16",
	"3r1rk1/p5pp/bpp1pp2/8/q1PP1P2/b3P3/P2NQRPP/1R2B1K1 b - - 6 22",
	"r1q2rk1/2p1bppp/2Pp4/p6b/Q1PNp3/4B3/PP1R1PPP/2K4R w - - 2 18",
	"4k2r/1pb2ppp/1p2p3/1R1p4/3P4/2r1PN2/P4PPP/1R4K1 b - - 3 22",
	"3q2k1/pb3p1p/4pbp1/2r5/PpN2N2/1P2P2P/5PP1/Q2R2K1 b - - 4 26",
	"6k1/6p1/6Pp/ppp5/3pn2P/1P3K2/1PP2P2/3N4 b - - 0 1",
	"3b4/5kp1/1p1p1p1p/pP1PpP1P/P1P1P3/3KN3/8/8 w - - 0 1",
	"2K5/p7/7P/5pR1/8/5k2/r7/8 w - - 0 1 moves g5g6 f3e3 g6g5 e3f3",
	"8/6pk/1p6/8/PP3p1p/5P2/4KP1q/3Q4 w - - 0 1",
	"7k/3p2pp/4q3/8/4Q3/5Kp1/P6b/8 w - - 0 1",
	"8/2p5/8/2kPKp1p/2p4P/2P5/3P4/8 w - - 0 1",
	"8/1p3pp1/7p/5P1P/2k3P1/8/2K2P2/8 w - - 0 1",
	"8/pp2r1k1/2p1p3/3pP2p/1P1P1P1P/P5KR/8/8 w - - 0 1",
	"8/3p4/p1bk3p/Pp6/1Kp1PpPp/2P2P1P/2P5/5B2 b - - 0 1",
	"5k2/7R/4P2p/5K2/p1r2P1p/8/8/8 b - - 0 1",
	"6k1/6p1/P6p/r1N5/5p2/7P/1b3PP1/4R1K1 w - - 0 1",
	"1r3k2/4q3/2Pp3b/3Bp3/2Q2p2/1p1P2P1/1P2KP2/3N4 w - - 0 1",
	"6k1/4pp1p/3p2p1/P1pPb3/R7/1r2P1PP/3B1P2/6K1 w - - 0 1",
	"8/3p3B/5p2/5P2/p7/PP5b/k7/6K1 w - - 0 1",
	"5rk1/q6p/2p3bR/1pPp1rP1/1P1Pp3/P3B1Q1/1K3P2/R7 w - - 93 90",
	"4rrk1/1p1nq3/p7/2p1P1pp/3P2bp/3Q1Bn1/PPPB4/1K2R1NR w - - 40 21",
	"r3k2r/3nnpbp/q2pp1p1/p7/Pp1PPPP1/4BNN1/1P5P/R2Q1RK1 w kq - 0 16",
	"3Qb1k1/1r2ppb1/pN1n2q1/Pp1Pp1Pr/4P2p/4BP2/4B1R1/1R5K b - - 11 40",
	"4k3/3q1r2/1N2r1b1/3ppN2/2nPP3/1B1R2n1/2R1Q3/3K4 w - - 5 1",
	"5k2/8/3PK3/8/8/8/8/8 w - - 0 1",
}

func HandleBench(depth int, brd *board.S_Board) {
	startTime := time.Now()
	var totalNodes uint64 = 0

	for _, fen := range benchmarkPositions {
		position.ResetBoard(brd)
		
		positionStr := "position fen " + fen
		ParsePosition(positionStr, brd)

		info := &board.S_SearchInfo{
			Depth:     depth,
			TimeSet:   false,
			Stopped:   false,
			StartTime: time.Now(),
		}

		// Create a quiet version of search for benchmarking
		runQuietSearch(brd, info)
		totalNodes += uint64(info.Nodes)
	}

	elapsed := time.Since(startTime).Milliseconds()
	if elapsed == 0 {
		elapsed = 1
	}
	nps := totalNodes * 1000 / uint64(elapsed)
	
	fmt.Printf("%d nodes %d nps\n", totalNodes, nps)
}

func runQuietSearch(brd *board.S_Board, info *board.S_SearchInfo) {
	search.ClearForSearch(brd, info)
	
	for currentDepth := 1; currentDepth <= info.Depth; currentDepth++ {
		search.AlphaBeta(-search.Inf, search.Inf, currentDepth, 1, brd, info)
		if info.Stopped {
			break
		}
	}
}

func UciLoop() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("id name %v JA\n", board.Name)
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

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if input == "isready" {
			fmt.Println("readyok")
		} else if strings.HasPrefix(input, "position") {
			ParsePosition(input, brd)
		} else if strings.HasPrefix(input, "ucinewgame") {
			ParsePosition("position startpos", brd)
		} else if strings.HasPrefix(input, "go") {
			ParseGo(input, info, brd)
		} else if strings.HasPrefix(input, "bench") {
			parts := strings.Split(input, " ")
			depth := 6 // default depth
			if len(parts) > 1 {
				if d, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
					depth = d
				}
			}
			HandleBench(depth, brd)
		} else if input == "stop" {
			m.Lock()
			if searchRunning && currentInfo != nil {
				currentInfo.Stopped = true
			}
			m.Unlock()
		} else if input == "quit" {
			m.Lock()
			if searchRunning && currentInfo != nil {
				currentInfo.Stopped = true
			}
			info.Quit = true
			m.Unlock()
			break
		} else if input == "uci" {
			fmt.Printf("id name %v JA\n", board.Name)
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

	movesIndex := -1
	for i, val := range lineList {
		if val == "moves" {
			movesIndex = i
			break
		}
	}

	if movesIndex != -1 && len(lineList) > movesIndex+1 {
		moveList := lineList[movesIndex+1:]
		fmt.Println(moveList)
		for _, move := range moveList {
			move = strings.TrimSpace(move)
			if move == "" {
				break
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
		if index == -1 || index+1 >= len(lineList) {
			fmt.Println("invalid index returned for binc")
			return
		}
		valueStr := strings.TrimSpace(lineList[index+1])
		inc, _ = strconv.Atoi(valueStr)
	}

	if hasWinc && brd.Side == board.WHITE {
		index := findIndex("winc", lineList)
		if index == -1 || index+1 >= len(lineList) {
			fmt.Println("invalid index returned for winc")
			return
		}
		valueStr := strings.TrimSpace(lineList[index+1])
		inc, _ = strconv.Atoi(valueStr)
	}

	if hasWtime && brd.Side == board.WHITE {
		index := findIndex("wtime", lineList)
		if index == -1 || index+1 >= len(lineList) {
			fmt.Println("invalid index returned for wtime")
			return
		}
		valueStr := strings.TrimSpace(lineList[index+1])
		Time, _ = strconv.Atoi(valueStr)
	}

	if hasBtime && brd.Side == board.BLACK {
		index := findIndex("btime", lineList)
		if index == -1 || index+1 >= len(lineList) {
			fmt.Println("invalid index returned for btime")
			return
		}
		valueStr := strings.TrimSpace(lineList[index+1])
		Time, _ = strconv.Atoi(valueStr)
	}

	if hasMovesToGo {
		index := findIndex("movestogo", lineList)
		if index == -1 || index+1 >= len(lineList) {
			fmt.Println("invalid index returned for movestogo")
			return
		}
		valueStr := strings.TrimSpace(lineList[index+1])
		movesToGo, _ = strconv.Atoi(valueStr)
	}

	if hasMoveTime {
		index := findIndex("movetime", lineList)
		if index == -1 || index+1 >= len(lineList) {
			fmt.Println("invalid index returned for movetime")
			return
		}
		valueStr := strings.TrimSpace(lineList[index+1])
		mt, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Printf("Error parsing movetime value '%s': %v\n", valueStr, err)
			return
		}
		moveTime = mt
	}

	if hasDepth {
		index := findIndex("depth", lineList)
		if index == -1 || index+1 >= len(lineList) {
			fmt.Println("invalid index returned for depth")
			return
		}
		valueStr := strings.TrimSpace(lineList[index+1])
		d, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Printf("Error parsing depth value '%s': %v\n", valueStr, err)
			return
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

	m.Lock()
	searchRunning = true
	currentInfo = info
	info.Stopped = false
	m.Unlock()

	go func() {
		search.SearchPositions(brd, info)
		m.Lock()
		searchRunning = false
		currentInfo = nil
		m.Unlock()
	}()
}

func findIndex(str string, arr []string) int {
	for i, val := range arr {
		if val == str {
			return i
		}
	}
	return -1
}
