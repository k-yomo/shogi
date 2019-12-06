package shogi

import (
	"fmt"
)

type Board [][]Piece

func NewBoard(firstPlayer Player, secondPlayer Player) (board Board) {
	board = append(board, initializeBase(secondPlayer)...)
	board = append(board, initializeMiddleZone()...)
	board = append(board, initializeBase(firstPlayer)...)
	return
}

// initializeBase initializes player's base zone
func initializeBase(p Player) [][]Piece {
	rows := [][]Piece{
		{NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p), NewPawn(p)},
		{nil,NewBishop(p) , nil, nil, nil, nil, nil, NewRook(p), nil},
		{NewLance(p), NewKnight(p), NewSilver(p), NewGold(p), NewKing(p), NewGold(p), NewSilver(p), NewKnight(p), NewLance(p)},
	}
	if !p.IsFirstPlayer() {
		reversePieceOrder(rows[1])
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
	p = b[pos.Y][pos.X]
	return p, p != nil
}

func (b Board) MovePiece(currentPlayer Player, curPos, distPos *Position) (isSucceeded bool) {
	piece, exist := b.FindPiece(curPos)
	if !exist {
		fmt.Println(fmt.Sprintf("piece doesn't exist at %v", curPos))
		return false
	}
	if !IsSamePlayer(currentPlayer, piece.Owner()) {
		fmt.Println(fmt.Sprintf("piece doesn't belong to %s", piece.Owner().Name()))
		return false
	}
	distPositionPiece, distPositionPieceExist := b.FindPiece(distPos)
	if distPositionPieceExist && IsSamePlayer(currentPlayer, distPositionPiece.Owner()) {
		fmt.Println(fmt.Sprintf("there is current user's piece at %v", distPos))
		return false
	}
	if !piece.IsMovableTo(curPos, distPos) {
		fmt.Println(fmt.Sprintf("the piece can't be moved to %v", distPos))
		return false
	}

	if distPositionPieceExist {
		currentPlayer.TakePiece(distPositionPiece)
	}
	b[distPos.Y][distPos.X] = b[curPos.Y][curPos.X]
	b[curPos.Y][curPos.X] = nil
	fmt.Println(currentPlayer.PiecesInHand())
	return true
}

func reversePieceOrder(s []Piece) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
