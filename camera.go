package rt

import (
	"math"
)

type Camera struct {
	hsize      int
	vsize      int
	fov        Value
	transform  Matrix
	pixelSize  Value
	halfWidth  Value
	halfHeight Value
}

func NewCamera(hsize, vsize int, fov Value) *Camera {
	var (
		halfView   = math.Tan(fov / 2)
		aspect     = Value(hsize) / Value(vsize)
		halfWidth  Value
		halfHeight Value
	)
	if aspect >= 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfHeight = halfView
		halfWidth = halfView * aspect
	}
	pixelSize := (halfWidth * 2) / Value(hsize)
	return &Camera{
		hsize:      hsize,
		vsize:      vsize,
		fov:        fov,
		transform:  MatrixIdentity4x4,
		halfWidth:  halfWidth,
		halfHeight: halfHeight,
		pixelSize:  pixelSize,
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
