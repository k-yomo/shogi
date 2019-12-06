package main

import (
	"fmt"
	"github.com/k-yomo/shogi/shogi"
)

func main() {
	game := shogi.NewGame()
	game.Print()
	if err := game.MovePiece(&shogi.Position{X: 1, Y: 6}, &shogi.Position{X: 1, Y: 5}); err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("moved!")
	}
	if err := game.MovePiece(&shogi.Position{X: 6, Y: 2}, &shogi.Position{X: 6, Y: 5}); err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("moved!")
	}
	game.Print()
}
