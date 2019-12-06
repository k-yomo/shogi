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
	var moveableRelativePositions []*Position
	for i := 1; i < 9; i++ {
		moveableRelativePositions = append(moveableRelativePositions, &Position{X: 0, Y: i})  // up to top edge
		moveableRelativePositions = append(moveableRelativePositions, &Position{X: i, Y: 0})  // up to right edge
		moveableRelativePositions = append(moveableRelativePositions, &Position{X: 0, Y: -i}) // up to bottom edge
		moveableRelativePositions = append(moveableRelativePositions, &Position{X: -i, Y: 0}) // up to left edge
	}
	return moveableRelativePositions
}

func PromotedRookMovableRelativePositions() []*Position {
	aditionalMovableRelativePositions := []*Position{
		{X: -1, Y: 1}, {X: 1, Y: 1},
		{X: -1, Y: -1}, {X: 1, Y: -1},
	}
	return append(RookMovableRelativePositions(), aditionalMovableRelativePositions...)
}

func BishopMovableRelativePositions() []*Position {
	var moveableRelativePositions []*Position
	for i := 1; i < 9; i++ {
		moveableRelativePositions = append(moveableRelativePositions, &Position{X: -i, Y: i})  // up to top-left corner
		moveableRelativePositions = append(moveableRelativePositions, &Position{X: i, Y: i})   // up to top-right corner
		moveableRelativePositions = append(moveableRelativePositions, &Position{X: i, Y: -i})  // up to bottom-right corner
		moveableRelativePositions = append(moveableRelativePositions, &Position{X: -i, Y: -i}) // up to bottom-left corner
	}
	return moveableRelativePositions
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
	var moveableRelativePositions []*Position
	for i := 1; i < 9; i++ {
		moveableRelativePositions = append(moveableRelativePositions, &Position{X: 0, Y: i}) // up to top edge
	}
	return moveableRelativePositions
}

func PawnMovableRelativePositions() []*Position {
	return []*Position{{X: 0, Y: 1}}
}
