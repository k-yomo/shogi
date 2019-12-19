package shogi

type Player interface {
	IsFirstPlayer() bool
	Name() string
	TakePiece(piece Piece)
	PiecesInHand() []Piece
}

type PlayerImpl struct {
	piecesInHand  []Piece
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
		return "先手"
	} else {
		return "後手"
	}
}

func (p *PlayerImpl) TakePiece(piece Piece) {
	p.piecesInHand = append(p.piecesInHand, piece)
	piece.TakenBy(p)
}

func (p *PlayerImpl) PiecesInHand() []Piece {
	return p.piecesInHand
}
