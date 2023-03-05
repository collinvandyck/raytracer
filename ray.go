package rt

type Ray struct {
	origin    Point
	direction Vector
}

func NewRay(origin Point, direction Vector) Ray {
	return Ray{
		origin:    origin,
		direction: direction,
	}
}

func (r Ray) Transform(m Matrix) Ray {
	return r
}

func (r Ray) Position(t Value) Point {
	d1 := r.direction.MultiplyBy(t)
	return d1.AddPoint(r.origin)
}

func (r Ray) Origin() Point {
	return r.origin
}

func (r Ray) Direction() Vector {
	return r.direction
}
