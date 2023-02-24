package raytracer

type float = float64

type tuple struct {
	x float
	y float
	z float
	w float
}

func point(x, y, z float) tuple {
	return tuple{x, y, z, 1}
}

func vector(x, y, z float) tuple {
	return tuple{x, y, z, 0}
}

func (t tuple) isPoint() bool {
	return t.w == 1
}

func (t tuple) isVector() bool {
	return t.w == 0
}
