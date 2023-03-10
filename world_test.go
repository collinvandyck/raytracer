package rt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorld(t *testing.T) {

	t.Run("Creating a world", func(t *testing.T) {
		w1 := NewWorld()
		require.Len(t, w1.Shapes(), 0)
		require.Nil(t, w1.Light())
	})

	t.Run("The default world", func(t *testing.T) {
		l1 := NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1))
		s1 := NewSphere()
		m1 := DefaultMaterial()
		m1.SetColor(NewColor(0.8, 1.0, 0.6))
		m1.SetDiffuse(0.7)
		m1.SetSpecular(0.2)
		s1.SetMaterial(m1)
		s2 := NewSphere()
		s2.SetTransform(Scaling(0.5, 0.5, 0.5))

		w1 := NewDefaultWorld()
		equalLight(t, l1, w1.Light())
		require.Len(t, w1.Shapes(), 2)
		require.Contains(t, w1.Shapes(), s1)
		require.Contains(t, w1.Shapes(), s2)
		require.EqualValues(t, []Shape{s1, s2}, w1.Shapes())
	})

	t.Run("Intersect a world with  a ray", func(t *testing.T) {
		w1 := NewDefaultWorld()
		r1 := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		xs := IntersectWorld(w1, r1)

		require.EqualValues(t, 4, xs.Len())
		require.EqualValues(t, 4, xs.Get(0).Value())
		require.EqualValues(t, 4.5, xs.Get(1).Value())
		require.EqualValues(t, 5.5, xs.Get(2).Value())
		require.EqualValues(t, 6, xs.Get(3).Value())
	})

	t.Run("Shading an intersection", func(t *testing.T) {
		w1 := NewDefaultWorld()
		r1 := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s1 := w1.Shape(0)
		i1 := NewIntersection(4, s1)
		cs := PrepareComputations(i1, r1)
		c1 := ShadeHit(w1, cs)
		equalColor(t, NewColor(0.38066, 0.47583, 0.2855), c1)
	})

	t.Run("Shading an intersection from the inside", func(t *testing.T) {
		w1 := NewDefaultWorld()
		w1.SetLight(NewPointLight(NewPoint(0, 0.25, 0), NewColor(1, 1, 1)))
		r1 := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		s1 := w1.Shape(1)
		i1 := NewIntersection(0.5, s1)
		cs := PrepareComputations(i1, r1)
		c1 := ShadeHit(w1, cs)
		equalColor(t, NewColor(0.90498, 0.90498, 0.90498), c1)
	})

	t.Run("The color when a ray misses", func(t *testing.T) {
		w1 := NewDefaultWorld()
		r1 := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))
		c1 := w1.ColorAt(r1)
		equalColor(t, NewColor(0, 0, 0), c1)
	})

	t.Run("The color when a ray hits", func(t *testing.T) {
		w1 := NewDefaultWorld()
		r1 := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		c1 := w1.ColorAt(r1)
		equalColor(t, NewColor(0.38066, 0.47583, 0.2855), c1)
	})

	// ray is inside the outer sphere, but outside the inner sphere, and pointing
	// at the inner sphere
	t.Run("The color with an intersection behind the ray", func(t *testing.T) {
		w1 := NewDefaultWorld()
		outer := w1.Shape(0)
		outer.Material().SetAmbient(1)
		inner := w1.Shape(1)
		inner.Material().SetAmbient(1)
		ray := NewRay(NewPoint(0, 0, 0.75), NewVector(0, 0, -1))
		c1 := w1.ColorAt(ray)
		equalColor(t, inner.Material().Color(), c1)
	})

}
