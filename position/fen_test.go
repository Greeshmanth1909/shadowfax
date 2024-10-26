package position

import (
	"github.com/Greeshmanth1909/shadowfax/board"
	"github.com/Greeshmanth1909/shadowfax/util"
	"testing"
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
	if boardStructure.Pieces[board.Square64to120[4]] != board.Wk {
		t.Fatalf("test fen parser: want %v, have %v", board.Wk, boardStructure.Pieces[board.Square64to120[4]])
	}
}

func TestCastlePareser(t *testing.T) {
	bits := getCastlingPerm("KQkq")
	if bits != 15 {
		t.Fatalf("something up with castling want %v, got %v", 16, bits)
	}
	bits = getCastlingPerm("Kk")
	if bits != 5 {
		t.Fatalf("something up with castling")
	}
	bits = getCastlingPerm("")
	if bits != 0 {
		t.Fatalf("something up with castling")
	}
}
