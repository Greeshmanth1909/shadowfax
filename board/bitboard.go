package board

import (
	"fmt"
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
