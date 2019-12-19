package shogi

import (
	"github.com/pkg/errors"
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
		{nil, NewRook(p), nil, nil, nil, nil, nil, NewBishop(p), nil},
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
		for i := 8; i >= 0; i-- {
			piece := row[i]
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
	p = b[pos.Y.Idx()][pos.X.Idx()]
	return p, p != nil
}

func (b Board) MovePiece(currentPlayer Player, curPos, distPos *Position) error {
	if err := b.IsMovablePieceTo(currentPlayer, curPos, distPos); err != nil {
		return err
	}
	distPositionPiece, distPositionPieceExist := b.FindPiece(distPos)
	if distPositionPieceExist {
		currentPlayer.TakePiece(distPositionPiece)
	}
	b[distPos.Y.Idx()][distPos.X.Idx()] = b[curPos.Y.Idx()][curPos.X.Idx()]
	b[curPos.Y.Idx()][curPos.X.Idx()] = nil
	return nil
}

func (b Board) IsMovablePieceTo(currentPlayer Player, curPos, distPos *Position) error {
	piece, exist := b.FindPiece(curPos)
	if !exist {
		return errors.Errorf("piece doesn't exist at %v", curPos)
	}
	if !IsSamePlayer(currentPlayer, piece.Owner()) {
		return errors.Errorf("piece doesn't belong to %s", piece.Owner().Name())
	}
	if !piece.IsMovableTo(curPos, distPos) {
		return errors.Errorf("the piece can't be moved to %v", distPos)
	}

	distPositionPiece, distPositionPieceExist := b.FindPiece(distPos)
	if distPositionPieceExist && IsSamePlayer(currentPlayer, distPositionPiece.Owner()) {
		return errors.Errorf("there is current user's piece at %v", distPos)
	}
	return nil
}

func reversePieceOrder(s []Piece) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
