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
	ro := m.MultiplyPoint(r.origin)
	rd := m.MultiplyVector(r.direction)
	return NewRay(ro, rd)
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

func (r *Ray) NormalizeDirection() {
	r.direction = r.direction.Normalize()
}
