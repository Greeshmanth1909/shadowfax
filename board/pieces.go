package board

// empty, pawn, bishop, night, rook, queen, king, pawn, bishop, night, rook, queen, king
var BigPiece = [13]bool{false, false, true, true, true, true, true, false, true, true, true, true, true}
var MajPiece = [13]bool{false, false, false, false, true, true, true, false, false, false, true, true, true}
var MinPiece = [13]bool{false, false, true, true, false, false, false, false, true, true, false, false, false}
var PieceVal = [13]int{0, 100, 325, 325, 550, 1000, 50000, 100, 325, 325, 550, 1000, 50000}
var PieceCol = [13]Color{BOTH, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, BLACK, BLACK, BLACK, BLACK, BLACK, BLACK}
