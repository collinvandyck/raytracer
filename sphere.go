package rt

import "math"

func IntersectSphere(sphere Sphere, ray Ray) []Value {
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
	return []Value{t1, t2}
}

type Sphere struct {
}

// todo: do i need to force an allocation here?
func NewSphere() Sphere {
	return Sphere{}
}

func (s Sphere) Equal(o Sphere) bool {
	return true
}

func (s Sphere) Point() Point {
	return NewPoint(0, 0, 0)
}

func (s Sphere) String() string {
	return "Sphere()"
}
