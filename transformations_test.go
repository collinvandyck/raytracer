package rt

import "testing"

func TestTransformations(t *testing.T) {

	t.Run("The transformation matrix for the default orientation", func(t *testing.T) {
		from := NewPoint(0, 0, 0)
		to := NewPoint(0, 0, -1)
		up := NewVector(0, 1, 0)
		m1 := ViewTransform(from, to, up)
		equalMatrix(t, MatrixIdentity4x4, m1)
	})

}
