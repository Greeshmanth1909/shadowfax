package board

import (
	"fmt"
	"math/bits"
)

func PrintBitBoard(bb uint64) {
	var ind uint64 = 1
	for rank := RANK_8; rank >= RANK_1; rank-- {
		for file := FILE_A; file <= FILE_H; file++ {
			sq := FRtoSq120(file, rank)
			if (ind<<sq)&bb != 0 {
				fmt.Print("x ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Print("\n")
	}
}

func CountBits(bb uint64) int {
	return bits.OnesCount64(bb)
}

func PopBits(bb *uint64) int {
	index := bits.TrailingZeros64(*bb)
	var one uint64 = 1
	one = one << index
	mask := ^one
	*bb = *bb & mask
	return index
}
