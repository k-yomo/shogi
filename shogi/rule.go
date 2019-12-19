package shogi

// Calculated RelativePositions are theoretically movable positions, and not sure if actually movable
// since it depends on the current position and surrounding pieces.
func KingMovableRelativePositions() []*Position {
	return []*Position{
		{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1},
		{X: -1, Y: 0}, {X: 1, Y: 0},
		{X: -1, Y: -1}, {X: 0, Y: - 1}, {X: 1, Y: -1},
	}
}

func RookMovableRelativePositions() []*Position {
	var movableRelativePositions []*Position
	for i := Axis(1); i < 9; i++ {
		movableRelativePositions = append(movableRelativePositions, &Position{X: 0, Y: i})  // up to top edge
		movableRelativePositions = append(movableRelativePositions, &Position{X: i, Y: 0})  // up to right edge
		movableRelativePositions = append(movableRelativePositions, &Position{X: 0, Y: -i}) // up to bottom edge
		movableRelativePositions = append(movableRelativePositions, &Position{X: -i, Y: 0}) // up to left edge
	}
	return movableRelativePositions
}

func PromotedRookMovableRelativePositions() []*Position {
	additionalMovableRelativePositions := []*Position{
		{X: -1, Y: 1}, {X: 1, Y: 1},
		{X: -1, Y: -1}, {X: 1, Y: -1},
	}
	return append(RookMovableRelativePositions(), additionalMovableRelativePositions...)
}

func BishopMovableRelativePositions() []*Position {
	var movableRelativePositions []*Position
	for i := Axis(1); i < 9; i++ {
		movableRelativePositions = append(movableRelativePositions, &Position{X: -i, Y: i})  // up to top-left corner
		movableRelativePositions = append(movableRelativePositions, &Position{X: i, Y: i})   // up to top-right corner
		movableRelativePositions = append(movableRelativePositions, &Position{X: i, Y: -i})  // up to bottom-right corner
		movableRelativePositions = append(movableRelativePositions, &Position{X: -i, Y: -i}) // up to bottom-left corner
	}
	return movableRelativePositions
}

func PromotedBishopMovableRelativePositions() []*Position {
	additionalMovableRelativePositions := []*Position{
		{X: -1, Y: 1}, {X: 1, Y: 1},
		{X: -1, Y: -1}, {X: 1, Y: -1},
	}
	return append(BishopMovableRelativePositions(), additionalMovableRelativePositions...)
}

func GoldMovableRelativePositions() []*Position {
	return []*Position{
		{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1},
		{X: -1, Y: 0}, {X: 1, Y: 0},
		{X: 0, Y: - 1},
	}
}

func SilverMovableRelativePositions() []*Position {
	return []*Position{
		{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1},
		{X: -1, Y: -1}, {X: 1, Y: -1},
	}
}

func KnightMovableRelativePositions() []*Position {
	return []*Position{
		{X: -1, Y: 2}, {X: 1, Y: 2},
	}
}

func LanceMovableRelativePositions() []*Position {
	var movableRelativePositions []*Position
	for i := Axis(1); i < 9; i++ {
		movableRelativePositions = append(movableRelativePositions, &Position{X: 0, Y: i}) // up to top edge
	}
	return movableRelativePositions
}

func PawnMovableRelativePositions() []*Position {
	return []*Position{{X: 0, Y: 1}}
}
