package rt

import (
	"testing"
)

func TestRays(t *testing.T) {

	t.Run("Creating and querying a ray", func(t *testing.T) {
		o1 := NewPoint(1, 2, 3)
		d1 := NewVector(4, 5, 6)
		r1 := NewRay(o1, d1)
		equalPoint(t, o1, r1.Origin())
		equalVector(t, d1, r1.Direction())
	})

	t.Run("Computing a point from a distance", func(t *testing.T) {
		o1 := NewPoint(2, 3, 4)
		d1 := NewVector(1, 0, 0)
		r1 := NewRay(o1, d1)

		equalPoint(t, NewPoint(2, 3, 4), r1.Position(0))
		equalPoint(t, NewPoint(3, 3, 4), r1.Position(1))
		equalPoint(t, NewPoint(1, 3, 4), r1.Position(-1))
		equalPoint(t, NewPoint(4.5, 3, 4), r1.Position(2.5))
	})

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
	})

}
