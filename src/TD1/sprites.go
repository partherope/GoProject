package td1

//this part is to realize the structure of the sprite
import (
	"math"
)

type Point2D struct {
	x, y float64
}

type Rectangle struct {
	//the rectangle is defined by two points
	//the first point is the top left point
	//the second point is the bottom right point
	p1, p2 *Point2D
}

type Sprite struct {
	//the sprite contains a position, a hitbox, a factor of zoom and the name of the file bitmap
	position *Point2D
	hitbox   *Rectangle
	zoom     float64
	file     string
}

// construction function of Point2D
func NewPoint2D(x, y float64) *Point2D {
	return &Point2D{x, y}
}

func (p Point2D) X() float64 {
	return p.x
}

func (p *Point2D) SetX(x float64) {
	p.x = x
}

func (p Point2D) Y() float64 {
	return p.y
}

func (p *Point2D) SetY(y float64) {
	p.y = y
}

func (p Point2D) Clone() *Point2D {
	//deep copy
	return &Point2D{p.x, p.y}
}

func (p Point2D) Module() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func NewRectangle(p1, p2 *Point2D) *Rectangle {
	return &Rectangle{p1, p2}
}

func (r Rectangle) P1() *Point2D {
	return r.p1
}

func (r *Rectangle) SetP1(p1 *Point2D) {
	r.p1 = p1
}

func (r Rectangle) P2() *Point2D {
	return r.p2
}

func (r *Rectangle) SetP2(p2 *Point2D) {
	r.p2 = p2
}

func (r Rectangle) Clone() *Rectangle {
	//deep copy
	return &Rectangle{r.p1.Clone(), r.p2.Clone()}
}

// construction function of Sprite
func NewSprite(hitbox *Rectangle, zoom float64, file string) *Sprite {
	//the position is the center of the hitbox
	center := NewPoint2D((hitbox.P1().X()+hitbox.P2().X())/2, (hitbox.P1().Y()+hitbox.P2().Y())/2)
	return &Sprite{center, hitbox, zoom, file}
}

func (s *Sprite) Move(dx, dy float64) {
	//move the sprite by dx and dy
	s.position.SetX(s.position.X() + dx)
	s.position.SetY(s.position.Y() + dy)
}

func (s Sprite) Collision(s2 Sprite) *Rectangle {
	sLeft := s.position.X() - s.hitbox.P1().X()
	sRight := s.position.X() + s.hitbox.P2().X()
	sTop := s.position.Y() - s.hitbox.P1().Y()
	sBottom := s.position.Y() + s.hitbox.P2().Y()

	s2Left := s2.position.X() - s2.hitbox.P1().X()
	s2Right := s2.position.X() + s2.hitbox.P2().X()
	s2Top := s2.position.Y() - s2.hitbox.P1().Y()
	s2Bottom := s2.position.Y() + s2.hitbox.P2().Y()

	if sLeft > s2Right || sRight < s2Left || sTop > s2Bottom || sBottom < s2Top {
		return nil
	}

	x1 := max(sLeft, s2Left)
	x2 := min(sRight, s2Right)
	y1 := max(sTop, s2Top)
	y2 := min(sBottom, s2Bottom)

	return NewRectangle(NewPoint2D(x1, y1), NewPoint2D(x2, y2))
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
