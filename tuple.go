package raytracer

type float = float64

var zeroVector = vector(0, 0, 0)

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

func (t tuple4) negate() tuple4 {
	zero := tuple4{0, 0, 0, 0}
	return zero.subtract(t)
}

func (t tuple4) multiply(val float) tuple4 {
	return tuple4{
		t.x * val,
		t.y * val,
		t.z * val,
		t.w * val,
	}
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

func point(x, y, z float) tuple4 {
	return tuple4{x, y, z, 1}
}

func vector(x, y, z float) tuple4 {
	return tuple4{x, y, z, 0}
}

func (t tuple4) isPoint() bool {
	return t.w == 1
}

func (t tuple4) isVector() bool {
	return t.w == 0
}
