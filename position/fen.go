package position

import (
    "strings"
    "strconv"
    "fmt"
    "github.com/Greeshmanth1909/shadowfax/board"
    "errors"
)

const StartPosition string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func Parse_FEN(fen *string, brd *board.S_Board) error {
    splitFen := strings.Split(*fen, " ")
    if len(splitFen) != 6 {
        return errors.New("Invalid Fen")
    }
    fenString := splitFen[0]
    side := splitFen[1]
    castling := splitFen[2]
    enp := splitFen[3]
    halfMove := splitFen[4]
    fullMove := splitFen[5]

    index := 0
    //piece := board.EMPTY

    for _, char := range fenString {
        switch char {
            case 'r':
                // Black rook
                brd.Pieces[board.Square64to120[index]] = board.Br
                index++
            case 'n':
                // Black night
                brd.Pieces[board.Square64to120[index]] = board.Bn
                index++
            case 'b':
                // Black bishop
                brd.Pieces[board.Square64to120[index]] = board.Bb
                index++
            case 'q':
                brd.Pieces[board.Square64to120[index]] = board.Bq
                index++
            case 'k':
                brd.Pieces[board.Square64to120[index]] = board.Bk
                index++
            case 'p':
                brd.Pieces[board.Square64to120[index]] = board.Bp
                index++

            // White pieces
            case 'R':
                brd.Pieces[board.Square64to120[index]] = board.Wr
                index++
            case 'N':
                brd.Pieces[board.Square64to120[index]] = board.Wn
                index++
            case 'B':
                brd.Pieces[board.Square64to120[index]] = board.Wb
                index++
            case 'Q':
                brd.Pieces[board.Square64to120[index]] = board.Wq
                index++
            case 'K':
                brd.Pieces[board.Square64to120[index]] = board.Wk
                index++
            case 'P':
                brd.Pieces[board.Square64to120[index]] = board.Wp
                index++

            // Empty Squares
            case '1', '2', '3', '4', '5', '6', '7', '8':
                inc, _ := strconv.Atoi(string(char))
                index += inc
            case '/':
                index++
            default:
                return errors.New(fmt.Sprintf("invalid character in fen string %v", char))
        }
    }
    if side == "w" {
        brd.Side = board.WHITE
    } else if side == "b" {
        brd.Side = board.BLACK
    } else {
        return errors.New(fmt.Sprintf("Invalid side, %v", side))
    }

    if enp != "-" {
        brd.EnP = convertSquareStringToSquare(enp)
    }

    brd.CastlePerm = getCastlingPerm(castling)
    brd.FiftyMove, _ = strconv.Atoi(halfMove)
    brd.Ply, _ = strconv.Atoi(fullMove)
    return nil
}

func getCastlingPerm(str string) int {
    castling := 0
    if strings.Contains(str, "K") {
        castling |= 1
    }
    if strings.Contains(str, "Q") {
        castling |= 1 << 1
    }
    if strings.Contains(str, "k") {
        castling |= 1 << 2
    }
    if strings.Contains(str, "q") {
        castling |= 1 << 3
    }
    return castling
}

func convertSquareStringToSquare(square string) board.Square {
    file := square[0]
    rank, _ := strconv.Atoi(string(square[1]))
    r := 0
    switch file {
        case 'a':
            r = 1
        case 'b':
            r = 2
        case 'c':
            r = 3
        case 'd':
            r = 4
        case 'e':
            r = 5
        case 'f':
            r = 6
        case 'g':
            r = 7
        case 'h':
            r = 8
    }
    return board.Square((rank + 1) * 10 + r)
}
