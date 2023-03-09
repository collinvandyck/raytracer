package rt

import "math"

func IntersectSphere(sphere *Sphere, ray Ray) Intersections {
	// transform the ray into object coordinates
	ray = ray.Transform(sphere.GetInverseTransform())

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

	i1 := NewIntersection(t1, sphere)
	i2 := NewIntersection(t2, sphere)
	return NewIntersections(i1, i2)
}

func NormalAtSphere(s1 *Sphere, p1 Point) Vector {
	t1 := s1.GetInverseTransform()          // get the inverse transform
	p1 = t1.MultiplyPoint(p1)               // multiply by the inverse to make p1 be in object coords
	n1 := p1.SubtractPoint(Origin)          // subtract the point from the origin to get the normal
	w1 := t1.Transpose().MultiplyVector(n1) // transpose the normal
	w1.SetW(0)                              // correct the transposition
	w1 = w1.Normalize()                     // finally normalize it
	return n1
}

type Sphere struct {
	transform Matrix
	inverse   Matrix
}

// todo: do i need to force an allocation here?
func NewSphere() *Sphere {
	return &Sphere{}
}

func (s *Sphere) NormalAt(point Point) Vector {
	return NormalAtSphere(s, point)
}

func (s *Sphere) GetInverseTransform() Matrix {
	if s.inverse.Empty() {
		m := s.GetTransform()
		s.inverse = m.Inverse()
	}
	return s.inverse
}

func (s *Sphere) GetTransform() Matrix {
	if s.transform.Empty() {
		return MatrixIdentity4x4
	}
	return s.transform
}

func (s *Sphere) SetTransform(m Matrix) {
	s.transform = m
	s.inverse = emptyMatrix
}

func (s *Sphere) EqualShape(o Shape) bool {
	os, ok := o.(*Sphere)
	if !ok {
		return false
	}
	return s.Equal(os)
}

func (s *Sphere) Equal(o *Sphere) bool {
	return true
}

func (s *Sphere) Point() Point {
	return NewPoint(0, 0, 0)
}

func (s *Sphere) String() string {
	return "Sphere()"
}
