package shogi

import "fmt"

type Piece interface {
	Name() string
	ShortName() string
	Owner() Player
	TakenBy(player Player)
	IsPromotable() bool
	IsPromoted() bool
	YDirectionNum() int
	MovablePositions(curPos *Position) PositionList
	IsMovableTo(curPos, distPos *Position) bool
}

type pieceImpl struct {
	Player       Player
	isPromotable bool
	isPromoted   bool
	isInHand     bool
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

func (p *pieceImpl) TakenBy(player Player) {
	p.Player = player
	p.isPromoted = false
	p.isInHand = true
}

// first player attacks from bottom to top which means -1 (descending)
// second player attacks from top to bottom which means 1 (ascending)
func (p *pieceImpl) YDirectionNum() int {
	if p.Player.IsFirstPlayer() {
		return -1
	} else {
		return 1
	}
}

type Position struct {
	X int
	Y int
}

type PositionList []*Position

func (pl PositionList) Contains(pos *Position) bool {
	for _, p := range pl {
		if p.X == pos.X && p.Y == pos.Y {
			return true
		}
	}
	return false
}

func (p *Position) String() string {
	return fmt.Sprintf("x: %d, y: %d", p.X, p.Y)
}

func (p *Position) IsValid() bool {
	return p.X >= 0 && p.X < 9 && p.Y >= 0 && p.Y < 9
}

func filterValidPositions(positions []*Position) []*Position {
	var validRelativePositions []*Position
	for _, pos := range positions {
		if pos.IsValid() {
			validRelativePositions = append(validRelativePositions, pos)
		}
	}
	return validRelativePositions
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

func (k *King) MovablePositions(curPos *Position) PositionList {
	return filterValidPositions(KingMovableRelativePositions())
}

func (k *King) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := k.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
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

func (r *Rook) MovablePositions(curPos *Position) PositionList {
	var movableRelativePositions []*Position
	if r.isPromoted {
		movableRelativePositions = PromotedRookMovableRelativePositions()
	} else {
		movableRelativePositions = RookMovableRelativePositions()
	}
	return calcMovablePositions(r, curPos, movableRelativePositions)
}

func (r *Rook) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := r.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
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

func (b *Bishop) MovablePositions(curPos *Position) PositionList {
	var movableRelativePositions []*Position
	if b.isPromoted {
		movableRelativePositions = PromotedBishopMovableRelativePositions()
	} else {
		movableRelativePositions = BishopMovableRelativePositions()
	}
	return calcMovablePositions(b, curPos, movableRelativePositions)
}

func (b *Bishop) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := b.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
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

func (g *Gold) MovablePositions(curPos *Position) PositionList {
	return calcMovablePositions(g, curPos, GoldMovableRelativePositions())
}

func (g *Gold) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := g.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
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

func (s *Silver) MovablePositions(curPos *Position) PositionList {
	var movableRelativePositions []*Position
	if s.isPromoted {
		movableRelativePositions = GoldMovableRelativePositions()
	} else {
		movableRelativePositions = SilverMovableRelativePositions()
	}
	return calcMovablePositions(s, curPos, movableRelativePositions)
}

func (s *Silver) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := s.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
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

func (k *Knight) MovablePositions(curPos *Position) PositionList {
	var movableRelativePositions []*Position
	if k.isPromoted {
		movableRelativePositions = GoldMovableRelativePositions()
	} else {
		movableRelativePositions = KnightMovableRelativePositions()
	}
	return calcMovablePositions(k, curPos, movableRelativePositions)
}

func (k *Knight) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := k.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
}

// 香車
type Lance struct {
	*pieceImpl
}

func NewLance(p Player) *Lance {
	return &Lance{pieceImpl: &pieceImpl{Player: p, isPromotable: true, isPromoted: false}}
}

func (l *Lance) Name() string {
	return "香車"
}

func (l *Lance) ShortName() string {
	return "香"
}

func (l *Lance) MovablePositions(curPos *Position) PositionList {
	var movableRelativePositions []*Position
	if l.isPromoted {
		movableRelativePositions = GoldMovableRelativePositions()
	} else {
		movableRelativePositions = LanceMovableRelativePositions()
	}
	return calcMovablePositions(l, curPos, movableRelativePositions)
}

func (l *Lance) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := l.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
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

func (p *Pawn) MovablePositions(curPos *Position) PositionList {
	var movableRelativePositions []*Position
	if p.isPromoted {
		movableRelativePositions = GoldMovableRelativePositions()
	} else {
		movableRelativePositions = PawnMovableRelativePositions()
	}
	return calcMovablePositions(p, curPos, movableRelativePositions)
}

func (p *Pawn) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := p.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
}

func calcMovablePositions(p Piece, curPos *Position, movableRelativePositions []*Position) []*Position {
	var movablePositions []*Position
	for _, relativePos := range movableRelativePositions {
		pos := &Position{X: curPos.X + relativePos.X, Y: curPos.Y + relativePos.Y*p.YDirectionNum()}
		movablePositions = append(movablePositions, pos)
	}
	return filterValidPositions(movablePositions)
}
