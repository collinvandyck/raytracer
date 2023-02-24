package raytracer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTuples(t *testing.T) {
	t.Run("A tuple with w=1 is a point", func(t *testing.T) {
		tup := tuple{4.3, -4.2, 3.1, 1.0}
		require.Equal(t, tup.x, 4.3)
		require.Equal(t, tup.y, -4.2)
		require.Equal(t, tup.z, 3.1)
		require.Equal(t, tup.w, 1.0)
		require.True(t, tup.isPoint())
		require.False(t, tup.isVector())
	})
	t.Run("A point creates a tuple with w=1", func(t *testing.T) {
		point := point(4, -4, 3)
		require.Equal(t, tuple{4, -4, 3, 1}, point)
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
	t.Run("A vector creates a tuple with w=0", func(t *testing.T) {
		vector := vector(4, -4, 3)
		require.Equal(t, tuple{4, -4, 3, 0}, vector)
	})
}
