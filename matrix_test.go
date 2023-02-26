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
	t.Run("Matrix equality with identical matrices", func(t *testing.T) {
		m1 := NewMatrixFromTable(`
			+---------------+
			| 1 | 2 | 3 | 4 |
			| 5 | 6 | 7 | 8 |
			| 9 | 8 | 7 | 6 |
			| 5 | 4 | 3 | 2 |
			+---------------+
		`)
		m2 := NewMatrixFromTable(`
			+---------------+
			| 1 | 2 | 3 | 4 |
			| 5 | 6 | 7 | 8 |
			| 9 | 8 | 7 | 6 |
			| 5 | 4 | 3 | 2 |
			+---------------+
		`)
		equalMatrix(t, m1, m2)
	})
	t.Run("Matrix equality with different matrices", func(t *testing.T) {
		m1 := NewMatrixFromTable(`
			+---------------+
			| 1 | 2 | 3 | 4 |
			| 5 | 6 | 7 | 8 |
			| 9 | 8 | 7 | 6 |
			| 5 | 4 | 3 | 2 |
			+---------------+
		`)
		m2 := NewMatrixFromTable(`
			+---------------+
			| 2 | 3 | 4 | 5 |
			| 6 | 7 | 8 | 9 |
			| 8 | 7 | 6 | 5 |
			| 4 | 3 | 2 | 1 |
			+---------------+
		`)
		notEqualMatrix(t, m1, m2)
	})
	t.Run("Multiplying two matrices", func(t *testing.T) {
		m1 := NewMatrixFromTable(`
			+---------------+
			| 1 | 2 | 3 | 4 |
			| 5 | 6 | 7 | 8 |
			| 9 | 8 | 7 | 6 |
			| 5 | 4 | 3 | 2 |
			+---------------+
		`)
		m2 := NewMatrixFromTable(`
			+-----------------+
			| -2 | 1 | 2 | 3  |
			| 3  | 2 | 1 | -1 |
			| 4  | 3 | 6 | 5  |
			| 1  | 2 | 7 | 8  |
			+-----------------+
		`)
		me := NewMatrixFromTable(`
			+---------------------+
			| 20 | 22 | 50  | 48  |
			| 44 | 54 | 114 | 108 |
			| 40 | 58 | 110 | 102 |
			| 16 | 26 | 46  | 42  |
			+---------------------+
		`)
		equalMatrix(t, me, m1.Multiply(m2))
	})
	t.Run("A matrix multiplied by a tuple", func(t *testing.T) {
		m1 := NewMatrixFromTable(`
			+---------------+
			| 1 | 2 | 3 | 4 |
			| 2 | 4 | 4 | 2 |
			| 8 | 6 | 4 | 1 |
			| 0 | 0 | 0 | 1 |
			+---------------+
		`)
		t1 := NewTuple(1, 2, 3, 1)
		te := NewTuple(18, 24, 33, 1)
		equalTuple(t, te, m1.MultiplyTuple4(t1))
	})
	t.Run("Multiplying a matrix by the identity matrix", func(t *testing.T) {
		m1 := NewMatrixFromTable(`
			+-----------------+
			| 0 | 1 | 2  | 4  |
			| 1 | 2 | 4  | 8  |
			| 2 | 4 | 8  | 16 |
			| 4 | 8 | 16 | 32 |
			+-----------------+
		`)
		e1 := NewMatrixFromTable(`
			+-----------------+
			| 0 | 1 | 2  | 4  |
			| 1 | 2 | 4  | 8  |
			| 2 | 4 | 8  | 16 |
			| 4 | 8 | 16 | 32 |
			+-----------------+
		`)
		equalMatrix(t, e1, m1.Multiply(MatrixIdentity4x4))
	})
	t.Run("Transposing a matrix", func(t *testing.T) {
		m1 := NewMatrixFromTable(`
			+---------------+
			| 0 | 9 | 3 | 0 |
			| 9 | 8 | 0 | 8 |
			| 1 | 8 | 5 | 3 |
			| 0 | 0 | 5 | 8 |
			+---------------+
		`)
		me := NewMatrixFromTable(`
			+---------------+
			| 0 | 9 | 1 | 0 |
			| 9 | 8 | 8 | 0 |
			| 3 | 0 | 5 | 5 |
			| 0 | 8 | 3 | 8 |
			+---------------+
		`)
		equalMatrix(t, me, m1.Transpose())
	})
}

func notEqualMatrix(t *testing.T, m1, m2 Matrix) {
	require.NotEqual(t, m1, m2)
	require.False(t, m1.Equal(m2))
	require.False(t, m2.Equal(m1))
}

func equalMatrix(t *testing.T, m1, m2 Matrix) {
	require.Equal(t, m1, m2)
	require.True(t, m1.Equal(m2))
	require.True(t, m2.Equal(m1))
}
