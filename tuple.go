package raytracer

import "math"

type (
	float = float64
)

type tuple4 struct {
	x float
	y float
	z float
	w float
}

func newTuple(x, y, z, w float) tuple4 {
	return tuple4{x, y, z, w}
}

type point tuple4

func newPoint(x, y, z float) point {
	return point(newTuple(x, y, z, 1))
}

func (p point) addVector(v2 vector) point {
	return point(tuple4(p).add(tuple4(v2)))
}

func (p point) subtractPoint(p2 point) vector {
	return vector(tuple4(p).subtract(tuple4(p2)))
}

func (p point) subtractVector(v2 vector) point {
	return point(tuple4(p).subtract(tuple4(v2)))
}

type vector tuple4

func newVector(x, y, z float) vector {
	return vector(newTuple(x, y, z, 0))
}

func (v vector) addVector(v2 vector) vector {
	return vector(tuple4(v).add(tuple4(v2)))
}

func (v vector) addPoint(p2 point) point {
	return point(tuple4(v).add(tuple4(p2)))
}

func (v vector) subtractVector(v2 vector) vector {
	return vector(tuple4(v).subtract(tuple4(v2)))
}

func (v vector) negate() vector {
	return vector(tuple4(v).negate())
}

func (v vector) magnitude() float {
	return math.Abs(v.x) + math.Abs(v.y) + math.Abs(v.z)
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

func (t tuple4) multiply(val float) tuple4 {
	return tuple4{
		t.x * val,
		t.y * val,
		t.z * val,
		t.w * val,
	}
}

func (t tuple4) divide(val float) tuple4 {
	return tuple4{
		t.x / val,
		t.y / val,
		t.z / val,
		t.w / val,
	}
}

var zeroTuple4 tuple4

func (t tuple4) negate() tuple4 {
	return zeroTuple4.subtract(t)
}

func (t tuple4) Equal(o tuple4) bool {
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
