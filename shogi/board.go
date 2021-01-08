package shogi

import (
	"github.com/pkg/errors"
)

type Board [][]Piece

var rows = []Axis{1, 2, 3, 4, 5, 6, 7, 8, 9}
var columns = []Axis{1, 2, 3, 4, 5, 6, 7, 8, 9}

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

func reversePieceOrder(s []Piece) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// initializeBase initializes middle empty 3 rows
func initializeMiddleZone() [][]Piece {
	rows := make([][]Piece, 3)
	for i := range rows {
		rows[i] = []Piece{nil, nil, nil, nil, nil, nil, nil, nil, nil}
	}
	return rows
}

func (b Board) opponentArea(player Player) [][]Piece {
	if player.IsFirstPlayer() {
		return b[0:4]
	} else {
		return b[6:9]
	}
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

func (b Board) FindPlayerPiece(curPlayer Player, pos *Position) (p Piece, exist bool) {
	p, exist = b.FindPiece(pos)
	if !exist {
		return nil, false
	}
	return p, IsSamePlayer(curPlayer, p.Owner())
}

func (b Board) MovePiece(curPlayer Player, curPos, distPos *Position) error {
	piece, exist := b.FindPlayerPiece(curPlayer, curPos)
	if !exist {
		return errors.Errorf("current player's piece doesn't exist at %v", curPos)
	}

	if !b.IsPieceMovableTo(curPos, distPos) {
		return errors.Errorf("the piece can't be moved to %v", distPos)
	}

	pieceAtDistPos, pieceExistsAtDistPos := b.FindPiece(distPos)
	if pieceExistsAtDistPos && IsSamePlayer(curPlayer, pieceAtDistPos.Owner()) {
		return errors.Errorf("there is current user's piece at %v", distPos)
	}

	_, isMovingKing := piece.(*King)
	if b.validateChecking(curPlayer, distPos, isMovingKing) {
		return errors.New("king is checked, so you must avoid it")
	}

	if pieceExistsAtDistPos {
		curPlayer.TakePiece(pieceAtDistPos)
	}
	b[distPos.Y.Idx()][distPos.X.Idx()] = b[curPos.Y.Idx()][curPos.X.Idx()]
	b[curPos.Y.Idx()][curPos.X.Idx()] = nil
	return nil
}

func (b Board) IsPieceMovableTo(curPos, distPos *Position) bool {
	piece, exist := b.FindPiece(curPos)
	if !exist {
		return false
	}
	positionsOnTheWay, isMovable := piece.PositionsOnTheWayTo(curPos, distPos)
	if !isMovable {
		return false
	}
	for _, pos := range positionsOnTheWay {
		if _, exist := b.FindPiece(pos); exist {
			return false
		}
	}
	return true
}

func (b Board) PieceMovablePosition(piecePos *Position) PositionList {
	piece, _ := b.FindPiece(piecePos)
	movablePositions := piece.MovablePositions(piecePos)
	var actualMovablePositions PositionList
	for _, distPos := range movablePositions {
		if b.IsPieceMovableTo(piecePos, distPos) {
			actualMovablePositions = append(actualMovablePositions, distPos)
		}
	}
	return actualMovablePositions
}

func (b Board) DropPiece(piece Piece, distPos *Position) error {
	_, pieceExistsAtDistPos := b.FindPiece(distPos)
	if pieceExistsAtDistPos {
		return errors.Errorf("there is a piece at %v", distPos)
	}
	if pawn, isPawn := piece.(*Pawn); isPawn {
		if b.isTwoPawnsOnSameColumn(pawn, distPos.X) {
			return errors.Errorf("there is a pawn on the same column: %v", distPos.X)
		}
	}
	if b.validateChecking(piece.Owner(), distPos, false) {
		return errors.New("king is checked, so you must avoid it")
	}
	b[distPos.X.Idx()][distPos.Y.Idx()] = piece
	return nil
}

func (b Board) validateChecking(curPlayer Player, distPos *Position, isMovingKing bool) bool {
	checkingPiecePos := b.checkingPiecePosition(curPlayer)
	if checkingPiecePos == nil {
		return true
	}
	if isMovingKing {
		if !b.PieceMovablePosition(checkingPiecePos).Contains(distPos) {
			return true
		}
	} else {
		p, _ := b.FindPiece(checkingPiecePos)
		positionsOnTheWayToKing, _ := p.PositionsOnTheWayTo(checkingPiecePos, b.findPlayerKingPosition(curPlayer))
		if positionsOnTheWayToKing.Contains(distPos) {
			return true
		}
	}
	return false
}

// CheckingPiecePositionOnTheWayToKing returns the checking piece's positions on the of the way to given player's king
// if king is not checked, it returns nil
func (b Board) checkingPiecePosition(player Player) *Position {
	kingPos := b.findPlayerKingPosition(player)
	for _, pos := range b.findOpponentPiecePositions(player) {
		if b.IsPieceMovableTo(pos, kingPos) {
			return pos
		}
	}
	return nil
}

func (b Board) findOpponentPiecePositions(curPlayer Player) PositionList {
	var opponentPiecePositions PositionList
	b.iterateThrough(func(pos *Position, piece Piece, exist bool) (finished bool) {
		if exist && !IsSamePlayer(curPlayer, piece.Owner()) {
			opponentPiecePositions = append(opponentPiecePositions, pos)
		}
		return false
	})
	return opponentPiecePositions
}

func (b Board) findPlayerKingPosition(curPlayer Player) *Position {
	var kingPos *Position
	b.iterateThrough(func(pos *Position, piece Piece, exist bool) (finished bool) {
		if king, ok := piece.(*King); ok && IsSamePlayer(curPlayer, king.Player) {
			kingPos = pos
			return true
		}
		return false
	})
	if kingPos != nil {
		return kingPos
	}
	panic("king is not found, something is wrong")
}

func (b Board) iterateThrough(inner func(pos *Position, piece Piece, exist bool) (finished bool)) {
	for _, y := range rows {
		for _, x := range columns {
			pos := &Position{X: x, Y: y}
			p, exist := b.FindPiece(pos)
			if finished := inner(pos, p, exist); finished {
				return
			}
		}
	}
}

func (b Board) isTwoPawnsOnSameColumn(pawn *Pawn, distPosX Axis) bool {
	for _, y := range rows {
		p, exist := b.FindPiece(&Position{X: distPosX, Y: y})
		twoPawnsOnSameColumn := exist && IsSamePlayer(p.Owner(), pawn.Owner()) && !p.IsPromoted()
		if twoPawnsOnSameColumn {
			return false
		}
	}
	return true
}
