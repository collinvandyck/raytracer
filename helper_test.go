package rt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func equalTuple(t *testing.T, t1, t2 Tuple4) {
	t.Helper()
	require.EqualValues(t, t1, t2)
	require.True(t, t1.equal(t2))
}

func equalVector(t *testing.T, v1, v2 Vector) {
	t.Helper()
	require.True(t, v1.Equal(v2), "Expected %s to equal %s", v2, v1)
}

func equalPoint(t *testing.T, p1, p2 Point) {
	t.Helper()
	require.True(t, p1.Equal(p2), "Expected %s to equal %s", p2, p1)
}

func notEqualMatrix(t *testing.T, m1, m2 Matrix) {
	require.False(t, m1.Equal(m2))
	require.False(t, m2.Equal(m1))
}

func equalMatrix(t *testing.T, me, m1 Matrix) {
	require.True(t, me.Equal(m1), "expected:\n%s\nactual:\n%s", me, m1)
	require.True(t, m1.Equal(me), "expected:\n%s\nactual:\n%s", m1, me)
}

func equalColor(t *testing.T, c1, c2 Color) {
	t.Helper()
	require.True(t, c1.Equal(c2))
}

func equalIntersection(t *testing.T, i1, i2 Intersection) {
	t.Helper()
	require.True(t, i1.Equal(i2), "Expected %v to equal %v", i2, i1)
}
