package rt

type Ray struct {
	origin    Point
	direction Vector
}

type Intersection struct {
	ts []float
}

func NewIntersection(ts ...float) Intersection {
	return Intersection{
		ts: ts,
	}
}

func (i Intersection) Get() []float {
	return i.ts
}

func (i Intersection) Len() int {
	return len(i.ts)
}

func (i Intersection) Equal(o Intersection) bool {
	if i.Len() != o.Len() {
		return false
	}
	for idx := 0; idx < i.Len(); idx++ {
		if !floatsEqual(i.ts[idx], o.ts[idx]) {
			return false
		}
	}
	return true
}

func NewRay(origin Point, direction Vector) Ray {
	return Ray{
		origin:    origin,
		direction: direction,
	}
}

func (r Ray) IntersectSphere(sphere Sphere) Intersection {
	return Intersection{}
}

func (r Ray) Position(t float) Point {
	d1 := r.direction.MultiplyBy(t)
	return d1.AddPoint(r.origin)
}

func (r Ray) Origin() Point {
	return r.origin
}

func (r Ray) Direction() Vector {
	return r.direction
}
