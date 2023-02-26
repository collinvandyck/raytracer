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
	t.Run("Transposing the identity matrix", func(t *testing.T) {
		m1 := MatrixIdentity4x4
		me := MatrixIdentity4x4
		equalMatrix(t, me, m1.Transpose())
	})
	t.Run("Determinants", func(t *testing.T) {
		t.Run("Calculating the determinant of a 2x2 matrix", func(t *testing.T) {
			m1 := NewMatrixFromTable(`
				+--------+
				|  1 | 5 |
				| -3 | 2 |
				+--------+
			`)
			require.EqualValues(t, 17, m1.Determinant())
		})
		t.Run("Calculating the determinant of a 3x3 matrix", func(t *testing.T) {
			m1 := NewMatrixFromTable(`
			+-------------------+
			|  1.0 | 2.0 |  6.0 |
			| -5.0 | 8.0 | -4.0 |
			|  2.0 | 6.0 |  4.0 |
			+-------------------+
			`)
			cf00 := m1.Cofactor(0, 0)
			cf01 := m1.Cofactor(0, 1)
			cf02 := m1.Cofactor(0, 2)
			dt1 := m1.Determinant()

			require.EqualValues(t, 56, cf00)
			require.EqualValues(t, 12, cf01)
			require.EqualValues(t, -46, cf02)
			require.EqualValues(t, -196, dt1)
		})
		t.Run("Calculating the determinant of a 4x4 matrix", func(t *testing.T) {
			m1 := NewMatrixFromTable(`
			+---------------------------+
			| -2.0 | -8.0 |  3.0 |  5.0 |
			| -3.0 |  1.0 |  7.0 |  3.0 |
			|  1.0 |  2.0 | -9.0 |  6.0 |
			| -6.0 |  7.0 |  7.0 | -9.0 |
			+---------------------------+
			`)
			cf00 := m1.Cofactor(0, 0)
			cf01 := m1.Cofactor(0, 1)
			cf02 := m1.Cofactor(0, 2)
			cf03 := m1.Cofactor(0, 3)
			dt1 := m1.Determinant()

			require.EqualValues(t, 690, cf00)
			require.EqualValues(t, 447, cf01)
			require.EqualValues(t, 210, cf02)
			require.EqualValues(t, 51, cf03)
			require.EqualValues(t, -4071, dt1)
		})
	})
	t.Run("Submatrixes", func(t *testing.T) {
		t.Run("A submatrix of a 3x3 matrix is a 2x2 matrix", func(t *testing.T) {
			m1 := NewMatrixFromTable(`
				+---------------+
				|  1 | 5 |  0  |
				| -3 | 2 |  7  |
				|  0 | 6 | -3  |
				+---------------+
			`)
			me := NewMatrixFromTable(`
				+--------+
				| -3 | 2 |
				|  0 | 6 |
				+--------+
			`)
			equalMatrix(t, me, m1.Submatrix(0, 2))
		})
		t.Run("A submatrix of a 4x4 matrix is a 3x3 matrix", func(t *testing.T) {
			m1 := NewMatrixFromTable(`
				+-----------------+
				| -6 | 1 |  1 | 6 |
				| -8 | 5 |  8 | 6 |
				| -1 | 0 |  8 | 2 |
				| -7 | 1 | -1 | 1 |
				+-----------------+
			`)
			me := NewMatrixFromTable(`
				+-------------+
				| -6 |  1 | 6 |
				| -8 |  8 | 6 |
				| -7 | -1 | 1 |
				+-------------+
			`)
			equalMatrix(t, me, m1.Submatrix(2, 1))
		})
	})
	t.Run("Minors", func(t *testing.T) {
		t.Run("Calculating a minor of a 3x3 matrix", func(t *testing.T) {
			m1 := NewMatrixFromTable(`
				+-------------+
				| 3 |  5 |  0 |
				| 2 | -1 | -7 |
				| 6 | -1 |  5 |
				+-------------+
			`)
			require.EqualValues(t, 25, m1.Minor(1, 0))
			m2 := m1.Submatrix(1, 0)
			require.EqualValues(t, 25, m2.Determinant())
		})
	})
	t.Run("Cofactors", func(t *testing.T) {
		t.Run("Calculating a cofactor of a 3x3 matrix", func(t *testing.T) {
			m1 := NewMatrixFromTable(`
				+-------------------+
				| 3.0 |  5.0 |  0.0 |
				| 2.0 | -1.0 | -7.0 |
				| 6.0 | -1.0 |  5.0 |
				+-------------------+
			`)
			require.EqualValues(t, -12, m1.Minor(0, 0))
			require.EqualValues(t, -12, m1.Cofactor(0, 0))
			require.EqualValues(t, 25, m1.Minor(1, 0))
			require.EqualValues(t, -25, m1.Cofactor(1, 0))
		})
	})
}

func BenchmarkMatrix4x4Determinant(b *testing.B) {
	m1 := NewMatrixFromTable(`
		+---------------------------+
		| -2.0 | -8.0 |  3.0 |  5.0 |
		| -3.0 |  1.0 |  7.0 |  3.0 |
		|  1.0 |  2.0 | -9.0 |  6.0 |
		| -6.0 |  7.0 |  7.0 | -9.0 |
		+---------------------------+
	`)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = m1.Determinant()
	}
}

func BenchmarkMatrix4x4Submatrix(b *testing.B) {
	m1 := NewMatrixFromTable(`
			+---------------+
			| 1 | 2 | 3 | 4 |
			| 5 | 6 | 7 | 8 |
			| 9 | 8 | 7 | 6 |
			| 5 | 4 | 3 | 2 |
			+---------------+
		`)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m1.Submatrix(1, 2)
	}
}

func BenchmarkMatrixMultiply(b *testing.B) {
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
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m1.Multiply(m2)
	}
}

func notEqualMatrix(t *testing.T, m1, m2 Matrix) {
	require.False(t, m1.Equal(m2))
	require.False(t, m2.Equal(m1))
}

func equalMatrix(t *testing.T, me, m1 Matrix) {
	require.True(t, me.Equal(m1), "me:\n%s\nwas not equal to\nm1:\n%s", me, m1)
	require.True(t, m1.Equal(me), "m1:\n%s\nwas not equal to\nme:\n%s", m1, me)
}
