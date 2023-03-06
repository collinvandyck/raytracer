package main

import (
	"log"
	"rt"
)

func main() {

	var (
		canvasPixels = 500                                      // the dimensions of the canvas
		wallSize     = rt.Value(7)                              // how big the wall is
		pixelSize    = wallSize / rt.Value(canvasPixels)        // pixel size in world coordinates
		half         = wallSize / 2                             // half the wall size
		color        = rt.NewColor(1, 0, 0)                     // color of the sphere
		canvas       = rt.NewCanvas(canvasPixels, canvasPixels) // size of the canvas
		rayOrigin    = rt.NewPoint(0, 0, -5)                    // ray origin
		sphere       = rt.NewSphere()                           // unit sphere
	)

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

	err := rt.WritePPMTo(canvas, "sphere.ppm")
	if err != nil {
		log.Fatal(err)
	}
}
