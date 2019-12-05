package shogi

import "fmt"

type Piece interface {
	Name() string
	ShortName() string
	Owner() Player
	IsPromotable() bool
	IsPromoted() bool
	MoveablePositions(curPos *Position, board Board) []*Position
	IsMoveableTo(curPos, nextPos *Position, board Board) bool
}

type pieceImpl struct {
	Player       Player
	isPromotable bool
	isPromoted   bool
}

func (p *pieceImpl) Owner() Player {
	return p.Player
}

func (p *pieceImpl) IsPromotable() bool {
	return p.isPromotable
}
func (p *pieceImpl) IsPromoted() bool {
	return p.isPromoted
}

type Position struct {
	X uint
	Y uint
}

func (p *Position) String() string {
	return fmt.Sprintf("x: %d, y: %d", p.X, p.Y)
}

// 王
type King struct {
	*pieceImpl
}

func NewKing(p Player) *King {
	return &King{pieceImpl: &pieceImpl{Player: p}}
}

func (k *King) Name() string {
	return k.ShortName()
}

func (k *King) ShortName() string {
	if k.Player.IsFirstPlayer() {
		return "玉"
	} else {
		return "王"
	}
}

func (k *King) MoveablePositions(curPos *Position, board Board) []*Position {
	return KingMoveablePositions(curPos, board, k.Player.IsFirstPlayer())
}

func (k *King) IsMoveableTo(curPos, nextPos *Position, board Board) bool {
	return IsKingMoveableTo(curPos, nextPos, board, k.Player.IsFirstPlayer())
}

// 飛車
type Rook struct {
	*pieceImpl
}

func NewRook(p Player) *Rook {
	return &Rook{pieceImpl: &pieceImpl{Player: p, isPromotable: true, isPromoted: false}}
}

func (r *Rook) Name() string {
	return "飛車"
}

func (r *Rook) ShortName() string {
	return "飛"
}

func (r *Rook) MoveablePositions(curPos *Position, board Board) []*Position {
	if r.isPromoted {
		return PromotedRookMoveablePositions(curPos, board, r.Player.IsFirstPlayer())
	} else {
		return RookMoveablePositions(curPos, board, r.Player.IsFirstPlayer())
	}
}

func (r *Rook) IsMoveableTo(curPos, nextPos *Position, board Board) bool {
	if r.isPromoted {
		return IsPromotedRookMoveableTo(curPos, nextPos, board, r.Player.IsFirstPlayer())
	} else {
		return IsRookMoveableTo(curPos, nextPos, board, r.Player.IsFirstPlayer())
	}
}

// 角
type Bishop struct {
	*pieceImpl
}

func NewBishop(p Player) *Bishop {
	return &Bishop{pieceImpl: &pieceImpl{Player: p, isPromotable: true, isPromoted: false}}
}

func (b *Bishop) Name() string {
	return b.ShortName()
}

func (b *Bishop) ShortName() string {
	return "角"
}

func (b *Bishop) MoveablePositions(curPos *Position, board Board) []*Position {
	if b.isPromoted {
		return PromotedBishopMoveablePositions(curPos, board, b.Player.IsFirstPlayer())
	} else {
		return BishopMoveablePositions(curPos, board, b.Player.IsFirstPlayer())
	}
}

func (b *Bishop) IsMoveableTo(curPos, nextPos *Position, board Board) bool {
	if b.isPromoted {
		return IsPromotedBishopMoveableTo(curPos, nextPos, board, b.Player.IsFirstPlayer())
	} else {
		return IsBishopMoveableTo(curPos, nextPos, board, b.Player.IsFirstPlayer())
	}
}

// 金
type Gold struct {
	*pieceImpl
}

func NewGold(p Player) *Bishop {
	return &Bishop{pieceImpl: &pieceImpl{Player: p}}
}

func (g *Gold) Name() string {
	return g.ShortName()
}

func (g *Gold) ShortName() string {
	return "金"
}

func (g *Gold) MoveablePositions(curPos *Position, board Board) []*Position {
	return GoldMoveablePositions(curPos, board, g.Player.IsFirstPlayer())
}

func (g *Gold) IsMoveableTo(curPos, nextPos *Position, board Board) bool {
	return IsGoldMoveableTo(curPos, nextPos, board, g.Player.IsFirstPlayer())
}

// 銀
type Silver struct {
	*pieceImpl
}

func NewSilver(p Player) *Silver {
	return &Silver{pieceImpl: &pieceImpl{Player: p, isPromotable: true, isPromoted: false}}
}

func (s *Silver) Name() string {
	return s.ShortName()
}

func (s *Silver) ShortName() string {
	return "銀"
}

func (s *Silver) MoveablePositions(curPos *Position, board Board) []*Position {
	if s.isPromoted {
		return GoldMoveablePositions(curPos, board, s.Player.IsFirstPlayer())
	}
	return nil
}

func (s *Silver) IsMoveableTo(curPos, nextPos *Position, board Board) bool {
	if s.isPromoted {
		return IsGoldMoveableTo(curPos, nextPos, board, s.Player.IsFirstPlayer())
	}
	return false
}

// 桂馬
type Knight struct {
	*pieceImpl
}

func NewKnight(p Player) *Knight {
	return &Knight{pieceImpl: &pieceImpl{Player: p, isPromotable: true, isPromoted: false}}
}

func (k *Knight) Name() string {
	return "桂馬"
}

func (k *Knight) ShortName() string {
	return "桂"
}

func (k *Knight) MoveablePositions(curPos *Position, board Board) []*Position {
	if k.isPromoted {
		return GoldMoveablePositions(curPos, board, k.Player.IsFirstPlayer())
	}
	return nil
}

func (k *Knight) IsMoveableTo(curPos, nextPos *Position, board Board) bool {
	if k.isPromoted {
		return IsGoldMoveableTo(curPos, nextPos, board, k.Player.IsFirstPlayer())
	}
	return false
}

// 香車
type Lance struct {
	*pieceImpl
}

func NewLance(p Player) *Lance {
	return &Lance{pieceImpl: &pieceImpl{Player: p, isPromotable: true, isPromoted: false}}
}

func (l *Lance) MoveablePositions(curPos *Position, board Board) []*Position {
	if l.isPromoted {
		return GoldMoveablePositions(curPos, board, l.Player.IsFirstPlayer())
	}
	return nil
}

func (l *Lance) IsMoveableTo(curPos, nextPos *Position, board Board) bool {
	if l.isPromoted {
		return IsGoldMoveableTo(curPos, nextPos, board, l.Player.IsFirstPlayer())
	}
	return false
}

func (l *Lance) Name() string {
	return "香車"
}

func (l *Lance) ShortName() string {
	return "香"
}

// 歩
type Pawn struct {
	*pieceImpl
}

func NewPawn(p Player) *Pawn {
	return &Pawn{pieceImpl: &pieceImpl{Player: p, isPromotable: true, isPromoted: false}}
}

func (p *Pawn) Name() string {
	return p.ShortName()
}

func (p *Pawn) ShortName() string {
	return "歩"
}

func (p *Pawn) MoveablePositions(curPos *Position, board Board) []*Position {
	if p.isPromoted {
		return GoldMoveablePositions(curPos, board, p.Player.IsFirstPlayer())
	}
	return nil
}

func (p *Pawn) IsMoveableTo(curPos, nextPos *Position, board Board) bool {
	if p.isPromoted {
		return IsGoldMoveableTo(curPos, nextPos, board, p.Player.IsFirstPlayer())
	}
	return false
}
