package main

import (
	"fmt"
	"github.com/k-yomo/shogi/shogi"
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
	// _ := game.MovePiece(&shogi.Position{X: 0, Y: 6}, &shogi.Position{X: 0, Y: 5})
	fmt.Println(game.FormatCurrentSituation())
}
