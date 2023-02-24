package raytracer

type float = float64

type tuple4 struct {
	x float
	y float
	z float
	w float
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
