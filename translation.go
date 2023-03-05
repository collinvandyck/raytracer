package rt

import "math"

// Translation applies a direction
func Translation(x, y, z Value) Matrix {
	m := NewMatrix(4, 4)
	m.Set(0, 0, 1)
	m.Set(1, 1, 1)
	m.Set(2, 2, 1)
	m.Set(3, 3, 1)
	m.Set(0, 3, x)
	m.Set(1, 3, y)
	m.Set(2, 3, z)
	return m
}

func Scaling(x, y, z Value) Matrix {
	m := NewMatrix(4, 4)
	m.Set(0, 0, x)
	m.Set(1, 1, y)
	m.Set(2, 2, z)
	m.Set(3, 3, 1)
	return m
}

func RotationX(rad Value) Matrix {
	m := NewMatrix(4, 4)
	m.Set(0, 0, 1)
	m.Set(1, 1, math.Cos(rad))
	m.Set(1, 2, -math.Sin(rad))
	m.Set(2, 1, math.Sin(rad))
	m.Set(2, 2, math.Cos(rad))
	m.Set(3, 3, 1)
	return m
}

func RotationY(rad Value) Matrix {
	m := NewMatrix(4, 4)
	m.Set(0, 0, math.Cos(rad))
	m.Set(0, 2, math.Sin(rad))
	m.Set(1, 1, 1)
	m.Set(2, 0, -math.Sin(rad))
	m.Set(2, 2, math.Cos(rad))
	m.Set(3, 3, 1)
	return m
}

func RotationZ(rad Value) Matrix {
	m := NewMatrix(4, 4)
	m.Set(0, 0, math.Cos(rad))
	m.Set(0, 1, -math.Sin(rad))
	m.Set(1, 0, math.Sin(rad))
	m.Set(1, 1, math.Cos(rad))
	m.Set(2, 2, 1)
	m.Set(3, 3, 1)
	return m
}

func Shearing(xy, xz, yx, yz, zx, zy Value) Matrix {
	m := NewMatrix(4, 4)
	m.Set(0, 0, 1)
	m.Set(1, 1, 1)
	m.Set(2, 2, 1)
	m.Set(3, 3, 1)
	m.Set(0, 1, xy)
	m.Set(0, 2, xz)
	m.Set(1, 0, yx)
	m.Set(1, 2, yz)
	m.Set(2, 0, zx)
	m.Set(2, 1, zy)
	return m
}
