package main

import (
	// "bufio"
	// "fmt"
	// "github.com/Greeshmanth1909/shadowfax/board"
	// "github.com/Greeshmanth1909/shadowfax/eval"
	// "github.com/Greeshmanth1909/shadowfax/position"
	// "github.com/Greeshmanth1909/shadowfax/search"
	"github.com/Greeshmanth1909/shadowfax/uci"
	"github.com/Greeshmanth1909/shadowfax/util"
	// "os"
	// "time"
)

func main() {
	util.InitAll()
	uci.UciLoop()
}
