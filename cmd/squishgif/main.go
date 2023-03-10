package main

import (
	"log"
	"rt"
	"rt/image"
)

func main() {
	const canvasPixels = 1024
	const scale = 0.1
	var cvs []*rt.Canvas
	for s := 1.0; s >= 0; s -= scale {
		canvas := render(canvasPixels, s)
		cvs = append(cvs, canvas)
	}
	err := image.WriteGIFTo(10, cvs, "squish.gif")
	if err != nil {
		log.Fatal(err)
	}
}

type Canvas *rt.Canvas

func render(canvasPixels int, scale rt.Value) Canvas {
	var (
		worldWallSize = rt.Value(7)                              // how big the wall is
		pixelSize     = worldWallSize / rt.Value(canvasPixels)   // pixel size in world coordinates
		canvas        = rt.NewCanvas(canvasPixels, canvasPixels) // size of the canvas
		rayOrigin     = rt.NewPoint(0, 0, -5)                    // ray origin
		sphere        = rt.NewSphere()                           // unit sphere
		material      = rt.DefaultMaterial()                     // the material of the sphere
		lightPos      = rt.NewPoint(-10, 10, -10)                // light position above and to the left of the eye
		lightColor    = rt.NewColor(1, 1, 1)                     // the light will be a white light
		light         = rt.NewPointLight(lightPos, lightColor)   // the singular light source
	)

	material.SetColor(rt.NewColor(1, 0.2, 1))
	sphere.SetMaterial(material)
	sphere.SetTransform(rt.Scaling(scale, scale, scale))

	for y := 0; y < canvasPixels; y++ {
		// compute worldY (top = +half, bottom = -half)
		worldY := (worldWallSize / 2) - (float64(y) * pixelSize)

		for x := 0; x < canvasPixels; x++ {
			// compute worldX (left = -half, right = +half)
			worldX := (-worldWallSize / 2) + (float64(x) * pixelSize)

			point := rt.NewPoint(worldX, worldY, 10) // the wall lives at z=10
			rayDirection := point.SubtractPoint(rayOrigin).Normalize()

			ray := rt.NewRay(rayOrigin, rayDirection)

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
