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
}
