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

func (p Point2D) clone() *Point2D {
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
