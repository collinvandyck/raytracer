package rt

import "testing"

func TestIntersect(t *testing.T) {

	t.Run("A ray intersects a sphere at two points", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s1 := NewSphere()
		vs := IntersectSphere(s1, r1)
		equalValueSlice(t, []value{4, 6}, vs)
	})

	t.Run("A ray intersects a sphere at a tangent", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 1, -5), NewVector(0, 0, 1))
		s1 := NewSphere()
		vs := IntersectSphere(s1, r1)
		equalValueSlice(t, []value{5, 5}, vs)
	})

	t.Run("A ray misses a sphere", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))
		s1 := NewSphere()
		vs := IntersectSphere(s1, r1)
		equalValueSlice(t, nil, vs)
	})

	t.Run("A ray originates inside a sphere", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		s1 := NewSphere()
		vs := IntersectSphere(s1, r1)
		equalValueSlice(t, []value{-1, 1}, vs)
	})

	t.Run("A sphere is behind a ray", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
		s1 := NewSphere()
		vs := IntersectSphere(s1, r1)
		equalValueSlice(t, []value{-6, -4}, vs)
	})

	t.Run("An intersection encapsulates t and object", func(t *testing.T) {
		s1 := NewSphere()
		i1 := NewSphereIntersection(3.5, s1)
		equalValue(t, 3.5, i1.GetT())
		equalShape(t, s1, i1.Shape())
	})

}
