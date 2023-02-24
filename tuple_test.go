package raytracer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTuples(t *testing.T) {
	t.Run("A tuple with w=0 is a point", func(t *testing.T) {
		tup := tuple{4.3, -4.2, 3.1, 1.0}
		require.Equal(t, tup.x, 4.3)
		require.Equal(t, tup.y, -4.2)
		require.Equal(t, tup.z, 3.1)
		require.Equal(t, tup.w, 1.0)
		require.True(t, tup.isPoint())
		require.False(t, tup.isVector())
	})
	t.Run("A point is a point", func(t *testing.T) {
		point := point(4.3, -4.2, 3.1)
		require.Equal(t, point.x, 4.3)
		require.Equal(t, point.y, -4.2)
		require.Equal(t, point.z, 3.1)
		require.Equal(t, point.w, 1.0)
		require.True(t, point.isPoint())
		require.False(t, point.isVector())
	})
	t.Run("A tuple with w=0 is a vector", func(t *testing.T) {
		tup := tuple{4.3, -4.2, 3.1, 0.0}
		require.Equal(t, tup.x, 4.3)
		require.Equal(t, tup.y, -4.2)
		require.Equal(t, tup.z, 3.1)
		require.Equal(t, tup.w, 0.0)
		require.False(t, tup.isPoint())
		require.True(t, tup.isVector())
	})
	t.Run("A vector is a vector", func(t *testing.T) {
		vector := vector(4.3, -4.2, 3.1)
		require.Equal(t, vector.x, 4.3)
		require.Equal(t, vector.y, -4.2)
		require.Equal(t, vector.z, 3.1)
		require.Equal(t, vector.w, 0.0)
		require.False(t, vector.isPoint())
		require.True(t, vector.isVector())
	})
}
