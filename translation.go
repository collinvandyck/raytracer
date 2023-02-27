package rt

type Translator Matrix

func Translation(x, y, z float) Translator {
	m := NewMatrix(4, 4)
	return Translator(m)
}

func (t Translator) MultiplyPoint(point Point) Point {
	return point
}
