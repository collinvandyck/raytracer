package rt

type Canvas struct {
	color [][]Color
}

func NewCanvas(width int, height int) *Canvas {
	if width <= 0 || height <= 0 {
		panic("dimensions must be positive")
	}
	res := &Canvas{color: make([][]Color, height)}
	for i := range res.color {
		res.color[i] = make([]Color, width)
	}
	return res
}

func (c *Canvas) Width() int {
	return len(c.color[0])
}

func (c *Canvas) Height() int {
	return len(c.color)
}
