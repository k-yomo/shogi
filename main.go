package main

import (
	"fmt"
	"github.com/k-yomo/shogi/shogi"
)

func main() {
	game := shogi.NewGame()
	game.Print()
	if err := game.MovePiece(&shogi.Position{X: 1, Y: 3}, &shogi.Position{X: 1, Y: 4}); err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("moved!")
	}
}
