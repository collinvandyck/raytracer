package rt

type Canvas struct {
	colors [][]Color
}

func NewCanvas(width int, height int) *Canvas {
	if width <= 0 || height <= 0 {
		panic("dimensions must be positive")
	}
	res := &Canvas{colors: make([][]Color, height)}
	for i := range res.colors {
		res.colors[i] = make([]Color, width)
	}
	return res
}

func (c *Canvas) WritePixel(x, y int, color Color) {
	c.colors[y][x] = color
}

func (c *Canvas) PixelAt(x, y int) Color {
	return c.colors[y][x]
}

func (c *Canvas) Width() int {
	return len(c.colors[0])
}

func (c *Canvas) Height() int {
	return len(c.colors)
}
