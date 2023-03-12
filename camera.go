package rt

type Camera struct {
	hsize int
	vsize int
	fov   Value
}

func NewCamera(hsize, vsize int, fov Value) *Camera {
	return &Camera{
		hsize: hsize,
		vsize: vsize,
		fov:   fov,
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
