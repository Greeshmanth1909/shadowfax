package util

import (
	"github.com/Greeshmanth1909/shadowfax/board"
)

func InitAll() {
	board.InitSquares64()
	board.InitSquares120()
	board.InitBitMasks()
	board.InitHashKeys()
}