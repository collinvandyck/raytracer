package rt

import (
	"math"
	"testing"
)

func TestSphere(t *testing.T) {

	t.Run("A sphere's default transformation", func(t *testing.T) {
		s1 := NewSphere()
		equalMatrix(t, MatrixIdentity4x4, s1.GetTransform())
	})

	t.Run("Changing a sphere's transformation", func(t *testing.T) {
		s1 := NewSphere()
		t1 := Scaling(2, 2, 3)
		s1.SetTransform(t1)
		equalMatrix(t, t1, s1.GetTransform())
	})

	t.Run("Intersecting a scaled sphere with a ray", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s1 := NewSphere()
		s1.SetTransform(Scaling(2, 2, 2))
		xs := IntersectSphere(s1, r1)
		xe := NewIntersections(NewIntersection(3, s1), NewIntersection(7, s1))
		equalIntersections(t, xe, xs)
	})

	t.Run("Intersecting a translated sphere with a ray", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s1 := NewSphere()
		s1.SetTransform(Translation(5, 0, 0))
		xs := IntersectSphere(s1, r1)
		xe := NewIntersections()
		equalIntersections(t, xe, xs)
	})

	t.Run("The normal on a sphere at a point on the x axis", func(t *testing.T) {
		s1 := NewSphere()
		n1 := s1.NormalAt(NewPoint(1, 0, 0))
		equalVector(t, NewVector(1, 0, 0), n1)
	})

	t.Run("The normal on a sphere at a point on the y axis", func(t *testing.T) {
		s1 := NewSphere()
		n1 := s1.NormalAt(NewPoint(0, 1, 0))
		equalVector(t, NewVector(0, 1, 0), n1)
	})

	t.Run("The normal on a sphere at a point on the z axis", func(t *testing.T) {
		s1 := NewSphere()
		n1 := s1.NormalAt(NewPoint(0, 0, 1))
		equalVector(t, NewVector(0, 0, 1), n1)
	})

	t.Run("The normal on a sphere at a non-axial point", func(t *testing.T) {
		s1 := NewSphere()
		n1 := s1.NormalAt(NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
		equalVector(t, NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3), n1)
	})

	t.Run("The normal is a normalized vector", func(t *testing.T) {
		s1 := NewSphere()
		n1 := s1.NormalAt(NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
		equalVector(t, n1.Normalize(), n1)
	})

	t.Run("Computing the normal on a translated sphere", func(t *testing.T) {
		s1 := NewSphere()
		s1.SetTransform(Translation(0, 1, 0))
		n1 := s1.NormalAt(NewPoint(0, 1.70711, -0.70711))
		equalVector(t, NewVector(0, 0.70711, -0.70711), n1)
	})

	t.Run("Computing the normal on a transformed sphere", func(t *testing.T) {
		s1 := NewSphere()
		m1 := Scaling(1, 0.5, 1).Multiply(RotationZ(Pi / 5))
		s1.SetTransform(m1)
		n1 := s1.NormalAt(NewPoint(0, Sqrt2/2, -Sqrt2/2))
		equalVector(t, NewVector(0, 0.97014, -0.24254), n1)
	})

}
