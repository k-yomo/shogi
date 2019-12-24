package shogi

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
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

func (g *Game) FormatCurrentSituation() string {
	return fmt.Sprintf(`
-------------------------------
| 先手: %s
-------------------------------
%s
-------------------------------
| 後手: %s
-------------------------------
`, g.FormatFirstPlayerPiecesInHand(), g.board.String(), g.FormatSecondPlayerPiecesInHand())
}

func (g *Game) CurrentPlayerName() string {
	return g.currentPlayer.Name()
}

func (g *Game) MovePiece(curPos, nextPos *Position) error {
	if err := g.board.MovePiece(g.currentPlayer, curPos, nextPos); err != nil {
		return errors.Wrap(err, "move piece")
	}
	g.switchPlayer()
	return nil
}

func (g *Game) DropPiece(piece Piece, distPos *Position) error {
	if err := g.board.DropPiece(piece, distPos); err != nil {
		return err
	}
	if err := g.currentPlayer.RemoveDroppedPiece(piece); err != nil {
		return err
	}
	return nil
}

func (g *Game) CurrentPlayerPiecesInHand() []Piece {
	return g.currentPlayer.PiecesInHand()
}

func (g *Game) FormatFirstPlayerPiecesInHand() string {
	return g.formatPlayerPiecesInHands(g.firstPlayer)
}

func (g *Game) FormatSecondPlayerPiecesInHand() string {
	return g.formatPlayerPiecesInHands(g.secondPlayer)
}

func (g *Game) switchPlayer() {
	if g.currentPlayer.IsFirstPlayer() {
		g.currentPlayer = g.secondPlayer
	} else {
		g.currentPlayer = g.firstPlayer
	}
}

func (g *Game) formatPlayerPiecesInHands(p Player) string {
	piecesInHand := p.PiecesInHand()
	pieceNamesInHand := make([]string, len(piecesInHand))
	for i, piece := range piecesInHand {
		pieceNamesInHand[i] = piece.ShortName()
	}
	var pieceNamesInHandStr string
	if len(pieceNamesInHand) > 0 {
		pieceNamesInHandStr = strings.Join(pieceNamesInHand, "、")
	} else {
		pieceNamesInHandStr = "無し"
	}
	return fmt.Sprintf("%s", pieceNamesInHandStr)
}
