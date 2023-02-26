package rt

import (
	"math"
)

type (
	float = float64
)

type Tuple4 struct {
	x float
	y float
	z float
	w float
}

func (t Tuple4) add(o Tuple4) Tuple4 {
	return Tuple4{
		t.x + o.x,
		t.y + o.y,
		t.z + o.z,
		t.w + o.w,
	}
}

func (t Tuple4) subtract(o Tuple4) Tuple4 {
	return Tuple4{
		t.x - o.x,
		t.y - o.y,
		t.z - o.z,
		t.w - o.w,
	}
}

func (t Tuple4) multiplyBy(val float) Tuple4 {
	return Tuple4{
		t.x * val,
		t.y * val,
		t.z * val,
		t.w * val,
	}
}

func (t Tuple4) multiply(o Tuple4) Tuple4 {
	return Tuple4{
		t.x * o.x,
		t.y * o.y,
		t.z * o.z,
		t.w * o.w,
	}
}

func (t Tuple4) divideBy(val float) Tuple4 {
	return Tuple4{
		t.x / val,
		t.y / val,
		t.z / val,
		t.w / val,
	}
}

func (t Tuple4) divide(o Tuple4) Tuple4 {
	return Tuple4{
		t.x / o.x,
		t.y / o.y,
		t.z / o.z,
		t.w / o.w,
	}
}

func (t Tuple4) dot(o Tuple4) float {
	x2 := t.x * o.x
	y2 := t.y * o.y
	z2 := t.z * o.z
	w2 := t.w * o.w
	return x2 + y2 + z2 + w2
}

func (t Tuple4) negate() Tuple4 {
	var zero Tuple4
	return zero.subtract(t)
}

func (t Tuple4) equal(o Tuple4) bool {
	xe := floatsEqual(t.x, o.x)
	ye := floatsEqual(t.y, o.y)
	ze := floatsEqual(t.z, o.z)
	we := floatsEqual(t.w, o.w)
	return xe && ye && ze && we
}

func (t Tuple4) isPoint() bool {
	return t.w == 1
}

func (t Tuple4) isVector() bool {
	return t.w == 0
}
func NewTuple(x, y, z, w float) Tuple4 {
	return Tuple4{x, y, z, w}
}

type Point Tuple4

func NewPoint(x, y, z float) Point {
	return Point(NewTuple(x, y, z, 1))
}

func (p Point) Scale(o Point) Point {
	return Point(Tuple4(p).multiply(Tuple4(o)))
}

func (p Point) AddVector(o Vector) Point {
	return Point(Tuple4(p).add(Tuple4(o)))
}

func (p Point) SubtractPoint(o Point) Vector {
	return Vector(Tuple4(p).subtract(Tuple4(o)))
}

func (p Point) SubtractVector(o Vector) Point {
	return Point(Tuple4(p).subtract(Tuple4(o)))
}

func (p Point) Equal(o Point) bool {
	return Tuple4(p).equal(Tuple4(o))
}

func (p Point) X() float {
	return p.x
}

func (p Point) Y() float {
	return p.y
}

func (p Point) Z() float {
	return p.z
}

type Vector Tuple4

func NewVector(x, y, z float) Vector {
	return Vector(NewTuple(x, y, z, 0))
}

func (v Vector) AddVector(o Vector) Vector {
	return Vector(Tuple4(v).add(Tuple4(o)))
}

func (v Vector) AddPoint(o Point) Point {
	return Point(Tuple4(v).add(Tuple4(o)))
}

func (v Vector) SubtractVector(o Vector) Vector {
	return Vector(Tuple4(v).subtract(Tuple4(o)))
}

func (v Vector) Negate() Vector {
	return Vector(Tuple4(v).negate())
}

func (v Vector) MultiplyBy(o float) Vector {
	return Vector(Tuple4(v).multiplyBy(o))
}

func (v Vector) Magnitude() float {
	m1 := Tuple4(v).multiply(Tuple4(v))
	return math.Sqrt(m1.x + m1.y + m1.z)
}

func (v Vector) Normalize() Vector {
	m1 := v.Magnitude()
	d1 := Tuple4(v).divideBy(m1)
	return NewVector(d1.x, d1.y, d1.z)
}

func (v Vector) Dot(o Vector) float {
	return Tuple4(v).dot(Tuple4(o))
}

func (v Vector) Cross(o Vector) Vector {
	x1 := v.y*o.z - v.z*o.y
	y1 := v.z*o.x - v.x*o.z
	z1 := v.x*o.y - v.y*o.x
	return NewVector(x1, y1, z1)
}

func (v Vector) Equal(o Vector) bool {
	return Tuple4(v).equal(Tuple4(o))
}

func (v Vector) X() float {
	return v.x
}

func (v Vector) Y() float {
	return v.y
}

func (v Vector) Z() float {
	return v.z
}
