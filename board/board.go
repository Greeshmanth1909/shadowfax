package board

// define board of 120 int array
var Board [120]uint64

const Name string = "Shadowfax v1.0.0"

const BrdSqrNum int = 120

const MAXGAMEMOVES int = 1028

// rank and file arrays
var FileArr [BrdSqrNum]File
var RankArr [BrdSqrNum]Rank

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

// define castling
type Castling int

const (
	WKCT Castling = 1
	WQCT          = 2
	BKCT          = 4
	BQCT          = 8
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
)

const (
	A2 Square = iota + 31
	B2
	C2
	D2
	E2
	F2
	G2
	H2
)

const (
	A3 Square = iota + 41
	B3
	C3
	D3
	E3
	F3
	G3
	H3
)

const (
	A4 Square = iota + 51
	B4
	C4
	D4
	E4
	F4
	G4
	H4
)

const (
	A5 Square = iota + 61
	B5
	C5
	D5
	E5
	F5
	G5
	H5
)

const (
	A6 Square = iota + 71
	B6
	C6
	D6
	E6
	F6
	G6
	H6
)

const (
	A7 Square = iota + 81
	B7
	C7
	D7
	E7
	F7
	G7
	H7
)

const (
	A8 Square = iota + 91
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

const (
	NO_SQ    Square = 99 // Used for invalid squares
	OFFBOARD Square = 100
)

type S_Board struct {
	Pieces     [BrdSqrNum]Piece
	Pawns      [3]uint64
	KingSquare [2]int
	EnP        Square
	Side       Color
	FiftyMove  int
	CastlePerm int
	Ply        int
	HisPly     int
	PosKey     uint64
	PieceNum   [13]int
	BigPiece   [2]int
	MinPiece   [2]int
	MajPiece   [2]int
	Material   [2]int
	History    [MAXGAMEMOVES]S_Undo
	PList      [13][10]int
	PvTable    *PvTable
}

type S_Undo struct {
	Move       uint32
	CastlePerm int
	EnP        Square
	FiftyMove  int
	PosKey     uint64
}

type PvEntry struct {
	PosKey uint64
	Move   uint32
}

type PvTable struct {
	PvTableEntries map[uint64]PvEntry
	NumEntries     int
}
