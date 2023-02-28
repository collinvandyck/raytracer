package rt

import "testing"

func TestTranslation(t *testing.T) {

	t.Run("Multiplying by a translation matrix", func(t *testing.T) {
		t1 := Translation(5, -3, 2)
		p1 := NewPoint(-3, 4, 5)
		equalPoint(t, NewPoint(2, 1, 7), t1.MultiplyPoint(p1))
	})

	t.Run("Multiplying by the inverse of a translation matrix", func(t *testing.T) {
		t1 := Translation(5, -3, 2)
		i1 := t1.Inverse()
		p1 := NewPoint(-3, 4, 5)
		equalPoint(t, NewPoint(-8, 7, 3), i1.MultiplyPoint(p1))
	})

	t.Run("Translation does not affect vectors", func(t *testing.T) {
		t1 := Translation(5, -3, 2)
		v1 := NewVector(-3, 4, 5)
		equalVector(t, NewVector(-3, 4, 5), t1.MultiplyVector(v1))
	})

}
