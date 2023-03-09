package rt

import "testing"

func TestMaterial(t *testing.T) {

	t.Run("The default material", func(t *testing.T) {
		m1 := DefaultMaterial()
		equalColor(t, NewColor(1, 1, 1), m1.Color())
		equalValue(t, 0.1, m1.Ambient())
		equalValue(t, 0.9, m1.Diffuse())
		equalValue(t, 0.9, m1.Specular())
		equalValue(t, 200, m1.Shininess())
	})

	t.Run("Lighting with the eye between the light and the surface", func(t *testing.T) {
		var (
			mat      = DefaultMaterial()
			position = NewPoint(0, 0, 0)
			eyev     = NewVector(0, 0, -1)
			normalv  = NewVector(0, 0, -1)
			light    = NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
		)
		color := Lighting(mat, light, position, eyev, normalv)
		equalColor(t, NewColor(1.9, 1.9, 1.9), color)
	})

	t.Run("Lighting with the eye between light and surface, eye offset 45°", func(t *testing.T) {
		var (
			mat      = DefaultMaterial()
			position = NewPoint(0, 0, 0)
			eyev     = NewVector(0, Sqrt2/2, -Sqrt2/2)
			normalv  = NewVector(0, 0, -1)
			light    = NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
		)
		color := Lighting(mat, light, position, eyev, normalv)
		equalColor(t, NewColor(1, 1, 1), color)
	})

	t.Run("Lighting with eye opposite surface, light offset 45°", func(t *testing.T) {
		var (
			mat      = DefaultMaterial()
			position = NewPoint(0, 0, 0)
			eyev     = NewVector(0, 0, -1)
			normalv  = NewVector(0, 0, -1)
			light    = NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1))
		)
		color := Lighting(mat, light, position, eyev, normalv)
		equalColor(t, NewColor(0.7364, 0.7364, 0.7364), color)
	})

	t.Run("Lighting with eye in the path of the reflection vector", func(t *testing.T) {
		var (
			mat      = DefaultMaterial()
			position = NewPoint(0, 0, 0)
			eyev     = NewVector(0, -Sqrt2/2, -Sqrt2/2)
			normalv  = NewVector(0, 0, -1)
			light    = NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1))
		)
		color := Lighting(mat, light, position, eyev, normalv)
		equalColor(t, NewColor(1.6364, 1.6364, 1.6364), color)
	})

	t.Run("Lighting with the light behind the surface", func(t *testing.T) {
		var (
			mat      = DefaultMaterial()
			position = NewPoint(0, 0, 0)
			eyev     = NewVector(0, 0, -1)
			normalv  = NewVector(0, 0, -1)
			light    = NewPointLight(NewPoint(0, 0, 10), NewColor(1, 1, 1))
		)
		color := Lighting(mat, light, position, eyev, normalv)
		equalColor(t, NewColor(0.1, 0.1, 0.1), color)
	})
}
