package rt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorld(t *testing.T) {

	t.Run("Creating a world", func(t *testing.T) {
		w1 := NewWorld()
		require.Len(t, w1.Shapes(), 0)
		require.Len(t, w1.Lights(), 0)
	})

}
