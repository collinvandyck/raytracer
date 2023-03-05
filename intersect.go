package rt

import "math"

func IntersectSphere(sphere Sphere, ray Ray) []value {
	// the vector from the sphere's center to the ray origin
	sphereToRay := ray.Origin().SubtractPoint(sphere.Point())

	a := ray.Direction().Dot(ray.Direction())
	b := 2 * ray.Direction().Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		// there was no intersection
		return nil
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	return []value{t1, t2}
}

type Intersection struct {
	t     value
	shape Shape
}

func NewIntersection(shape Shape, t value) Intersection {
	return Intersection{
		shape: shape,
		t:     t,
	}
}

func (i Intersection) Equal(o Intersection) bool {
	if !floatsEqual(i.t, o.t) {
		return false
	}
	if i.shape == nil {
		return o.shape == nil
	}
	return i.shape.Equal(o.shape)
}
