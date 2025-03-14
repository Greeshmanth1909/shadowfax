package board

import (
	"fmt"
	"testing"
)

func TestPiece(t *testing.T) {
	// IMPORTANT: util package cannot be imported because as board is its dependency and importing it would result in a cycle
	InitSquares64()
	InitSquares120()
	have := fmt.Sprintf("%v", Wp)
	want := fmt.Sprintf("%v", 1)
	if have != want {
		t.Fatalf("TestPiece mismatch: have %v want %v", have, want)
	}
}

func TestSquaresArr(t *testing.T) {
	have := Square64to120
	want := [64]int{
		21, 22, 23, 24, 25, 26, 27, 28,
		31, 32, 33, 34, 35, 36, 37, 38,
		41, 42, 43, 44, 45, 46, 47, 48,
		51, 52, 53, 54, 55, 56, 57, 58,
		61, 62, 63, 64, 65, 66, 67, 68,
		71, 72, 73, 74, 75, 76, 77, 78,
		81, 82, 83, 84, 85, 86, 87, 88,
		91, 92, 93, 94, 95, 96, 97, 98,
	}
	if have != want {
		t.Fatalf("Squares Array 64 Mismatch: have- %v \nwant- %v", have, want)
	}

	have120 := Square120to64
	want120 := [120]int{
		65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
		65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
		65, 0, 1, 2, 3, 4, 5, 6, 7, 65,
		65, 8, 9, 10, 11, 12, 13, 14, 15, 65,
		65, 16, 17, 18, 19, 20, 21, 22, 23, 65,
		65, 24, 25, 26, 27, 28, 29, 30, 31, 65,
		65, 32, 33, 34, 35, 36, 37, 38, 39, 65,
		65, 40, 41, 42, 43, 44, 45, 46, 47, 65,
		65, 48, 49, 50, 51, 52, 53, 54, 55, 65,
		65, 56, 57, 58, 59, 60, 61, 62, 63, 65,
		65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
		65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
	}

	if have120 != want120 {
		t.Fatalf("Squares Array 64 Mismatch: have- %v \nwant- %v", have, want)
	}
}

func TestPopCount(t *testing.T) {
	var num uint64 = 9
	count := CountBits(num)
	if count != 2 {
		t.Fatalf("population count error: have %v, want %v", count, 2)
	}

	num = 24
	ind := PopBits(&num)
	if num != 16 {
		t.Fatalf("PopBits error: popbits(24) == 16; got %v", num)
	}
	if ind != 3 {
		t.Fatalf("PopBits index error: popbits(24) == 16 (pop index should equal 3); got %v", ind)
	}

}
