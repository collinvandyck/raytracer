package raytracer

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
	return floatsEqual(t.x, o.x) &&
		floatsEqual(t.y, o.y) &&
		floatsEqual(t.z, o.z) &&
		floatsEqual(t.w, o.w)
}

func tuple(x, y, z, w float) tuple4 {
	return tuple4{x, y, z, w}
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

type point tuple4

func newPoint(x, y, z float) point {
	return point(newTuple(x, y, z, 1))
}

func (p point) addVector(o vector) point {
	return point(tuple4(p).add(tuple4(o)))
}

func (p point) subtractPoint(o point) vector {
	return vector(tuple4(p).subtract(tuple4(o)))
}

func (p point) subtractVector(o vector) point {
	return point(tuple4(p).subtract(tuple4(o)))
}

type vector tuple4

func newVector(x, y, z float) vector {
	return vector(newTuple(x, y, z, 0))
}

func (v vector) addVector(o vector) vector {
	return vector(tuple4(v).add(tuple4(o)))
}

func (v vector) addPoint(o point) point {
	return point(tuple4(v).add(tuple4(o)))
}

func (v vector) subtractVector(o vector) vector {
	return vector(tuple4(v).subtract(tuple4(o)))
}

func (v vector) negate() vector {
	return vector(tuple4(v).negate())
}

func (v vector) magnitude() float {
	m1 := tuple4(v).multiply(tuple4(v))
	return math.Sqrt(m1.x + m1.y + m1.z)
}

func (v vector) normalize() vector {
	m1 := v.magnitude()
	m2 := tuple4(v).divideBy(m1)
	return newVector(m2.x, m2.y, m2.z)
}

func (v vector) dot(o vector) float {
	return tuple4(v).dot(tuple4(o))
}

func (v vector) cross(o vector) vector {
	x1 := v.y*o.z - v.z*o.y
	y1 := v.z*o.x - v.x*o.z
	z1 := v.x*o.y - v.y*o.x
	return newVector(x1, y1, z1)
}
