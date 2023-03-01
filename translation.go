package rt

import "math"

func Translation(x, y, z float) Translator {
	m := NewMatrix(4, 4)
	m.Set(0, 0, 1)
	m.Set(1, 1, 1)
	m.Set(2, 2, 1)
	m.Set(3, 3, 1)
	m.Set(0, 3, x)
	m.Set(1, 3, y)
	m.Set(2, 3, z)
	return Translator(m)
}

func Scaling(x, y, z float) Translator {
	m := NewMatrix(4, 4)
	m.Set(0, 0, x)
	m.Set(1, 1, y)
	m.Set(2, 2, z)
	m.Set(3, 3, 1)
	return Translator(m)
}

func RotationX(rad float) Translator {
	m := NewMatrix(4, 4)
	m.Set(0, 0, 1)
	m.Set(1, 1, math.Cos(rad))
	m.Set(1, 2, -math.Sin(rad))
	m.Set(2, 1, math.Sin(rad))
	m.Set(2, 2, math.Cos(rad))
	m.Set(3, 3, 1)
	return Translator(m)
}

func RotationY(rad float) Translator {
	m := NewMatrix(4, 4)
	m.Set(0, 0, math.Cos(rad))
	m.Set(0, 2, math.Sin(rad))
	m.Set(1, 1, 1)
	m.Set(2, 0, -math.Sin(rad))
	m.Set(2, 2, math.Cos(rad))
	m.Set(3, 3, 1)
	return Translator(m)
}

type Translator Matrix

func (t Translator) MultiplyPoint(point Point) Point {
	return Point(Matrix(t).MultiplyTuple4(Tuple4(point)))
}

// todo: optimize by just returning the vector?
func (t Translator) MultiplyVector(vector Vector) Vector {
	return Vector(Matrix(t).MultiplyTuple4(Tuple4(vector)))
}

func (t Translator) Inverse() Translator {
	return Translator(Matrix(t).Inverse())
}
