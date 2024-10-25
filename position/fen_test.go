package position

import (
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/Greeshmanth1909/shadowfax/util"
	"testing"
    "fmt"
)

func TestSquare(t *testing.T) {
    // IMPORTANT: util.InitAll must be called within the scope of any unit test to ensure proper results
    util.InitAll()
	have := convertSquareStringToSquare("e4")
	want := board.E4
	if have != want {
		t.Fatalf("fen square conv error, have %v want %v", have, want)
	}

	have = convertSquareStringToSquare("a1")
	want = board.A1
	if have != want {
		t.Fatalf("fen square conv error, have %v want %v", have, want)
	}

	have = convertSquareStringToSquare("h8")
	want = board.H8
	if have != want {
		t.Fatalf("fen square conv error, have %v want %v", have, want)
	}

	have = convertSquareStringToSquare("h1")
	want = board.H1
	if have != want {
		t.Fatalf("fen square conv error, have %v want %v", have, want)
	}

	have = convertSquareStringToSquare("a8")
	want = board.A8
	if have != want {
		t.Fatalf("fen square conv error, have %v want %v", have, want)
	}

}

func TestParseFen(t *testing.T) {
    startString := StartPosition
    var boardStructure board.S_Board
    Parse_FEN(&startString, &boardStructure)
    fmt.Println(boardStructure.Pieces[board.Square64to120[4]] == board.Bk)
    if boardStructure.Pieces[board.Square64to120[4]] != board.Bk {
        t.Fatalf("test fen parser: want %v, have %v", boardStructure.Pieces[board.Square64to120[4]], board.Bk)
    }
}
