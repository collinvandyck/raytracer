package rt

type Translator Matrix

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

func (t Translator) MultiplyPoint(point Point) Point {
	return Point(Matrix(t).MultiplyTuple4(Tuple4(point)))
}

func (t Translator) MultiplyVector(vector Vector) Vector {
	return Vector(Matrix(t).MultiplyTuple4(Tuple4(vector)))
}

func (t Translator) Inverse() Translator {
	return Translator(Matrix(t).Inverse())
}
