package rt

import "fmt"

type Canvas struct {
	pixels    [][]Color
	pointSize int
}

func NewCanvas(width int, height int) *Canvas {
	if width <= 0 || height <= 0 {
		panic("dimensions must be positive")
	}
	res := &Canvas{
		pixels:    make([][]Color, height),
		pointSize: 1,
	}
	for i := range res.pixels {
		res.pixels[i] = make([]Color, width)
	}
	return res
}

// Returns a unique set of colors for this canvas
func (c *Canvas) Colors() Colors {
	res := make(Colors)
	for x := 0; x < c.Width(); x++ {
		for y := 0; y < c.Height(); y++ {
			color := c.PixelAt(x, y)
			res[color] = true
		}
	}
	return res
}

func (c *Canvas) SetPointSize(val int) {
	if val < 1 {
		panic(fmt.Sprintf("invalid point size: %v", val))
	}
	c.pointSize = val
}

func (c *Canvas) Fill(color Color) {
	for ri := 0; ri < c.Height(); ri++ {
		for ci := 0; ci < c.Width(); ci++ {
			c.write(ci, ri, color)
		}
	}
}

func (c *Canvas) WritePoint(point Point, color Color) {
	x := int(point.x)
	y := int(point.y)
	c.WritePixel(x, y, color)
}

func (c *Canvas) WritePixel(x, y int, color Color) {
	for xi := x; xi < x+c.pointSize; xi++ {
		for yi := y; yi < y+c.pointSize; yi++ {
			c.write(xi, yi, color)
		}
	}
}

func (c *Canvas) write(x, y int, color Color) {
	if y < 0 || y >= len(c.pixels) {
		return
	}
	if x < 0 || x >= len(c.pixels[y]) {
		return
	}
	c.pixels[y][x] = color
}

func (c *Canvas) PixelAt(x, y int) Color {
	return c.pixels[y][x]
}

func (c *Canvas) Width() int {
	return len(c.pixels[0])
}

func (c *Canvas) Height() int {
	return len(c.pixels)
}
