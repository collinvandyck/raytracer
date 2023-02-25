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
}
