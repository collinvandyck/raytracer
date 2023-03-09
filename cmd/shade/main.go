package main

import (
	"fmt"
	"log"
	"rt"
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
	err := rt.WritePPMTo(canvas, "shade.ppm")
	if err != nil {
		log.Fatal(err)
	}
}

func render(canvasPixels int) *rt.Canvas {
	var (
		wallSize   = rt.Value(7)                              // how big the wall is
		pixelSize  = wallSize / rt.Value(canvasPixels)        // pixel size in world coordinates
		half       = wallSize / 2                             // half the wall size
		canvas     = rt.NewCanvas(canvasPixels, canvasPixels) // size of the canvas
		rayOrigin  = rt.NewPoint(0, 0, -5)                    // ray origin
		sphere     = rt.NewSphere()                           // unit sphere
		material   = rt.DefaultMaterial()                     // the material of the sphere
		lightPos   = rt.NewPoint(-10, 10, -10)                // light position above and to the left of the eys
		lightColor = rt.NewColor(1, 1, 1)                     // the light will be a white light
		light      = rt.NewPointLight(lightPos, lightColor)   // the singular light source
	)

	material.SetColor(rt.NewColor(1, 0.2, 1))
	sphere.SetMaterial(material)

	for y := 0; y < canvasPixels; y++ {
		// compute worldY (top = +half, bottom = -half)
		worldY := half - (float64(y) * pixelSize)

		for x := 0; x < canvasPixels; x++ {
			worldX := half - (float64(x) * pixelSize)
			point := rt.NewPoint(worldX, worldY, 10) // the wall lives at z=10
			rayDirection := point.SubtractPoint(rayOrigin).Normalize()

			ray := rt.NewRay(rayOrigin, rayDirection)
			// TODO: we are already normalizing our vector, so we don't need to do this, possibly
			ray.NormalizeDirection()

			xs := rt.IntersectSphere(sphere, ray)
			hit, ok := rt.Hit(xs)
			if !ok {
				// we missed the sphere. no need to continue
				continue
			}

			// here we have a hit.
			hitPoint := ray.Position(hit.Value())
			hitNormal := hit.Shape().NormalAt(hitPoint)
			hitEye := ray.Direction().Negate()

			// calculate the color.
			hitColor := rt.Lighting(hit.Shape().Material(), light, hitPoint, hitEye, hitNormal)
			canvas.WritePixel(x, y, hitColor)

		}
	}
	return canvas
}
