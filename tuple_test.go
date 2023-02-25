package rt

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTuple(t *testing.T) {
	t.Run("A tuple with w=1 is a point", func(t *testing.T) {
		tup := tuple(4.3, -4.2, 3.1, 1.0)
		require.Equal(t, tup.x, 4.3)
		require.Equal(t, tup.y, -4.2)
		require.Equal(t, tup.z, 3.1)
		require.Equal(t, tup.w, 1.0)
		require.True(t, tup.isPoint())
		require.False(t, tup.isVector())
	})
	t.Run("A point creates a tuple with w=1", func(t *testing.T) {
		p1 := NewPoint(4, -4, 3)
		require.Equal(t, Point(tuple4{4, -4, 3, 1}), p1)
	})
	t.Run("A vector creates a tuple with w=0", func(t *testing.T) {
		v1 := NewVector(4, -4, 3)
		require.Equal(t, Vector(tuple4{4, -4, 3, 0}), v1)
	})
	t.Run("A tuple with w=0 is a vector", func(t *testing.T) {
		tup := tuple(4.3, -4.2, 3.1, 0.0)
		require.Equal(t, tup.x, 4.3)
		require.Equal(t, tup.y, -4.2)
		require.Equal(t, tup.z, 3.1)
		require.Equal(t, tup.w, 0.0)
		require.False(t, tup.isPoint())
		require.True(t, tup.isVector())
	})
}

func TestTupleAdd(t *testing.T) {
	t.Run("Adding two tuples", func(t *testing.T) {
		t1 := tuple(3, -2, 5, 1)
		t2 := tuple(-2, 3, 1, 0)
		re := tuple(1, 1, 6, 1)
		equalTuple(t, re, t1.add(t2))
	})
	t.Run("Adding two vectors is a vector", func(t *testing.T) {
		v1 := NewVector(3, -2, 5)
		v2 := NewVector(-2, 3, 1)
		re := NewVector(1, 1, 6)
		equalVector(t, re, v1.AddVector(v2))

	})
	t.Run("Adding vector and a point results in a point", func(t *testing.T) {
		p1 := NewPoint(3, -2, 5)
		v1 := NewVector(-2, 3, 1)
		re := NewPoint(1, 1, 6)
		equalPoint(t, re, p1.AddVector(v1))
		equalPoint(t, re, v1.AddPoint(p1))
	})
}

func TestTupleSubtract(t *testing.T) {
	t.Run("Subtracing two tuples", func(t *testing.T) {
		t1 := tuple(3, -2, 5, 1)
		t2 := tuple(-2, 3, 1, 0)
		re := tuple(5, -5, 4, 1)
		equalTuple(t, re, t1.subtract(t2))
	})
	t.Run("Subtracting two points is a vector", func(t *testing.T) {
		p1 := NewPoint(3, -2, 5)
		p2 := NewPoint(-2, 3, 1)
		re := NewVector(5, -5, 4)
		equalVector(t, re, p1.SubtractPoint(p2))
	})
	t.Run("Subtracting two vectors is a vector", func(t *testing.T) {
		v1 := NewVector(3, -2, 5)
		v2 := NewVector(-2, 3, 1)
		re := NewVector(5, -5, 4)
		equalVector(t, re, v1.SubtractVector(v2))
	})
	t.Run("Subtracting a vector from a point is a point", func(t *testing.T) {
		p1 := NewPoint(3, 2, 1)
		v1 := NewVector(5, 6, 7)
		re := NewPoint(-2, -4, -6)
		equalPoint(t, re, p1.SubtractVector(v1))
	})
	t.Run("Subtracing a vector from the zero vector", func(t *testing.T) {
		v1 := NewVector(0, 0, 0)
		v2 := NewVector(1, -2, -3)
		re := NewVector(-1, 2, 3)
		equalVector(t, re, v1.SubtractVector(v2))
	})
	t.Run("Negating a vector", func(t *testing.T) {
		v1 := NewVector(1, -2, -3)
		re := NewVector(-1, 2, 3)
		equalVector(t, re, v1.Negate())
	})
}

