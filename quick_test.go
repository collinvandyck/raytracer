package rt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuickSubmatrix(t *testing.T) {
	m1 := NewMatrixFromTable(`
		+---------------------------+
		| -2.0 | -8.0 |  3.0 |  5.0 |
		| -3.0 |  1.0 |  7.0 |  3.0 |
		|  1.0 |  2.0 | -9.0 |  6.0 |
		| -6.0 |  7.0 |  7.0 | -9.0 |
		+---------------------------+
	`)
	m2 := m1.Submatrix(0, 0)
	equalMatrix(t, NewMatrixFromTable(`
		+--------------------+
		|  1.0 |  7.0 |  3.0 |
		|  2.0 | -9.0 |  6.0 |
		|  7.0 |  7.0 | -9.0 |
		+--------------------+
	`), m2)
	m3 := m2.Multiply(m2)
	equalMatrix(t, NewMatrixFromTable(`
		+--------------------+
		|  36  |  -35 | 18   |
		|  26  | 137  | -102 |
		|  -42 | -77  | 144  |
		+--------------------+
	`), m3)
	m4 := m2.Submatrix(0, 0)
	equalMatrix(t, NewMatrixFromTable(`
		+---------+
		| -9 |  6 |
		|  7 | -9 |
		+---------+
	`), m4)
	m5 := m4.Multiply(m4)
	equalMatrix(t, NewMatrixFromTable(`
		+-------------+
		| 123  | -108 |
		| -126 |  123 |
		+-------------+ `), m5)
}

func TestQuickCofactor(t *testing.T) {
	m1 := NewMatrixFromTable(`
		+---------------------------+
		| -2.0 | -8.0 |  3.0 |  5.0 |
		| -3.0 |  1.0 |  7.0 |  3.0 |
		|  1.0 |  2.0 | -9.0 |  6.0 |
		| -6.0 |  7.0 |  7.0 | -9.0 |
		+---------------------------+
	`)
	m1.SetDebug(true)

	require.EqualValues(t, 690, m1.Cofactor(0, 0))
	require.EqualValues(t, 447, m1.Cofactor(0, 1))
	require.EqualValues(t, 210, m1.Cofactor(0, 2))
	require.EqualValues(t, 51, m1.Cofactor(0, 3))
	require.EqualValues(t, -4071, m1.Determinant())
}
