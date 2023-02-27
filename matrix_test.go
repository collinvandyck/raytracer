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
	t.Run("Testing an invertible matrix for invertibility", func(t *testing.T) {
		m1 := NewMatrixFromTable(`
			+-------------------------+
			| 6.0 |  4.0 | 4.0 |  4.0 |
			| 5.0 |  5.0 | 7.0 |  6.0 |
			| 4.0 | -9.0 | 3.0 | -7.0 |
			| 9.0 |  1.0 | 7.0 | -6.0 |
			+-------------------------+
		`)
		require.EqualValues(t, -2120, m1.Determinant())
		require.True(t, m1.IsInvertible())
	})
	t.Run("Testing a noninvertible matrix for invertibility", func(t *testing.T) {
		m1 := NewMatrixFromTable(`
			+---------------------------+
			| -4.0 |  2.0 | -2.0 | -3.0 |
			|  9.0 |  6.0 |  2.0 |  6.0 |
			|  0.0 | -5.0 |  1.0 | -5.0 |
			|  0.0 |  0.0 |  0.0 |  0.0 |
			+---------------------------+
		`)
		require.EqualValues(t, 0, m1.Determinant())
		require.False(t, m1.IsInvertible())
	})
	t.Run("Calculating the inverse of a matrix", func(t *testing.T) {
		a1 := NewMatrixFromTable(`
			+---------------------------+
			| -5.0 |  2.0 |  6.0 | -8.0 |
			|  1.0 | -5.0 |  1.0 |  8.0 |
			|  7.0 |  7.0 | -6.0 | -7.0 |
			|  1.0 | -3.0 |  7.0 |  4.0 |
			+---------------------------+
		`)
		b1 := a1.Inverse()
		require.EqualValues(t, 532, a1.Determinant())
		require.EqualValues(t, -160, a1.Cofactor(2, 3))
		require.EqualValues(t, -160.0/532.0, b1.Get(3, 2))
		require.EqualValues(t, 105, a1.Cofactor(3, 2))
		require.EqualValues(t, 105.0/532.0, b1.Get(2, 3))
		ie := NewMatrixFromTable(`
			+-------------------------------------------+
			|  0.21805 |  0.45113 |   0.2406 | -0.04511 |
			| -0.80827 | -1.45677 | -0.44361 |  0.52068 |
			| -0.07895 | -0.22368 | -0.05263 |  0.19737 |
			| -0.52256 | -0.81391 | -0.30075 |  0.30639 |
			+-------------------------------------------+
		`)
		equalMatrix(t, ie, b1)
	})
	t.Run("Calculating the inverse of another matrix", func(t *testing.T) {
		a1 := NewMatrixFromTable(`
			+-------------------+
			|  8 | -5 |  9 |  2 |
			|  7 |  5 |  6 |  1 |
			| -6 |  0 |  9 |  6 |
			| -3 |  0 | -9 | -4 |
			+-------------------+
		`)
		b1 := a1.Inverse()
		ie := NewMatrixFromTable(`
			+-------------------------------------------+
			| -0.15385 | -0.15385 | -0.28205 | -0.53846 |
			| -0.07692 |  0.12308 |  0.02564 |  0.03077 |
			|  0.35897 |  0.35897 |   0.4359 |  0.92308 |
			| -0.69231 | -0.69231 | -0.76923 | -1.92308 |
			+-------------------------------------------+
		`)
		equalMatrix(t, ie, b1)
	})
	t.Run("Calculating the inverse of a third matrix", func(t *testing.T) {
		a1 := NewMatrixFromTable(`
			+-------------------+
			|  9 |  3 |  0 |  9 |
			| -5 | -2 | -6 | -3 |
			| -4 |  9 |  6 |  4 |
			| -7 |  6 |  6 |  2 |
			+-------------------+
		`)
		b1 := a1.Inverse()
		ie := NewMatrixFromTable(`
			+-------------------------------------------+
			| -0.04074 | -0.07778 |  0.14444 | -0.22222 |
			| -0.07778 |  0.03333 |  0.36667 | -0.33333 |
			| -0.02901 |  -0.1463 | -0.10926 |  0.12963 |
			|  0.17778 |  0.06667 | -0.26667 |  0.33333 |
			+-------------------------------------------+		
		`)
		equalMatrix(t, ie, b1)
	})
	t.Run("Multiplying a product by its inverse", func(t *testing.T) {
		a1 := NewMatrixFromTable(`
			+-------------------+
			|  3 | -9 |  7 |  3 |
			|  3 | -8 |  2 | -9 |
			| -4 |  4 |  4 |  1 |
			| -6 |  5 | -1 |  1 |
			+-------------------+
		`)
		b1 := NewMatrixFromTable(`
			+----------------+
			| 8 |  2 | 2 | 2 |
			| 3 | -1 | 7 | 0 |
			| 7 |  0 | 5 | 4 |
			| 6 | -2 | 0 | 5 |
			+----------------+
		`)
		c1 := a1.Multiply(b1)
		equalMatrix(t, a1, c1.Multiply(b1.Inverse()))
	})
	t.Run("Inverting the identity matrix", func(t *testing.T) {
		a1 := MatrixIdentity4x4.Inverse()
		equalMatrix(t, MatrixIdentity4x4, a1)
	})
	t.Run("Multiplying a matrix by its inverse", func(t *testing.T) {
		a1 := NewMatrixFromTable(`
			+------------------+
			|  3 | -8 |  2 | 1 |
			| -4 |  4 |  4 | 2 |
			| -6 |  5 | -1 | 3 |
			| -5 |  4 | -2 | 2 |
			+------------------+
		`)
		i1 := a1.Inverse()
		r1 := a1.Multiply(i1)
		equalMatrix(t, MatrixIdentity4x4, r1)
	})
	t.Run("Comparing inverting transpose vs transposing inverted", func(t *testing.T) {
		a1 := NewMatrixFromTable(`
			+------------------+
			|  3 | -8 |  2 | 1 |
			| -4 |  4 |  4 | 2 |
			| -6 |  5 | -1 | 3 |
			| -5 |  4 | -2 | 2 |
			+------------------+
		`)
		r1 := a1.Transpose().Inverse()
		r2 := a1.Inverse().Transpose()
		equalMatrix(t, r1, r2)
	})
	t.Run("Multiplying identity and non-identity by a tuple", func(t *testing.T) {
		t1 := Tuple4{1, 2, 3, 4}
		m1 := MatrixIdentity4x4
		equalTuple(t, t1, m1.MultiplyTuple4(t1))

		i2 := NewMatrixFromTable(`
			+---------------+
			| 1 | 0 | 0 | 9 |
			| 0 | 1 | 0 | 8 |
			| 0 | 0 | 2 | 7 |
			| 0 | 0 | 0 | 1 |
			+---------------+
		`)
		equalTuple(t, NewTuple(37, 34, 34, 4), i2.MultiplyTuple4(t1))
	})
}

func notEqualMatrix(t *testing.T, m1, m2 Matrix) {
	require.False(t, m1.Equal(m2))
	require.False(t, m2.Equal(m1))
}

func equalMatrix(t *testing.T, me, m1 Matrix) {
	require.True(t, me.Equal(m1), "expected:\n%s\nactual:\n%s", me, m1)
	require.True(t, m1.Equal(me), "expected:\n%s\nactual:\n%s", m1, me)
}
