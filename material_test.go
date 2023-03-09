package rt

import "testing"

func TestMaterial(t *testing.T) {

	t.Run("The default material", func(t *testing.T) {
		m1 := DefaultMaterial()
		equalColor(t, NewColor(1, 1, 1), m1.GetColor())
		equalValue(t, 0.1, m1.GetAmbient())
		equalValue(t, 0.9, m1.GetDiffuse())
		equalValue(t, 0.9, m1.GetSpecular())
		equalValue(t, 200, m1.GetShininess())
	})
}
