package shogi

type Piece interface {
	Name() string
	ShortName() string
	Owner() Player
	TakenBy(player Player)
	IsPromotable() bool
	IsPromoted() bool
	YDirectionNum() int
	// MovablePositions returns positions where the piece can move to if there is no obstacles on the way.
	// It means, depending on the other piece positions, actual movable positions can be more limited.
	MovablePositions(curPos *Position) PositionList
	IsMovableTo(curPos, distPos *Position) bool
	// PositionsOnTheWayTo returns positions where the piece passes on the way to the destination position.
	// Return array doesn't include curPos and distPos
	PositionsOnTheWayTo(curPos, distPos *Position) (positionsOnTheWay PositionList, movableToThePosition bool)
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
	return calcMovableAbsPositions(k, curPos, KingMovableRelativePositions())
}

func (k *King) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := k.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
}

func (k *King) PositionsOnTheWayTo(curPos, distPos *Position) (positionsOnTheWay PositionList, movableToThePosition bool) {
	return nil, k.IsMovableTo(curPos, distPos)
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
	var movableRelativePositions PositionList
	if r.isPromoted {
		movableRelativePositions = PromotedRookMovableRelativePositions()
	} else {
		movableRelativePositions = RookMovableRelativePositions()
	}
	return calcMovableAbsPositions(r, curPos, movableRelativePositions)
}

func (r *Rook) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := r.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
}

func (r *Rook) PositionsOnTheWayTo(curPos, distPos *Position) (positionsOnTheWay PositionList, movableToThePosition bool) {
	if !r.IsMovableTo(curPos, distPos) {
		return nil, false
	}

	switch {
	case curPos.X < distPos.X:
		for x := curPos.X + 1; x < distPos.X; x++ {
			positionsOnTheWay = append(positionsOnTheWay, &Position{X: x, Y: curPos.X})
		}
	case curPos.X > distPos.X:
		for x := curPos.X - 1; x < distPos.X; x-- {
			positionsOnTheWay = append(positionsOnTheWay, &Position{X: x, Y: curPos.X})
		}
	case curPos.Y < distPos.Y:
		for y := curPos.Y + 1; y < distPos.Y; y++ {
			positionsOnTheWay = append(positionsOnTheWay, &Position{X: curPos.X, Y: y})
		}
	case curPos.Y > distPos.Y:
		for y := curPos.Y - 1; y > distPos.Y; y-- {
			positionsOnTheWay = append(positionsOnTheWay, &Position{X: curPos.X, Y: y})
		}
	}
	return positionsOnTheWay, true
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
	var movableRelativePositions PositionList
	if b.isPromoted {
		movableRelativePositions = PromotedBishopMovableRelativePositions()
	} else {
		movableRelativePositions = BishopMovableRelativePositions()
	}
	return calcMovableAbsPositions(b, curPos, movableRelativePositions)
}

func (b *Bishop) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := b.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
}

func (b *Bishop) PositionsOnTheWayTo(curPos, distPos *Position) (positionsOnTheWay PositionList, movableToThePosition bool) {
	if !b.IsMovableTo(curPos, distPos) {
		return nil, false
	}
	switch {
	// to top right
	case curPos.X < distPos.X && curPos.Y < distPos.Y:
		for i := 1; Axis(i) < distPos.X - curPos.X; i++ {
			positionsOnTheWay = append(positionsOnTheWay, &Position{X: curPos.X + Axis(i), Y: curPos.Y + Axis(i)})
		}
	// to bottom right
	case curPos.X < distPos.X && curPos.Y > distPos.Y:
		for i := 1; Axis(i) < distPos.X - curPos.X; i++ {
			positionsOnTheWay = append(positionsOnTheWay, &Position{X: curPos.X + Axis(i), Y: curPos.Y + Axis(-i)})
		}
	// to top left
	case curPos.X > distPos.X && curPos.Y < distPos.Y:
		for i := 1; Axis(i) < curPos.X - distPos.X; i++ {
			positionsOnTheWay = append(positionsOnTheWay, &Position{X: curPos.X + Axis(-i), Y: curPos.Y + Axis(i)})
		}
	// to bottom left
	case curPos.X > distPos.X && curPos.Y > distPos.Y:
		for i := 1; Axis(i) < curPos.X - distPos.X; i++ {
			positionsOnTheWay = append(positionsOnTheWay, &Position{X: curPos.X + Axis(-i), Y: curPos.Y + Axis(-i)})
		}
	}
	return positionsOnTheWay, true
}

