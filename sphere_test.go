package rt

import "testing"

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
		s1.SetTransform(Scaling(2, 2, 3))
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

}
