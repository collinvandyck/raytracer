package rt

var emptyMatrix Matrix

var MatrixIdentity4x4 = NewMatrixFromValues([][]Value{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
})

var Origin = NewPoint(0, 0, 0)

var defaultMaterial = DefaultMaterial()
