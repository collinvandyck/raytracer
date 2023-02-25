package rt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColors(t *testing.T) {
	t.Run("Colors are (red, green, blue) tuples", func(t *testing.T) {
		c1 := NewColor(-0.5, 0.4, 1.7)
		require.Equal(t, -0.5, c1.Red())
		require.Equal(t, 0.4, c1.Green())
		require.Equal(t, 1.7, c1.Blue())
	})
	t.Run("Adding colors", func(t *testing.T) {
		c1 := NewColor(0.9, 0.6, 0.75)
		c2 := NewColor(0.7, 0.1, 0.25)
		equalColor(t, NewColor(1.6, 0.7, 1.0), c1.Add(c2))
	})

}

func equalColor(t *testing.T, c1, c2 Color) {
	t.Helper()
	require.EqualValues(t, c1, c2)
	require.True(t, c1.Equal(c2))
}
