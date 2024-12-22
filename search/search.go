package search

import (
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/Greeshmanth1909/shadowfax/eval"
	"time"
)

const Inf int = 30000
const Mate int = 29000

func CheckUp(info *board.S_SearchInfo) {
	// check if time up or interrupt from GUI
    startTime := info.StartTime
    dur := time.Since(startTime).Milliseconds()
    if info.TimeSet && dur > info.StopTime {
       info.Stopped = true 
    }

}
func IsRepetition(brd *board.S_Board) bool {
	pKey := brd.PosKey

	for i := brd.HisPly - brd.FiftyMove; i < brd.HisPly; i++ {
		if brd.History[i].PosKey == pKey {
			return true
		}
	}
	return false
}

func SearchPositions(brd *board.S_Board, info *board.S_SearchInfo) {
	var bestMove uint32
	var bestScore int
	var currentDepth int
	var pvMoves int

	bestScore = -Inf
	ClearForSearch(brd, info)

	for currentDepth = 1; currentDepth < info.Depth; currentDepth++ {
		bestScore = AlphaBeta(-Inf, Inf, currentDepth, 1, brd, info)
        if info.Stopped {
            break
        }
        pvMoves = eval.GetPvLine(currentDepth, brd)
        bestMove = brd.PvArray[0]

        var m eval.S_Move
        m.Move = bestMove

        fmt.Printf("info score cp %v depth %v nodes %v time %v pv", bestScore, currentDepth, info.Nodes, time.Since(info.StartTime).Milliseconds())
        for i := 0; i < pvMoves; i++ {
            val := brd.PvArray[i]
            var mv eval.S_Move
            mv.Move = val
            eval.PrintMove(&mv)
        }
        fmt.Print("\n")
        fmt.Printf("ordering: %v\n", info.FhF/info.Fh)
        fmt.Printf("bestmove")
        eval.PrintMove(&m)
        fmt.Print("\n")
	}
}

func ClearForSearch(brd *board.S_Board, info *board.S_SearchInfo) {
	for i := 0; i < 13; i++ {
		for j := 0; j < board.BrdSqrNum; j++ {
			brd.SearchHistoryArray[i][j] = 0
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < board.MAXDEPTH; j++ {
			brd.SearchKillers[i][j] = 0
		}
	}

	// Clear pvtable
	for key := range brd.PvTable.PvTableEntries {
		delete(brd.PvTable.PvTableEntries, key)
	}
	brd.PvTable.NumEntries = 0
	brd.Ply = 0

	info.StartTime = time.Now()
	info.Stopped = false
	info.Nodes = 0
	info.Fh = 0
	info.FhF = 0
}

func Quiescence(alpha, beta int, brd *board.S_Board, info *board.S_SearchInfo) int {
	// board.CheckBoard(brd)
    if info.Nodes & 2047 == 0 {
       CheckUp(info) 
    }

	info.Nodes++
	if IsRepetition(brd) || brd.FiftyMove >= 100 {
		return 0
	}

	score := EvalPosition(brd)
	if score >= beta {
		return beta
	}
	if score > alpha {
		alpha = score
	}

	var list eval.S_MoveList
	eval.GenerateAllCaps(brd, &list)

	legal := 0
	oldAlpha := alpha
	bestMove := uint32(0)
	score = -Inf

	for i := 0; i < list.Count; i++ {
		PickNextMove(i, &list)
		mv := list.MoveList[i]
		if !eval.MakeMove(brd, &mv) {
			continue
		}
		legal++
		score = -Quiescence(-beta, -alpha, brd, info)
		eval.TakeMove(brd)

        if info.Stopped {
            return 0
        }

		if score > alpha {
			if score >= beta {
				if legal == 1 {
					info.FhF++
				}
				info.Fh++
				return beta
			}
			alpha = score
			bestMove = mv.Move
		}
	}

	if alpha != oldAlpha {
		board.StorePvMove(brd, brd.PosKey, bestMove)
	}
	return alpha
}

func AlphaBeta(alpha, beta, depth, doNull int, brd *board.S_Board, info *board.S_SearchInfo) int {
	// board.CheckBoard(brd)
	if depth == 0 {
		return Quiescence(alpha, beta, brd, info)
	}

    if info.Nodes & 2047 == 0 {
       CheckUp(info) 
    }

	info.Nodes++
	if IsRepetition(brd) || brd.FiftyMove >= 100 {
		return 0
	}

	if brd.Ply > board.MAXDEPTH-1 {
		return EvalPosition(brd)
	}

	var list eval.S_MoveList
	eval.GenerateAllMoves(brd, &list)

	legal := 0
	oldAlpha := alpha
	bestMove := uint32(0)
	score := -Inf
	pvMove := board.ProbePvTable(brd, brd.PosKey)

	if pvMove != 0 {
		for i := 0; i < list.Count; i++ {
			mv := list.MoveList[i]
			if mv.Move == pvMove {
				list.MoveList[i].Score = 2000000
				break
			}
		}
	}

	for i := 0; i < list.Count; i++ {
		PickNextMove(i, &list)
		mv := list.MoveList[i]
		if !eval.MakeMove(brd, &mv) {
			continue
		}
		legal++
		score = -AlphaBeta(-beta, -alpha, depth-1, 1, brd, info)
		eval.TakeMove(brd)

        if info.Stopped {
            return 0
        }

		if score > alpha {
			if score >= beta {
				if legal == 1 {
					info.FhF++
				}
				info.Fh++
				if eval.GetCapturedPiece(&mv) == 0 {
					brd.SearchKillers[1][brd.Ply] = brd.SearchKillers[0][brd.Ply]
					brd.SearchKillers[0][brd.Ply] = mv.Move
				}
				return beta
			}
			alpha = score
			bestMove = mv.Move
			if eval.GetCapturedPiece(&mv) == 0 {
				brd.SearchHistoryArray[brd.Pieces[eval.GetFromSquare(&mv)]][eval.GetToSquare(&mv)] += uint32(depth)

			}
		}
	}

	if legal == 0 {
		if eval.SquareAttacked(board.Square(brd.KingSquare[brd.Side]), brd.Side^1, brd) {
			return -Mate + brd.Ply
		}
		return 0
	}

	if alpha != oldAlpha {
		board.StorePvMove(brd, brd.PosKey, bestMove)
	}
	return alpha
}

func PickNextMove(moveNum int, list *eval.S_MoveList) {
	var temp eval.S_Move
	var index, bestScore int
	var bestNum int = moveNum
	for index = moveNum; index < list.Count; index++ {
		if list.MoveList[index].Score > bestScore {
			bestScore = list.MoveList[index].Score
			bestNum = index
		}
	}
	temp = list.MoveList[moveNum]
	list.MoveList[moveNum] = list.MoveList[bestNum]
	list.MoveList[bestNum] = temp
}