// 金
type Gold struct {
	*pieceImpl
}

func NewGold(p Player) *Gold {
	return &Gold{pieceImpl: &pieceImpl{Player: p}}
}

func (g *Gold) Name() string {
	return g.ShortName()
}

func (g *Gold) ShortName() string {
	return "金"
}

func (g *Gold) MovablePositions(curPos *Position) PositionList {
	return calcMovableAbsPositions(g, curPos, GoldMovableRelativePositions())
}

func (g *Gold) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := g.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
}

func (g *Gold) PositionsOnTheWayTo(curPos, distPos *Position) (positionsOnTheWay PositionList, movableToThePosition bool) {
	return nil, g.IsMovableTo(curPos, distPos)
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
	var movableRelativePositions PositionList
	if s.isPromoted {
		movableRelativePositions = GoldMovableRelativePositions()
	} else {
		movableRelativePositions = SilverMovableRelativePositions()
	}
	return calcMovableAbsPositions(s, curPos, movableRelativePositions)
}

func (s *Silver) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := s.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
}

func (s *Silver) PositionsOnTheWayTo(curPos, distPos *Position) (positionsOnTheWay PositionList, movableToThePosition bool) {
	return nil, s.IsMovableTo(curPos, distPos)
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
	var movableRelativePositions PositionList
	if k.isPromoted {
		movableRelativePositions = GoldMovableRelativePositions()
	} else {
		movableRelativePositions = KnightMovableRelativePositions()
	}
	return calcMovableAbsPositions(k, curPos, movableRelativePositions)
}

func (k *Knight) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := k.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
}

func (k *Knight) PositionsOnTheWayTo(curPos, distPos *Position) (positionsOnTheWay PositionList, movableToThePosition bool) {
	return nil, k.IsMovableTo(curPos, distPos)
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
	var movableRelativePositions PositionList
	if l.isPromoted {
		movableRelativePositions = GoldMovableRelativePositions()
	} else {
		movableRelativePositions = LanceMovableRelativePositions()
	}
	return calcMovableAbsPositions(l, curPos, movableRelativePositions)
}

func (l *Lance) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := l.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
}

func (l *Lance) PositionsOnTheWayTo(curPos, distPos *Position) (positionsOnTheWay PositionList, movableToThePosition bool) {
	if !l.IsMovableTo(curPos, distPos) {
		return nil, l.IsMovableTo(curPos, distPos)
	}

	if curPos.Y < distPos.Y {
		for y := curPos.Y + 1; y < distPos.Y; y++ {
			positionsOnTheWay = append(positionsOnTheWay, &Position{X: curPos.X, Y: y})
		}
	} else {
		for y := curPos.Y - 1; y > distPos.Y; y-- {
			positionsOnTheWay = append(positionsOnTheWay, &Position{X: curPos.X, Y: y})
		}
	}
	return positionsOnTheWay, true
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
	var movableRelativePositions PositionList
	if p.isPromoted {
		movableRelativePositions = GoldMovableRelativePositions()
	} else {
		movableRelativePositions = PawnMovableRelativePositions()
	}
	return calcMovableAbsPositions(p, curPos, movableRelativePositions)
}

func (p *Pawn) IsMovableTo(curPos, distPos *Position) bool {
	movablePositions := p.MovablePositions(curPos)
	return movablePositions.Contains(distPos)
}

func (p *Pawn) PositionsOnTheWayTo(curPos, distPos *Position) (positionsOnTheWay PositionList, movableToThePosition bool) {
	return nil, p.IsMovableTo(curPos, distPos)
}

// calcMovableAbsPositions calculates absolute positions where the piece can be moved to
func calcMovableAbsPositions(p Piece, curPos *Position, movableRelativePositions PositionList) PositionList {
		movablePositions := PositionList{}
		for _, relativePos := range movableRelativePositions {
		pos := &Position{X: curPos.X + relativePos.X, Y: curPos.Y + relativePos.Y*Axis(p.YDirectionNum())}
		movablePositions = append(movablePositions, pos)
	}
		return movablePositions
}
