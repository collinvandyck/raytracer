package rt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanvas(t *testing.T) {
	t.Run("Creating a canvas", func(t *testing.T) {
		c1 := NewCanvas(10, 20)
		require.Equal(t, 10, c1.Width())
		require.Equal(t, 20, c1.Height())
	})
}
