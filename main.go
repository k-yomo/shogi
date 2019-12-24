package main

import (
	"fmt"
	"github.com/k-yomo/shogi/shogi"
	"log"
)

func main() {
	game := shogi.NewGame()
	fmt.Println(game.FormatCurrentSituation())
	_ = game.MovePiece(&shogi.Position{X: 1, Y: 7}, &shogi.Position{X: 1, Y: 6})
	fmt.Println(game.FormatCurrentSituation())
	_ = game.MovePiece(&shogi.Position{X: 1, Y: 3}, &shogi.Position{X: 1, Y: 4})
	fmt.Println(game.FormatCurrentSituation())
	_ = game.MovePiece(&shogi.Position{X: 1, Y: 6}, &shogi.Position{X: 1, Y: 5})
	fmt.Println(game.FormatCurrentSituation())
	_ = game.MovePiece(&shogi.Position{X: 1, Y: 4}, &shogi.Position{X: 1, Y: 5})
	fmt.Println(game.FormatCurrentSituation())
	_ = game.MovePiece(&shogi.Position{X: 9, Y: 3}, &shogi.Position{X: 9, Y: 4})
	pieces := game.CurrentPlayerPiecesInHand()
	if err := game.DropPiece(pieces[0], &shogi.Position{X: 9, Y: 3}); err != nil {
		log.Fatal(err)
	}
	fmt.Println(game.FormatCurrentSituation())
}
