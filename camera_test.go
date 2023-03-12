package rt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCamera(t *testing.T) {

	t.Run("Constructing a camera", func(t *testing.T) {
		var (
			hsize = 160
			vsize = 120
			fov   = Pi / 2
		)
		c := NewCamera(hsize, vsize, fov)
		require.Equal(t, 160, c.HSize())
		require.Equal(t, 120, c.VSize())
		equalValue(t, Pi/2, c.FOV())
		equalMatrix(t, MatrixIdentity4x4, c.Transform())
	})

	t.Run("The pixel size for a horizontal canvas", func(t *testing.T) {
		c := NewCamera(200, 125, Pi/2)
		equalValue(t, 0.01, c.PixelSize())
	})

	t.Run("The pixel size for a vertical canvas", func(t *testing.T) {
		c := NewCamera(125, 200, Pi/2)
		equalValue(t, 0.01, c.PixelSize())
	})

	t.Run("Constructing a ray through the center of the canvas", func(t *testing.T) {
		c := NewCamera(201, 101, Pi/2)
		ray := c.RayForPixel(100, 50)
		equalPoint(t, NewPoint(0, 0, 0), ray.Origin())
		equalVector(t, NewVector(0, 0, -1), ray.Direction())
	})

	t.Run("Constructing a ray through a corner of the canvas", func(t *testing.T) {
		c := NewCamera(201, 101, Pi/2)
		ray := c.RayForPixel(0, 0)
		equalPoint(t, NewPoint(0, 0, 0), ray.Origin())
		equalVector(t, NewVector(0.66519, 0.33259, -0.66851), ray.Direction())
	})

	t.Run("Constructing a ray when the camera is transformed", func(t *testing.T) {
		c := NewCamera(201, 101, Pi/2)
		xf := RotationY(Pi / 4).Multiply(Translation(0, -2, 5))
		c.SetTransform(xf)
		ray := c.RayForPixel(100, 50)
		equalPoint(t, NewPoint(0, 2, -5), ray.Origin())
		equalVector(t, NewVector(Sqrt2/2, 0, -Sqrt2/2), ray.Direction())
	})

}
