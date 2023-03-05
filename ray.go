package rt

import "math"

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

func (r Ray) IntersectSphere(sphere Sphere) (res Intersection) {
	// the vector from the sphere's center to the ray origin
	sphereToRay := r.Origin().SubtractPoint(sphere.Point())

	a := r.Direction().Dot(r.Direction())
	b := 2 * r.Direction().Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		// there was no intersection
		return res
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	return sphere.Intersection(t1, t2)
}

func (r Ray) Position(t value) Point {
	d1 := r.direction.MultiplyBy(t)
	return d1.AddPoint(r.origin)
}

func (r Ray) Origin() Point {
	return r.origin
}

func (r Ray) Direction() Vector {
	return r.direction
}