func TestTupleMultiply(t *testing.T) {
	t.Run("Multiplying a tuple by a scalar", func(t *testing.T) {
		const factor float = 3.5
		t1 := tuple(1, -2, 3, -4)
		re := tuple(1*factor, -2*factor, 3*factor, -4*factor)
		equalTuple(t, re, t1.multiplyBy(factor))
	})
	t.Run("Multiplying a tuple by a fraction", func(t *testing.T) {
		const factor float = 0.5
		t1 := tuple(1, -2, 3, -4)
		re := tuple(1*factor, -2*factor, 3*factor, -4*factor)
		equalTuple(t, re, t1.multiplyBy(factor))
	})
}

func TestTupleDivide(t *testing.T) {
	t.Run("Dividing a tuple by a scalar", func(t *testing.T) {
		const factor float = 2
		t1 := tuple(1, -2, 3, -4)
		re := tuple(1/factor, -2/factor, 3/factor, -4/factor)
		equalTuple(t, re, t1.divideBy(factor))
	})
}

func TestVectorMagnitude(t *testing.T) {
	t.Run("Computing the magnitude of vector(1, 0, 0)", func(t *testing.T) {
		v1 := NewVector(1, 0, 0)
		require.EqualValues(t, 1, v1.Magnitude())
	})
	t.Run("Computing the magnitude of vector(0, 1, 0)", func(t *testing.T) {
		v1 := NewVector(0, 1, 0)
		require.EqualValues(t, 1, v1.Magnitude())
	})
	t.Run("Computing the magnitude of vector(1, 2, 3)", func(t *testing.T) {
		v1 := NewVector(1, 2, 3)
		require.EqualValues(t, math.Sqrt(14), v1.Magnitude())
	})
	t.Run("Computing the magnitude of vector(-1, -2, -3)", func(t *testing.T) {
		v1 := NewVector(1, 2, 3)
		require.EqualValues(t, math.Sqrt(14), v1.Magnitude())
	})
}

func TestVectorNormalize(t *testing.T) {
	t.Run("Normalizing vector(4, 0, 0) gives (1, 0, 0)", func(t *testing.T) {
		v1 := NewVector(4, 0, 0)
		re := NewVector(1, 0, 0)
		equalVector(t, re, v1.Normalize())
	})
	t.Run("Normalizing vector(1, 2, 3)", func(t *testing.T) {
		v1 := NewVector(1, 2, 3)
		ex := 1.0 / math.Sqrt(14)
		ey := 2.0 / math.Sqrt(14)
		ez := 3.0 / math.Sqrt(14)
		re := NewVector(ex, ey, ez)
		equalVector(t, re, v1.Normalize())
	})
	t.Run("The magnitude of a normalized vector", func(t *testing.T) {
		v1 := NewVector(1, 2, 3)
		require.EqualValues(t, 1, v1.Normalize().Magnitude())
	})
}

func TestVectorDotProduct(t *testing.T) {
	t.Run("The dot product of two tuples", func(t *testing.T) {
		v1 := NewVector(1, 2, 3)
		v2 := NewVector(2, 3, 4)
		require.EqualValues(t, 20, v1.Dot(v2))
	})
}

func TestVectorCrossProduct(t *testing.T) {
	t.Run("The cross product of two vectors", func(t *testing.T) {
		v1 := NewVector(1, 2, 3)
		v2 := NewVector(2, 3, 4)
		equalVector(t, NewVector(-1, 2, -1), v1.Cross(v2))
		equalVector(t, NewVector(1, -2, 1), v2.Cross(v1))
	})
}

func equalTuple(t *testing.T, t1, t2 tuple4) {
	t.Helper()
	require.EqualValues(t, t1, t2)
	require.True(t, t1.equal(t2))
}

func equalVector(t *testing.T, v1, v2 Vector) {
	t.Helper()
	require.EqualValues(t, v1, v2)
	require.True(t, v1.Equal(v2))
}

func equalPoint(t *testing.T, p1, p2 Point) {
	t.Helper()
	require.EqualValues(t, p1, p2)
	require.True(t, p1.Equal(p2))
}
