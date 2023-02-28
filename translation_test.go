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

func TestScaling(t *testing.T) {

	t.Run("A scaling matrix applied to a point", func(t *testing.T) {
		t1 := Scaling(2, 3, 4)
		p1 := NewPoint(-4, 6, 8)
		equalPoint(t, NewPoint(-8, 18, 32), t1.MultiplyPoint(p1))
	})

}
