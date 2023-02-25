package rt

type Canvas struct {
	pixels [][]Color
}

func NewCanvas(width int, height int) Canvas {
	if width <= 0 || height <= 0 {
		panic("dimensions must be positive")
	}
	res := Canvas{pixels: make([][]Color, height)}
	for i := range res.pixels {
		res.pixels[i] = make([]Color, width)
	}
	return res
}

func (c *Canvas) Fill(color Color) {
	for ri := 0; ri < c.Height(); ri++ {
		for ci := 0; ci < c.Width(); ci++ {
			c.WritePixel(ci, ri, color)
		}
	}
}

func (c *Canvas) WritePixel(x, y int, color Color) {
	if y < 0 || y >= len(c.pixels) {
		return
	}
	if x < 0 || x >= len(c.pixels[y]) {
		return
	}
	c.pixels[y][x] = color
}

func (c Canvas) PixelAt(x, y int) Color {
	return c.pixels[y][x]
}

func (c Canvas) Width() int {
	return len(c.pixels[0])
}

func (c Canvas) Height() int {
	return len(c.pixels)
}
