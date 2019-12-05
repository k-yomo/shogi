package shogi

import "sort"

type Board [][]Piece

func NewBoard(firstPlayer Player, secondPlayer Player) (board Board) {
	board = append(board, initializeBase(firstPlayer)...)
	board = append(board, initializeMiddleZone()...)
	board = append(board, initializeBase(secondPlayer)...)
	return
}

// initializeBase initializes player's base zone
func initializeBase(p Player) [][]Piece {
	rows := [][]Piece{
		{NewLance(p), NewKnight(p), NewSilver(p), NewGold(p), NewKing(p), NewGold(p), NewSilver(p), NewKnight(p), NewLance(p)},
		{nil, NewRook(p), nil, nil, nil, nil, nil, NewBishop(p), nil},
		{NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p)},
	}
	if !p.IsFirstPlayer() {
		sort.Slice(rows[1], func(i, j int) bool {
			return i > j
		})
		rows[0], rows[2] = rows[2], rows[0]
	}
	return rows
}

// initializeBase initializes middle empty 3 rows
func initializeMiddleZone() [][]Piece {
	rows := make([][]Piece, 3)
	for i := range rows {
		rows[i] = []Piece{nil, nil, nil, nil, nil, nil, nil, nil, nil}
	}
	return rows
}

func (b Board) String() string {
	var boardString string
	for _, row := range b {
		for _, piece := range row {
			if piece != nil {
				boardString += piece.ShortName()
			} else {
				boardString += "ー"
			}
		}
		boardString += "\n"
	}
	return boardString
}

func (b Board) FindPiece(pos *Position) (p Piece, exist bool) {
	p = b[pos.Y-1][pos.X-1]
	return p, p != nil
}
