package rt

import (
	"math"
)

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
	mt := s1.GetTransposedInverseTransform()         // transposed inverse transform
	op := s1.GetInverseTransform().MultiplyPoint(p1) // object space
	on := op.SubtractPoint(Origin)                   // object normal
	wn := mt.MultiplyVector(on)                      // world normal
	wn.SetW(0)
	normalized := wn.Normalize()
	return normalized
}

type Sphere struct {
	transform         Matrix
	inverse           Matrix
	transposedInverse Matrix
}

func NewSphere() *Sphere {
	return &Sphere{}
}

func (s *Sphere) NormalAt(point Point) Vector {
	return NormalAtSphere(s, point)
}

func (s *Sphere) GetTransposedInverseTransform() Matrix {
	if s.transposedInverse.Empty() {
		m := s.GetInverseTransform()
		s.transposedInverse = m.Transpose()
	}
	return s.transposedInverse
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
	s.transposedInverse = emptyMatrix
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
