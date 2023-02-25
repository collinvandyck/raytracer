package rt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatrix(t *testing.T) {
	t.Run("Constructing and testing a 4x4 matrix from table", func(t *testing.T) {
		m := NewMatrixFromTable(`
		    +---------------------------+
			| 1    | 2    | 3    | 4    |
			| 5.5  | 6.5  | 7.5  | 8.5  |
			| 9    | 10   | 11   | 12   |
			| 13.5 | 14.5 | 15.5 | 16.5 |
		    +---------------------------+
		`)
		require.Equal(t, 1.0, m.Get(0, 0))
		require.Equal(t, 4.0, m.Get(0, 3))
		require.Equal(t, 5.5, m.Get(1, 0))
		require.Equal(t, 7.5, m.Get(1, 2))
		require.Equal(t, 11.0, m.Get(2, 2))
		require.Equal(t, 13.5, m.Get(3, 0))
		require.Equal(t, 15.5, m.Get(3, 2))
	})
	t.Run("A 2x2 matrix ought to be representable", func(t *testing.T) {
		m := NewMatrixFromTable(`
		    +---------+
			| -3 |  5 |
			|  1 | -2 |
		    +---------+
		`)
		require.Equal(t, -3.0, m.Get(0, 0))
		require.Equal(t, 5.0, m.Get(0, 1))
		require.Equal(t, 1.0, m.Get(1, 0))
		require.Equal(t, -2.0, m.Get(1, 1))
	})
	t.Run("A 3x3 matrix ought to be representable", func(t *testing.T) {
		m := NewMatrixFromTable(`
		    +---------------+
			| -3 |  5 |  0  |
			|  1 | -2 | -7  |
			|  0 |  1 |  1  |
		    +---------------+
		`)
		require.Equal(t, -3.0, m.Get(0, 0))
		require.Equal(t, 5.0, m.Get(0, 1))
		require.Equal(t, 0.0, m.Get(0, 2))
		require.Equal(t, 1.0, m.Get(1, 0))
		require.Equal(t, -2.0, m.Get(1, 1))
		require.Equal(t, -7.0, m.Get(1, 2))
		require.Equal(t, 0.0, m.Get(2, 0))
		require.Equal(t, 1.0, m.Get(2, 1))
		require.Equal(t, 1.0, m.Get(2, 2))
	})
}
