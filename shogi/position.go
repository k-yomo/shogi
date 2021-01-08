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

// IsSamePosition checks if receiver position is at the same position with given position.
func (p *Position) IsSamePosition(pos *Position) bool {
	return p.X == pos.X && p.Y == pos.Y
}
