package rt

import "testing"

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
	})

}
