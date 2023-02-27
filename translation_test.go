package rt

import "testing"

func TestTranslation(t *testing.T) {
	t.Run("Multiplying by a translation matrix", func(t *testing.T) {
		t1 := Translation(5, -3, 2)
		p1 := NewPoint(-3, 4, 5)
		equalPoint(t, NewPoint(2, 1, 7), t1.MultiplyPoint(p1))
	})
}
