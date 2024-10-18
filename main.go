package main

import (
	"fmt"
	"github.com/Greeshmanth1909/shadowfax/board"
)

func main() {
    s1 := board.InitSquares64()
	fmt.Printf("%v\nStatus: running\n", board.Name)
    for _, val := range s1 {
        fmt.Printf("%v ", val)
    }
}
