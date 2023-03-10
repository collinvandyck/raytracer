package image

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"io"
	"os"
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

func WritePNGTo(c *rt.Canvas, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return WritePNG(c, f)
}

func WriteGIFTo(delay int, cvs []*rt.Canvas, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return WriteGIF(delay, cvs, f)
}

func WritePNG(c *rt.Canvas, w io.Writer) error {
	rgb := CanvasToRGBA(c)
	return png.Encode(w, rgb)
}

func WriteGIF(delay int, cvs []*rt.Canvas, w io.Writer) error {
	res := CanvasesToGIF(delay, cvs)
	return gif.EncodeAll(w, &res)
}

func CanvasesToGIF(delay int, cvs []*rt.Canvas) gif.GIF {
	palette := make(color.Palette, 256)
	for i := 0; i < 256; i++ {
		c := color.RGBA{uint8(i), uint8(i), uint8(i), 255}
		palette[i] = c
	}
	palette = palette[:255]
	pals := []*image.Paletted{}
	var dels []int
	for _, cv := range cvs {
		fmt.Println("Creating RGBA")
		rgb := CanvasToRGBA(cv)
		fmt.Println("Converting to palleted")
		pal := convertRGBAtoPaletted(rgb, palette)
		pals = append(pals, pal)
		dels = append(dels, delay)
	}
	fmt.Println("Returning GIF")
	return gif.GIF{
		Image: pals,
		Delay: dels,
	}
}

func CanvasToRGBA(c *rt.Canvas) *image.RGBA {
	rect := image.Rect(0, 0, c.Width(), c.Height())
	res := image.NewRGBA(rect)
	for x := 0; x < c.Width(); x++ {
		for y := 0; y < c.Height(); y++ {
			cc := c.PixelAt(x, y)
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

// Helper function to convert an RGBA image to a paletted image
func convertRGBAtoPaletted(img *image.RGBA, p color.Palette) *image.Paletted {
	bounds := img.Bounds()
	paletted := image.NewPaletted(bounds, p)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			paletted.Set(x, y, img.At(x, y))
		}
	}
	return paletted
}
