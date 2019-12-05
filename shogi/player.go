package shogi

type Player interface {
	IsFirstPlayer() bool
	Name() string
}

type PlayerImpl struct {
	takenPieces   []Piece
	isFirstPlayer bool
}

func NewPlayer(isFirstPlayer bool) *PlayerImpl {
	return &PlayerImpl{isFirstPlayer: isFirstPlayer}
}

func IsSamePlayer(p1, p2 Player) bool {
	return p1.IsFirstPlayer() == p2.IsFirstPlayer()
}

func (p *PlayerImpl) IsFirstPlayer() bool {
	return p.isFirstPlayer
}

func (p *PlayerImpl) Name() string {
	if p.isFirstPlayer {
		return "後手"
	} else {
		return "先手"
	}
}
