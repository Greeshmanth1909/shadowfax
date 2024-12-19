package util

import (
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/Greeshmanth1909/shadowfax/eval"
)

func InitAll() {
	board.InitSquares64()
	board.InitSquares120()
	board.InitBitMasks()
	board.InitHashKeys()
	board.InitFileRankArrays()
	eval.InitMvvLva()
}
