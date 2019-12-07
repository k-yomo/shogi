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
	if !g.board.MovePiece(g.currentPlayer, curPos, nextPos) {
		return errors.New(fmt.Sprintf("piece can't be moved or not exist"))
	}
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

