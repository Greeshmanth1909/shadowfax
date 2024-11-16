package eval

import (
	"github.com/Greeshmanth1909/shadowfax/board"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

const MAXPOSITIONMOVES = 256

type S_MoveList struct {
	MoveList [MAXPOSITIONMOVES]S_Move
	Count    int
}

func AddQuietMove(brd *board.S_Board, move uint32, list *S_MoveList) {
	list.MoveList[list.Count].Move = move
	list.MoveList[list.Count].Score = 0
	list.Count++
}

func AddCaptureMove(brd *board.S_Board, move uint32, list *S_MoveList) {
	list.MoveList[list.Count].Move = move
	list.MoveList[list.Count].Score = 0
	list.Count++
}
func AddEnPassantMove(brd *board.S_Board, move uint32, list *S_MoveList) {
	list.MoveList[list.Count].Move = move
	list.MoveList[list.Count].Score = 0
	list.Count++
}

func GenerateAllMoves(brd *board.S_Board, list *S_MoveList) {
	list.Count++
}
