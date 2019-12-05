package shogi

import (
	"errors"
	"fmt"
)

type Game struct {
	currentPlayer Player
	firstPlayer   Player
	secondPlayer  Player
	board         Board
}

// NewGame starts new shogi game
func NewGame() *Game {
	firstPlayer := NewPlayer(true)
	secondPlayer := NewPlayer(false)
	board := NewBoard(firstPlayer, secondPlayer)
	return &Game{
		currentPlayer: firstPlayer,
		firstPlayer:   firstPlayer,
		secondPlayer:  secondPlayer,
		board:         board,
	}
}

func (g *Game) Print() {
	fmt.Println(g.board.String())
}

func (g *Game) CurrentPlayerName() string {
	return g.currentPlayer.Name()
}

func (g *Game) MovePiece(curPos, nextPos *Position) error {
	piece, exist := g.board.FindPiece(curPos)
	if !exist {
		return errors.New(fmt.Sprintf("piece is not found at %v", curPos))
	}
	if !IsSamePlayer(g.currentPlayer, piece.Owner()) {
		return errors.New("piece does not belong to current player")
	}
	nextPositionPiece, exist := g.board.FindPiece(nextPos)
	if exist && IsSamePlayer(g.currentPlayer, nextPositionPiece.Owner()) {
		return errors.New(fmt.Sprintf("current user's piece exists at %v", nextPos))
	}

	if !piece.IsMoveableTo(curPos, nextPos, g.board) {
		return errors.New(fmt.Sprintf("%s can't be moved from %v to %v", piece.Name(), curPos, nextPos))
	}

	fmt.Println(g.board.String())
	g.switchPlayer()
	return nil
}

func (g *Game) switchPlayer() {
	if g.currentPlayer.IsFirstPlayer() {
		g.currentPlayer = g.secondPlayer
	} else {
		g.currentPlayer = g.firstPlayer
	}
}
