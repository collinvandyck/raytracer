package raytracer

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
		p1 := newPoint(4, -4, 3)
		require.Equal(t, point(tuple4{4, -4, 3, 1}), p1)
	})
	t.Run("A vector creates a tuple with w=0", func(t *testing.T) {
		v1 := newVector(4, -4, 3)
		require.Equal(t, vector(tuple4{4, -4, 3, 0}), v1)
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
		v1 := newVector(3, -2, 5)
		v2 := newVector(-2, 3, 1)
		re := newVector(1, 1, 6)
		equalVector(t, re, v1.addVector(v2))

	})
	t.Run("Adding vector and a point results in a point", func(t *testing.T) {
		p1 := newPoint(3, -2, 5)
		v1 := newVector(-2, 3, 1)
		re := newPoint(1, 1, 6)
		equalPoint(t, re, p1.addVector(v1))
		equalPoint(t, re, v1.addPoint(p1))
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
		p1 := newPoint(3, -2, 5)
		p2 := newPoint(-2, 3, 1)
		re := newVector(5, -5, 4)
		equalVector(t, re, p1.subtractPoint(p2))
	})
	t.Run("Subtracting two vectors is a vector", func(t *testing.T) {
		v1 := newVector(3, -2, 5)
		v2 := newVector(-2, 3, 1)
		re := newVector(5, -5, 4)
		equalVector(t, re, v1.subtractVector(v2))
	})
	t.Run("Subtracting a vector from a point is a point", func(t *testing.T) {
		p1 := newPoint(3, 2, 1)
		v1 := newVector(5, 6, 7)
		re := newPoint(-2, -4, -6)
		equalPoint(t, re, p1.subtractVector(v1))
	})
	t.Run("Subtracing a vector from the zero vector", func(t *testing.T) {
		v1 := newVector(0, 0, 0)
		v2 := newVector(1, -2, -3)
		re := newVector(-1, 2, 3)
		equalVector(t, re, v1.subtractVector(v2))
	})
	t.Run("Negating a vector", func(t *testing.T) {
		v1 := newVector(1, -2, -3)
		re := newVector(-1, 2, 3)
		equalVector(t, re, v1.negate())
	})
}

func TestTupleMultiply(t *testing.T) {
	t.Run("Multiplying a tuple by a scalar", func(t *testing.T) {
		const factor float = 3.5
		t1 := tuple(1, -2, 3, -4)
		re := tuple(1*factor, -2*factor, 3*factor, -4*factor)
		equalTuple(t, re, t1.multiply(factor))
	})
	t.Run("Multiplying a tuple by a fraction", func(t *testing.T) {
		const factor float = 0.5
		t1 := tuple(1, -2, 3, -4)
		re := tuple(1*factor, -2*factor, 3*factor, -4*factor)
		equalTuple(t, re, t1.multiply(factor))
	})
}

func TestTupleDivide(t *testing.T) {
	t.Run("Dividing a tuple by a scalar", func(t *testing.T) {
		const factor float = 2
		t1 := tuple(1, -2, 3, -4)
		re := tuple(1/factor, -2/factor, 3/factor, -4/factor)
		equalTuple(t, re, t1.divide(factor))
	})
}

func TestVectorMagnitude(t *testing.T) {
	t.Run("Computing the magnitude of vector(1, 0, 0)", func(t *testing.T) {
		v1 := newVector(1, 0, 0)
		require.EqualValues(t, 1, v1.magnitude())
	})
	t.Run("Computing the magnitude of vector(0, 1, 0)", func(t *testing.T) {
		v1 := newVector(0, 1, 0)
		require.EqualValues(t, 1, v1.magnitude())
	})
	t.Run("Computing the magnitude of vector(1, 2, 3)", func(t *testing.T) {
		v1 := newVector(1, 2, 3)
		require.EqualValues(t, math.Sqrt(14), v1.magnitude())
	})
	t.Run("Computing the magnitude of vector(-1, -2, -3)", func(t *testing.T) {
		v1 := newVector(1, 2, 3)
		require.EqualValues(t, math.Sqrt(14), v1.magnitude())
	})
}

func TestVectorNormalize(t *testing.T) {
	t.Run("Normalizing vector(4, 0, 0) gives (1, 0, 0)", func(t *testing.T) {
		v1 := newVector(4, 0, 0)
		re := newVector(1, 0, 0)
		equalVector(t, re, v1.normalize())
	})
	t.Run("Normalizing vector(1, 2, 3)", func(t *testing.T) {
		v1 := newVector(4, 0, 0)
		ex := 1.0 / math.Sqrt(14)
		ey := 2.0 / math.Sqrt(14)
		ez := 3.0 / math.Sqrt(14)
		re := newVector(ex, ey, ez)
		equalVector(t, re, v1.normalize())
	})
	t.Run("The magnitude of a normalized vector", func(t *testing.T) {
		v1 := newVector(1, 2, 3)
		n1 := v1.normalize()
		e1 := n1.magnitude()
		require.Equal(t, 1, e1)
	})
}

func equalTuple(t *testing.T, t1, t2 tuple4) {
	require.EqualValues(t, t1, t2)
}

func equalVector(t *testing.T, v1, v2 vector) {
	require.EqualValues(t, v1, v2)
}

func equalPoint(t *testing.T, p1, p2 point) {
	require.EqualValues(t, p1, p2)
}
