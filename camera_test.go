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

}
