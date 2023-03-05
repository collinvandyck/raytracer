package rt

type Color Tuple4

func NewColor(r, g, b value) Color {
	return Color(NewTuple(r, g, b, 0))
}

func (c Color) Add(o Color) Color {
	return Color(Tuple4(c).add(Tuple4(o)))
}

func (c Color) Subtract(o Color) Color {
	return Color(Tuple4(c).subtract(Tuple4(o)))
}

func (c Color) Multiply(o Color) Color {
	return Color(Tuple4(c).multiply(Tuple4(o)))
}

func (c Color) MultiplyBy(val value) Color {
	return Color(Tuple4(c).multiplyBy(val))
}

func (c Color) Red() value {
	return c.x
}

func (c Color) Green() value {
	return c.y
}

func (c Color) Blue() value {
	return c.z
}

func (c Color) Equal(o Color) bool {
	return Tuple4(c).equal(Tuple4(o))
}
