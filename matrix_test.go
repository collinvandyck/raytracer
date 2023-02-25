package rt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatrix(t *testing.T) {
	t.Run("Constructing and testing a 4x4 matrix", func(t *testing.T) {
		m := Matrix{
			[]float{1, 2, 3, 4},
			[]float{5.5, 6.5, 7.5, 8.5},
			[]float{9, 10, 11, 12},
			[]float{13.5, 14.5, 15.5, 16.5},
		}
		require.Equal(t, 1.0, m.Get(0, 0))
	})
}
