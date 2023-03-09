package main

import (
	"fmt"
	"log"
	"rt"
	"rt/image"
	"time"
)

func main() {
	const canvasPixels = 1024
	start := time.Now()
	canvas := render(canvasPixels)
	dur := time.Since(start)
	durPixel := dur / time.Duration(canvasPixels*canvasPixels)
	fmt.Printf("Total time: %s\n", dur.Truncate(time.Millisecond))
	fmt.Printf("Per pixel : %s\n", durPixel)
	err := image.WritePNGTo(canvas, "sphere.png")
	if err != nil {
		log.Fatal(err)
	}
}

func render(canvasPixels int) *rt.Canvas {
	var (
		wallSize  = rt.Value(7)                              // how big the wall is
		pixelSize = wallSize / rt.Value(canvasPixels)        // pixel size in world coordinates
		half      = wallSize / 2                             // half the wall size
		color     = rt.NewColor(1, 0, 0)                     // color of the sphere
		canvas    = rt.NewCanvas(canvasPixels, canvasPixels) // size of the canvas
		rayOrigin = rt.NewPoint(0, 0, -5)                    // ray origin
		sphere    = rt.NewSphere()                           // unit sphere
	)

	//sphere.SetTransform(rt.Shearing(1, 0, 0, 0, 0, 0).Multiply(rt.Scaling(0.5, 1, 1)))

	for y := 0; y < canvasPixels; y++ {
		// compute worldY (top = +half, bottom = -half)
		worldY := half - (float64(y) * pixelSize)

		for x := 0; x < canvasPixels; x++ {
			worldX := half - (float64(x) * pixelSize)
			point := rt.NewPoint(worldX, worldY, 10) // the wall lives at z=10
			vector := point.SubtractPoint(rayOrigin).Normalize()
			ray := rt.NewRay(rayOrigin, vector)
			xs := rt.IntersectSphere(sphere, ray)
			_, hit := rt.Hit(xs)
			if hit {
				canvas.WritePixel(x, y, color)
			}

		}
	}
	return canvas
}
