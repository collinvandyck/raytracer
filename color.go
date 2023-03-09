package rt

import "fmt"

var black = NewColor(0, 0, 0)

type Color Tuple4

func NewColor(r, g, b Value) Color {
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

func (c Color) MultiplyBy(val Value) Color {
	return Color(Tuple4(c).multiplyBy(val))
}

func (c Color) Red() Value {
	return c.x
}

func (c Color) Green() Value {
	return c.y
}

func (c Color) Blue() Value {
	return c.z
}

func (c Color) Equal(o Color) bool {
	return Tuple4(c).equal(Tuple4(o))
}

func (c Color) String() string {
	return fmt.Sprintf("Color(%s,%s,%s)",
		formatFloat(c.x), formatFloat(c.y), formatFloat(c.z))
}
