package rt

import (
	"math"
)

type (
	float = float64
)

type tuple4 struct {
	x float
	y float
	z float
	w float
}

func (t tuple4) add(o tuple4) tuple4 {
	return tuple4{
		t.x + o.x,
		t.y + o.y,
		t.z + o.z,
		t.w + o.w,
	}
}

func (t tuple4) subtract(o tuple4) tuple4 {
	return tuple4{
		t.x - o.x,
		t.y - o.y,
		t.z - o.z,
		t.w - o.w,
	}
}

func (t tuple4) multiplyBy(val float) tuple4 {
	return tuple4{
		t.x * val,
		t.y * val,
		t.z * val,
		t.w * val,
	}
}

func (t tuple4) multiply(o tuple4) tuple4 {
	return tuple4{
		t.x * o.x,
		t.y * o.y,
		t.z * o.z,
		t.w * o.w,
	}
}

func (t tuple4) divideBy(val float) tuple4 {
	return tuple4{
		t.x / val,
		t.y / val,
		t.z / val,
		t.w / val,
	}
}

func (t tuple4) divide(o tuple4) tuple4 {
	return tuple4{
		t.x / o.x,
		t.y / o.y,
		t.z / o.z,
		t.w / o.w,
	}
}

func (t tuple4) dot(o tuple4) float {
	x2 := t.x * o.x
	y2 := t.y * o.y
	z2 := t.z * o.z
	w2 := t.w * o.w
	return x2 + y2 + z2 + w2
}

func (t tuple4) negate() tuple4 {
	var zero tuple4
	return zero.subtract(t)
}

func (t tuple4) equal(o tuple4) bool {
	xe := floatsEqual(t.x, o.x)
	ye := floatsEqual(t.y, o.y)
	ze := floatsEqual(t.z, o.z)
	we := floatsEqual(t.w, o.w)
	return xe && ye && ze && we
}

func (t tuple4) isPoint() bool {
	return t.w == 1
}

func (t tuple4) isVector() bool {
	return t.w == 0
}
func newTuple(x, y, z, w float) tuple4 {
	return tuple4{x, y, z, w}
}

type Point tuple4

func NewPoint(x, y, z float) Point {
	return Point(newTuple(x, y, z, 1))
}

func (p Point) Scale(o Point) Point {
	return Point(tuple4(p).multiply(tuple4(o)))
}

func (p Point) AddVector(o Vector) Point {
	return Point(tuple4(p).add(tuple4(o)))
}

func (p Point) SubtractPoint(o Point) Vector {
	return Vector(tuple4(p).subtract(tuple4(o)))
}

func (p Point) SubtractVector(o Vector) Point {
	return Point(tuple4(p).subtract(tuple4(o)))
}

func (p Point) Equal(o Point) bool {
	return tuple4(p).equal(tuple4(o))
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

type Vector tuple4

func NewVector(x, y, z float) Vector {
	return Vector(newTuple(x, y, z, 0))
}

func (v Vector) AddVector(o Vector) Vector {
	return Vector(tuple4(v).add(tuple4(o)))
}

func (v Vector) AddPoint(o Point) Point {
	return Point(tuple4(v).add(tuple4(o)))
}

func (v Vector) SubtractVector(o Vector) Vector {
	return Vector(tuple4(v).subtract(tuple4(o)))
}

func (v Vector) Negate() Vector {
	return Vector(tuple4(v).negate())
}

func (v Vector) MultiplyBy(o float) Vector {
	return Vector(tuple4(v).multiplyBy(o))
}

func (v Vector) Magnitude() float {
	m1 := tuple4(v).multiply(tuple4(v))
	return math.Sqrt(m1.x + m1.y + m1.z)
}

func (v Vector) Normalize() Vector {
	m1 := v.Magnitude()
	d1 := tuple4(v).divideBy(m1)
	return NewVector(d1.x, d1.y, d1.z)
}

func (v Vector) Dot(o Vector) float {
	return tuple4(v).dot(tuple4(o))
}

func (v Vector) Cross(o Vector) Vector {
	x1 := v.y*o.z - v.z*o.y
	y1 := v.z*o.x - v.x*o.z
	z1 := v.x*o.y - v.y*o.x
	return NewVector(x1, y1, z1)
}

func (v Vector) Equal(o Vector) bool {
	return tuple4(v).equal(tuple4(o))
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
