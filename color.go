package rt

import (
	"fmt"
	"image/color"
)

var black = NewColor(0, 0, 0)

type Colors map[Color]bool

func (c Colors) Add(color Color) {
	c[color] = true
}

func (c Colors) AddAll(o Colors) {
	for oc := range o {
		c.Add(oc)
	}
}

func (c Colors) ToPallete() color.Palette {
	scale := func(v Value) uint8 {
		v *= 255
		if v > 255 {
			v = 255
		}
		return uint8(v)
	}
	res := make(color.Palette, 0, len(c))
	for cc := range c {
		pc := color.RGBA{
			R: scale(cc.Red()),
			G: scale(cc.Green()),
			B: scale(cc.Blue()),
			A: 255,
		}
		res = append(res, pc)
	}
	return res
}

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
