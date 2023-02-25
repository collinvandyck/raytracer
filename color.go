package rt

type Color tuple4

func NewColor(r, g, b float) Color {
	return Color(tuple(r, g, b, 0))
}

func (c Color) Add(o Color) Color {
	return c
}

func (c Color) Subtract(o Color) Color {
	return c
}

func (c Color) MultiplyBy(val float) Color {
	return c
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
