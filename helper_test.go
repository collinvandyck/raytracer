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
	t.Helper()
	require.False(t, m1.Equal(m2))
	require.False(t, m2.Equal(m1))
}

func equalMatrix(t *testing.T, me, m1 Matrix) {
	t.Helper()
	require.True(t, me.Equal(m1), "expected:\n%s\nactual:\n%s", me, m1)
	require.True(t, m1.Equal(me), "expected:\n%s\nactual:\n%s", m1, me)
}

func equalColor(t *testing.T, c1, c2 Color) {
	t.Helper()
	require.True(t, c1.Equal(c2), "Expected %s to equal %s", c2, c1)
}

func equalValue(t *testing.T, v1, v2 Value) {
	t.Helper()
	require.True(t, floatsEqual(v1, v2), "Expected %v to equal %v", v2, v1)
}

var _ = equalValueSlice

func equalValueSlice(t *testing.T, v1, v2 []Value) {
	t.Helper()
	require.Len(t, v2, len(v1), "Expected %v to equal %v", v2, v1)
	for i := 0; i < len(v1); i++ {
		require.True(t, floatsEqual(v1[i], v2[i]), "Expected %v to equal %v", v2, v1)
	}
}

func equalShape(t *testing.T, s1, s2 Shape) {
	t.Helper()
	require.EqualValues(t, s2, s1)
}

var _ = equalIntersection

func equalIntersection(t *testing.T, i1, i2 Intersection) {
	t.Helper()
	require.True(t, i1.Equal(i2), "Expected %s to equal %s", i2, i1)
}

func equalIntersections(t *testing.T, i1, i2 Intersections) {
	t.Helper()
	require.True(t, i1.Equal(i2), "Expected %s to equal %s", i2, i1)
}

func equalMaterial(t *testing.T, m1, m2 Material) {
	t.Helper()
	require.True(t, m1.Equal(m2), "Expected %s to equal %s", m2, m1)
}

func equalLight(t *testing.T, l1, l2 *PointLight) {
	t.Helper()
	require.True(t, l1.Equal(l2), "Expected %s to equal %s", l2, l1)
}
