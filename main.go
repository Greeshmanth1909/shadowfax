package main

import (
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
)

func main() {
	s1 := board.InitSquares64()
	s2 := board.InitSquares120()
	fmt.Printf("%v\nStatus: running\n", board.Name)

	fmt.Println(s1)
	fmt.Println(s2)
}
