package board

// define board of 120 int array
var Board [120]uint64

const Name string = "Shadowfax v1.0.0"

var BrdSqrNum int = 120

// define piece types
type Piece int

const (
	EMPTY Piece = iota
	Wp
	Wn
	Wb
	Wr
	Wq
	Wk
	Bp
	Bn
	Bb
	Br
	Bq
	Bk
)

// define files and ranks
type File int

const (
	FILE_A File = iota
	FILE_B
	FILE_C
	FILE_D
	FILE_E
	FILE_F
	FILE_G
	FILE_H
	FILE_NONE
)

type Rank int

const (
	RANK_1 Rank = iota
	RANK_2
	RANK_3
	RANK_4
	RANK_5
	RANK_6
	RANK_7
	RANK_8
	RANK_NONE
)

// define colors
type Color int

const (
	WHITE Color = iota
	BLACK
	BOTH
)

// define squares
type Square int

const (
	A1 Square = iota + 21
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2 = iota + 2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3 = iota + 2
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4 = iota + 2 // 51
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5 = iota + 2 // 61
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6 = iota + 2 // 71
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7 = iota + 2 // 81
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8 = iota + 2 // 91
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	NO_SQ Square = 99 // Used for invalid squares
)

type S_Board struct {
    Pieces [BrdSqrNum]int
    Pawns [3]uint64
    KingSquare [2]int
    EnP int
    Side int
    FiftyMove int
    Ply int
    HisPly int
    PosKey uint64
    PieceNum int
    BigPiece [3]int
    MinPiece [3]int
    MajPiece [3]int
}
