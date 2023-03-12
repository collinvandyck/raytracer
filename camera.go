package rt

type Camera struct {
	hsize     int
	vsize     int
	fov       Value
	transform Matrix
	pixelSize Value
}

func NewCamera(hsize, vsize int, fov Value) *Camera {
	return &Camera{
		hsize:     hsize,
		vsize:     vsize,
		fov:       fov,
		transform: MatrixIdentity4x4,
	}
}

func (c *Camera) HSize() int {
	return c.hsize
}

func (c *Camera) VSize() int {
	return c.vsize
}

func (c *Camera) FOV() Value {
	return c.fov
}

func (c *Camera) Transform() Matrix {
	return c.transform
}

func (c *Camera) PixelSize() Value {
	return c.pixelSize
}
