package rt

import "math"

var noIntersection Intersection

type Intersection struct {
	ts    []value
	shape Shape
}

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

func NewIntersection(ts ...value) Intersection {
	return Intersection{
		ts: ts,
	}
}

func (i Intersection) Shape() Shape {
	return i.shape
}

func (i *Intersection) SetShape(shape Shape) {
	i.shape = shape
}

func (i Intersection) Get() []value {
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
	if i.shape == nil {
		return o.shape == nil
	}
	return i.shape.Equal(o.shape)
}
