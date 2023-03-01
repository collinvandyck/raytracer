package rt

import (
	"testing"
)

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

	t.Run("A scaling matrix applied to a vector", func(t *testing.T) {
		t1 := Scaling(2, 3, 4)
		v1 := NewVector(-4, 6, 8)
		equalVector(t, NewVector(-8, 18, 32), t1.MultiplyVector(v1))
	})

	t.Run("Multiplying by the inverse of a scaling matrix", func(t *testing.T) {
		t1 := Scaling(2, 3, 4)
		i1 := t1.Inverse()
		v1 := NewVector(-4, 6, 8)
		equalVector(t, NewVector(-2, 2, 2), i1.MultiplyVector(v1))
	})

	t.Run("Reflection is scaling by a negative value", func(t *testing.T) {
		t1 := Scaling(1, -1, 1)
		p1 := NewPoint(2, 3, 4)
		equalPoint(t, NewPoint(2, -3, 4), t1.MultiplyPoint(p1))
	})

}

func TestRotation(t *testing.T) {

	t.Run("Rotating a point around the x axis", func(t *testing.T) {
		p1 := NewPoint(0, 1, 0)
		halfQuarter := RotationX(Pi / 4)
		fullQuarter := RotationX(Pi / 2)
		equalPoint(t, NewPoint(0, Sqrt2/2, Sqrt2/2), halfQuarter.MultiplyPoint(p1))
		equalPoint(t, NewPoint(0, 0, 1), fullQuarter.MultiplyPoint(p1))
	})

	t.Run("The inverse of an x-rotation rotates in the opposite direction", func(t *testing.T) {
		p1 := NewPoint(0, 1, 0)
		halfQuarter := RotationX(Pi / 4)
		inv := halfQuarter.Inverse()
		equalPoint(t, NewPoint(0, Sqrt2/2, -Sqrt2/2), inv.MultiplyPoint(p1))
	})

	t.Run("Rotating a point around the y axis", func(t *testing.T) {
		p1 := NewPoint(0, 0, 1)
		halfQuarter := RotationY(Pi / 4)
		fullQuarter := RotationY(Pi / 2)
		equalPoint(t, NewPoint(Sqrt2/2, 0, Sqrt2/2), halfQuarter.MultiplyPoint(p1))
		equalPoint(t, NewPoint(1, 0, 0), fullQuarter.MultiplyPoint(p1))
	})

	t.Run("Rotating a point around the z axis", func(t *testing.T) {
		p1 := NewPoint(0, 1, 0)
		halfQuarter := RotationZ(Pi / 4)
		fullQuarter := RotationZ(Pi / 2)
		equalPoint(t, NewPoint(-Sqrt2/2, Sqrt2/2, 0), halfQuarter.MultiplyPoint(p1))
		equalPoint(t, NewPoint(-1, 0, 0), fullQuarter.MultiplyPoint(p1))
	})

}

func TestShearing(t *testing.T) {

	t.Run("A shearing transformation moves x in proportion to y", func(t *testing.T) {
		t1 := Shearing(1, 0, 0, 0, 0, 0)
		p1 := NewPoint(2, 3, 4)
		equalPoint(t, NewPoint(5, 3, 4), t1.MultiplyPoint(p1))
	})

	t.Run("A shearing transformation moves x in proportion to z", func(t *testing.T) {
		t1 := Shearing(0, 1, 0, 0, 0, 0)
		p1 := NewPoint(2, 3, 4)
		equalPoint(t, NewPoint(6, 3, 4), t1.MultiplyPoint(p1))
	})

	t.Run("A shearing transformation moves y in proportion to x", func(t *testing.T) {
		t1 := Shearing(0, 0, 1, 0, 0, 0)
		p1 := NewPoint(2, 3, 4)
		equalPoint(t, NewPoint(2, 5, 4), t1.MultiplyPoint(p1))
	})

	t.Run("A shearing transformation moves y in proportion to z", func(t *testing.T) {
		t1 := Shearing(0, 0, 0, 1, 0, 0)
		p1 := NewPoint(2, 3, 4)
		equalPoint(t, NewPoint(2, 7, 4), t1.MultiplyPoint(p1))
	})

	t.Run("A shearing transformation moves z in proportion to x", func(t *testing.T) {
		t1 := Shearing(0, 0, 0, 0, 1, 0)
		p1 := NewPoint(2, 3, 4)
		equalPoint(t, NewPoint(2, 3, 6), t1.MultiplyPoint(p1))
	})

	t.Run("A shearing transformation moves z in proportion to y", func(t *testing.T) {
		t1 := Shearing(0, 0, 0, 0, 0, 1)
		p1 := NewPoint(2, 3, 4)
		equalPoint(t, NewPoint(2, 3, 7), t1.MultiplyPoint(p1))
	})

}

func TestChaining(t *testing.T) {

	t.Run("Individual transformations are applied in sequence", func(t *testing.T) {
		p1 := NewPoint(1, 0, 1)
		a1 := RotationX(Pi / 2)
		b1 := Scaling(5, 5, 5)
		c1 := Translation(10, 5, 7)

		// apply rotation first
		p2 := a1.MultiplyPoint(p1)
		equalPoint(t, NewPoint(1, -1, 0), p2)

		// then apply scaling
		p3 := b1.MultiplyPoint(p2)
		equalPoint(t, NewPoint(5, -5, 0), p3)

		// then apply translation
		p4 := c1.MultiplyPoint(p3)
		equalPoint(t, NewPoint(15, 0, 7), p4)
	})
}
