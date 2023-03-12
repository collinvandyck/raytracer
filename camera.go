package rt

import (
	"math"
)

type Camera struct {
	hsize      int
	vsize      int
	fov        Value
	transform  Matrix
	inverse    Matrix
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

func (c *Camera) Render(world *World) *Canvas {
	canvas := NewCanvas(c.hsize, c.vsize)
	for y := 0; y < c.vsize-1; y++ {
		for x := 0; x < c.hsize-1; x++ {
			ray := c.RayForPixel(x, y)
			color := world.ColorAt(ray)
			canvas.WritePixel(x, y, color)
		}
	}
	return canvas
}

// Compute the ray from the camera to the (x,y) pixel on the canvas
func (c *Camera) RayForPixel(px, py int) Ray {
	var (
		pxVal   = Value(px)
		pyVal   = Value(py)
		xOffset = (pxVal + 0.5) * c.pixelSize // x offset from edge of canvas to pixel center
		yOffset = (pyVal + 0.5) * c.pixelSize // y offset from edge of canvas to pixel center
		worldX  = c.halfWidth - xOffset       // cam looks toward -z so +x is to the left
		worldY  = c.halfHeight - yOffset      // y still operates normally
		canvasZ = Value(-1)                   // canvas is positioned at z-1
		inverse = c.InverseTransform()        // we use this to move the camera
	)

	// tranform the canvas point
	pixel := inverse.MultiplyPoint(NewPoint(worldX, worldY, canvasZ))

	// transform the camera origin
	origin := inverse.MultiplyPoint(Origin)

	direction := pixel.SubtractPoint(origin).Normalize()
	ray := NewRay(origin, direction)
	return ray
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

func (c *Camera) SetTransform(m Matrix) {
	c.transform = m
	c.inverse = emptyMatrix
}

func (c *Camera) InverseTransform() Matrix {
	if c.inverse.Empty() {
		c.inverse = c.Transform().Inverse()
	}
	return c.inverse
}

func (c *Camera) PixelSize() Value {
	return c.pixelSize
}
