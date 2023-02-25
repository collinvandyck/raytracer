package raytracer

import (
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
		point := point(4, -4, 3)
		require.Equal(t, tuple4{4, -4, 3, 1}, point)
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
	t.Run("A vector creates a tuple with w=0", func(t *testing.T) {
		vector := vector(4, -4, 3)
		require.Equal(t, tuple4{4, -4, 3, 0}, vector)
	})
}

func TestTupleAdd(t *testing.T) {
	t.Run("Adding two tuples", func(t *testing.T) {
		t1 := tuple(3, -2, 5, 1)
		t2 := tuple(-2, 3, 1, 0)
		require.True(t, tuple(1, 1, 6, 1).Equal(t1.add(t2)))
	})
	t.Run("Adding two points is undefined", func(t *testing.T) {
		p1 := point(3, -2, 5)
		p2 := point(-2, 3, 1)
		re := p1.add(p2)
		require.False(t, re.isPoint())
		require.False(t, re.isVector())
	})
	t.Run("Adding two vectors is a vector", func(t *testing.T) {
		v1 := vector(3, -2, 5)
		v2 := vector(-2, 3, 1)
		re := v1.add(v2)
		require.Equal(t, vector(1, 1, 6), re)
	})
	t.Run("Adding vector and a point results in a point", func(t *testing.T) {
		p1 := point(3, -2, 5)
		v1 := vector(-2, 3, 1)
		re := p1.add(v1)
		require.Equal(t, point(1, 1, 6), re)
	})
}

func TestTupleSubtract(t *testing.T) {
	t.Run("Subtracing two tuples", func(t *testing.T) {
		t1 := tuple(3, -2, 5, 1)
		t2 := tuple(-2, 3, 1, 0)
		re := tuple(5, -5, 4, 1)
		require.Equal(t, re, t1.subtract(t2))
	})
	t.Run("Subtracting two points is a vector", func(t *testing.T) {
		p1 := point(3, -2, 5)
		p2 := point(-2, 3, 1)
		re := p1.subtract(p2)
		require.Equal(t, vector(5, -5, 4), re)
	})
	t.Run("Subtracting two vectors is a vector", func(t *testing.T) {
		v1 := vector(3, -2, 5)
		v2 := vector(-2, 3, 1)
		re := v1.subtract(v2)
		require.Equal(t, vector(5, -5, 4), re)
	})
	t.Run("Subtracting a vector from a point is a point", func(t *testing.T) {
		p1 := point(3, 2, 1)
		v1 := vector(5, 6, 7)
		re := point(-2, -4, -6)
		require.Equal(t, re, p1.subtract(v1))
	})
}
