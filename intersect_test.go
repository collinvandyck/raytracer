package rt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersect(t *testing.T) {

	t.Run("A ray intersects a sphere at two points", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s1 := NewSphere()
		xs := IntersectSphere(s1, r1)
		xe := NewIntersections(NewIntersection(4, s1), NewIntersection(6, s1))
		equalIntersections(t, xe, xs)
	})

	t.Run("A ray intersects a sphere at a tangent", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 1, -5), NewVector(0, 0, 1))
		s1 := NewSphere()
		xs := IntersectSphere(s1, r1)
		xe := NewIntersections(NewIntersection(5, s1), NewIntersection(5, s1))
		equalIntersections(t, xe, xs)
	})

	t.Run("A ray misses a sphere", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))
		s1 := NewSphere()
		xs := IntersectSphere(s1, r1)
		xe := NewIntersections()
		equalIntersections(t, xe, xs)
	})

	t.Run("A ray originates inside a sphere", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		s1 := NewSphere()
		xs := IntersectSphere(s1, r1)
		xe := NewIntersections(NewIntersection(-1, s1), NewIntersection(1, s1))
		equalIntersections(t, xe, xs)
	})

	t.Run("A sphere is behind a ray", func(t *testing.T) {
		r1 := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
		s1 := NewSphere()
		xs := IntersectSphere(s1, r1)
		xe := NewIntersections(NewIntersection(-6, s1), NewIntersection(-4, s1))
		equalIntersections(t, xe, xs)
	})

	t.Run("An intersection encapsulates t and object", func(t *testing.T) {
		s1 := NewSphere()
		i1 := NewIntersection(3.5, s1)
		equalValue(t, 3.5, i1.Value())
		equalShape(t, s1, i1.Shape())
	})

	t.Run("Aggregating intersections", func(t *testing.T) {
		s1 := NewSphere()
		i1 := NewIntersection(1, s1)
		i2 := NewIntersection(2, s1)
		xs := NewIntersections(i1, i2)

		require.Len(t, xs, 2)
		equalValue(t, 1, xs[0].Value())
		equalValue(t, 2, xs[1].Value())
	})

	t.Run("The hit, when all intersections have positive t", func(t *testing.T) {
		s1 := NewSphere()
		i1 := NewIntersection(1, s1)
		i2 := NewIntersection(2, s1)
		xs := NewIntersections(i2, i1)
		h1, ok := Hit(xs)
		require.True(t, ok)
		equalIntersection(t, i1, h1)
	})

	t.Run("The hit, when some intersections have negative t", func(t *testing.T) {
		s1 := NewSphere()
		i1 := NewIntersection(-1, s1)
		i2 := NewIntersection(+1, s1)
		xs := NewIntersections(i2, i1)
		h1, ok := Hit(xs)
		require.True(t, ok)
		equalIntersection(t, i2, h1)
	})

}
