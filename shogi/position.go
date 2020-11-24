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

// NOTE: those two functions may be better to be in board.go since the size of board is knowledge of board.
func (p *Position) IsValid() bool {
	return p.X >= 1 && p.X <= 9 && p.Y >= 1 && p.Y <= 9
}

// IsSamePosition checks if receiver position is at the same position with given position.
func (p *Position) IsSamePosition(pos *Position) bool {
	return p.X == pos.X && p.Y == pos.Y
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
