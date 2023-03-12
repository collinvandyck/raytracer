package rt

import "testing"

func TestTransformations(t *testing.T) {

	// starting at origin and looking in the -z direction is the identity transform
	t.Run("The transformation matrix for the default orientation", func(t *testing.T) {
		from := NewPoint(0, 0, 0)
		to := NewPoint(0, 0, -1)
		up := NewVector(0, 1, 0)
		m1 := ViewTransform(from, to, up)
		equalMatrix(t, MatrixIdentity4x4, m1)
	})

	// starting at origin and looking in the +z direction is like looking in a mirror.
	// front and back are swapped, and left and right are swapped. Reflection is the same
	// as scaling by a negative value.
	//
	// TODO: why is y scaled negatively in this case?
	t.Run("A view transformation matrix looking in positive z direction", func(t *testing.T) {
		from := NewPoint(0, 0, 0)
		to := NewPoint(0, 0, 1)
		up := NewVector(0, 1, 0)
		m1 := ViewTransform(from, to, up)
		equalMatrix(t, Scaling(-1, 1, -1), m1)
	})

	// Eye at a point +8 on the z axis looking at the origin translates the world
	// by -8 in the z direction
	t.Run("The view transformation moves the world", func(t *testing.T) {
		from := NewPoint(0, 0, 8)
		to := NewPoint(0, 0, 0)
		up := NewVector(0, 1, 0)
		m1 := ViewTransform(from, to, up)
		equalMatrix(t, Scaling(0, 0, -8), m1)
	})

	t.Run("An arbitrary view transformation", func(t *testing.T) {
		from := NewPoint(1, 3, 2)
		to := NewPoint(4, -2, 8)
		up := NewVector(1, 1, 0)
		m1 := ViewTransform(from, to, up)
		me := NewMatrixFromTable(`
			+------------------------------------------+
			| -0.50709 | 0.50709 |  0.67612 | -2.36643 |
			|  0.76772 | 0.60609 |  0.12122 | -2.82843 |
			| -0.35857 | 0.59761 | -0.71714 |        0 |
			|        0 |       0 |        0 |        1 |
			+------------------------------------------+
		`)
		equalMatrix(t, me, m1)
	})

}
