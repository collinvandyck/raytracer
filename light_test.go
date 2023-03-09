package rt

import "testing"

func TestLights(t *testing.T) {
	t.Run("A point light has a position and intensity", func(t *testing.T) {
		i1 := NewColor(1, 1, 1)
		p1 := NewPoint(0, 0, 0)
		l1 := NewPointLight(p1, i1)

		equalPoint(t, p1, l1.Position())
		equalColor(t, i1, l1.Intensity())
	})
}
