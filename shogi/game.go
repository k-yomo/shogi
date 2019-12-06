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

func (g *Game) isMovable(curPos, distPos *Position) bool {
	piece, exist := g.board.FindPiece(curPos)
	if !exist {
		fmt.Println(fmt.Sprintf("piece doesn't exist at %v", curPos))
		return false
	}
	if !IsSamePlayer(g.currentPlayer, piece.Owner()) {
		fmt.Println(fmt.Sprintf("piece doesn't belong to %s", piece.Owner().Name()))
		return false
	}
	distPositionPiece, exist := g.board.FindPiece(distPos)
	if exist && IsSamePlayer(g.currentPlayer, distPositionPiece.Owner()) {
		fmt.Println(fmt.Sprintf("there is current user's piece at %v", distPos))
		return false
	}
	if !piece.IsMovableTo(curPos, distPos) {
		fmt.Println(fmt.Sprintf("the piece can't be moved to %v", distPos))
		return false
	}
	g.board[distPos.Y][distPos.X] = g.board[curPos.Y][curPos.X]
	return true
}
