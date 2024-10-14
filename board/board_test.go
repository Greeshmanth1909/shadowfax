package board

import (
    "testing"
    "fmt"
)

func TestPiece(t *testing.T) {
    have := fmt.Sprintf("%v", Wp)
    want := fmt.Sprintf("%v", 1)
    if have != want {
        t.Fatalf("TestPiece mismatch: have %v want %v", have, want)
    }
}
