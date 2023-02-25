package rt

type Color tuple4

func NewColor(r, g, b float) Color {
	return Color(tuple(r, g, b, 0))
}

func (c Color) Add(o Color) Color {
	return Color(tuple4(c).add(tuple4(o)))
}

func (c Color) Subtract(o Color) Color {
	return Color(tuple4(c).subtract(tuple4(o)))
}

func (c Color) MultiplyBy(val float) Color {
	return Color(tuple4(c).multiplyBy(val))
}

func (c Color) Red() float {
	return c.x
}

func (c Color) Green() float {
	return c.y
}

func (c Color) Blue() float {
	return c.z
}

func (c Color) Equal(o Color) bool {
	return tuple4(c).equal(tuple4(o))
}
