package image

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"rt"
)

func scaleRGB(val rt.Value) uint8 {
	const max = 255
	val *= max
	if val > max {
		val = max
	}
	return uint8(val)
}

func WritePNG(c rt.Canvas, w io.Writer) error {
	rgb := CanvasToRGBA(c)
	return png.Encode(w, rgb)
}

func CanvasToRGBA(c rt.Canvas) *image.RGBA {
	rect := image.Rect(0, 0, c.Width(), c.Height())
	res := image.NewRGBA(rect)
	for x := 0; x < c.Width(); x++ {
		for y := 0; y < c.Height(); y++ {
			cc := c.PixelAt(x, y)
			cc.Red()
			color := color.RGBA{
				R: scaleRGB(cc.Red()),
				G: scaleRGB(cc.Green()),
				B: scaleRGB(cc.Blue()),
				A: 255,
			}
			res.Set(x, y, color)
		}
	}
	return res
}
