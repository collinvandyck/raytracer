package rt

import (
	"testing"

	"github.com/stretchr/testify/require"
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

	t.Run("An intersection has values", func(t *testing.T) {
		i1 := NewIntersection()
		require.Equal(t, 0, i1.Len())
		require.Len(t, i1.Get(), 0)
		equalIntersection(t, NewIntersection(), i1)

		i2 := NewIntersection(1, 2)
		require.Equal(t, 2, i2.Len())
		require.EqualValues(t, []float{1, 2}, i2.Get())
		equalIntersection(t, NewIntersection(1, 2), i2)
	})

	t.Run("A ray intersects a sphere at two points", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s1 := NewSphere()

		i1 := r1.IntersectSphere(s1)
		equalIntersection(t, NewIntersection(4, 6), i1)
	})

	t.Run("A ray intersects a sphere at a tangent", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 1, -5), NewVector(0, 0, 1))
		s1 := NewSphere()

		i1 := r1.IntersectSphere(s1)
		equalIntersection(t, NewIntersection(5, 5), i1)
	})

	t.Run("A ray misses a sphere", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))
		s1 := NewSphere()

		i1 := r1.IntersectSphere(s1)
		equalIntersection(t, NewIntersection(), i1)
	})

	t.Run("A ray originates inside a sphere", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		s1 := NewSphere()

		i1 := r1.IntersectSphere(s1)
		equalIntersection(t, NewIntersection(-1, 1), i1)
	})

	t.Run("A sphere is behind a ray", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
		s1 := NewSphere()

		i1 := r1.IntersectSphere(s1)
		equalIntersection(t, NewIntersection(-6, -4), i1)
	})

}
