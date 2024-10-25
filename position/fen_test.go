package position

import (
	"github.com/Greeshmanth1909/shadowfax/board"
	"testing"
)

func TestSquare(t *testing.T) {
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
