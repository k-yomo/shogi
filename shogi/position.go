package shogi

import "fmt"


type Axis int

func (a Axis) Idx() int {
	return int(a - 1)
}

type Position struct {
	X Axis
	Y Axis
}
type PositionList []*Position

func (p *Position) String() string {
	return fmt.Sprintf("x: %d, y: %d", p.X, p.Y)
}

func (pl PositionList) Contains(pos *Position) bool {
	for _, p := range pl {
		if p.X == pos.X && p.Y == pos.Y {
			return true
		}
	}
	return false
}

// TODO: those two functions below should be in board.go since the size of board is about board domain.
func (p *Position) IsValid() bool {
	return p.X >= 1 && p.X <= 9 && p.Y >= 1 && p.Y <= 9
}

func (pl PositionList) SelectValid() PositionList {
	var validRelativePositions PositionList
	for _, pos := range pl {
		if pos.IsValid() {
			validRelativePositions = append(validRelativePositions, pos)
		}
	}
	return validRelativePositions
}

