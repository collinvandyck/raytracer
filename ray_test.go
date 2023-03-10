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

	t.Run("Translating a ray", func(t *testing.T) {
		o1 := NewPoint(1, 2, 3)
		d1 := NewVector(0, 1, 0)
		r1 := NewRay(o1, d1)
		m1 := Translation(3, 4, 5)
		r2 := r1.Transform(m1)

		equalPoint(t, NewPoint(4, 6, 8), r2.Origin())
		equalVector(t, NewVector(0, 1, 0), r2.Direction())
	})

	t.Run("Scaling a ray", func(t *testing.T) {
		o1 := NewPoint(1, 2, 3)
		d1 := NewVector(0, 1, 0)
		r1 := NewRay(o1, d1)
		m1 := Scaling(2, 3, 4)
		r2 := r1.Transform(m1)

		equalPoint(t, NewPoint(2, 6, 12), r2.Origin())
		equalVector(t, NewVector(0, 3, 0), r2.Direction())
	})

}
