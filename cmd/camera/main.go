package main

import (
	"fmt"
	"log"
	"rt"
	"rt/image"
	"time"
)

func main() {
	// the floor is a flattened sphere with a matte texture
	floor := rt.NewSphere()
	floor.SetTransform(rt.Scaling(10, 0.01, 10))
	floor.Material().SetColor(rt.NewColor(1, 0.9, 0.9))
	floor.Material().SetSpecular(0)
	floor.Material().SetColor(rt.NewColor(0, 0.7, 1))

	// the wall on the left has the same scale and color as the
	// floor but is also rotated and translated.
	leftWall := rt.NewSphere()
	leftWall.SetMaterial(floor.Material())
	leftWall.SetTransform(
		rt.Translation(0, 0, 5).Multiply(
			rt.RotationY(-rt.Pi / 4).Multiply(
				rt.RotationX(rt.Pi / 2).Multiply(
					rt.Scaling(10, 0.01, 10),
				),
			),
		),
	)

	// the right wall is the same but it's rotated opposite wrt y
	rightWall := rt.NewSphere()
	rightWall.SetMaterial(floor.Material())
	rightWall.SetTransform(
		rt.Translation(0, 0, 5).Multiply(
			rt.RotationY(rt.Pi / 4).Multiply(
				rt.RotationX(rt.Pi / 2).Multiply(
					rt.Scaling(10, 0.01, 10),
				),
			),
		),
	)

	// the large sphere in the middle is a unit sphere, translated
	// upward slightly and also colored green
	middle := rt.NewSphere()
	middle.SetTransform(rt.Translation(-0.5, 1, 0.5))
	middle.Material().SetColor(rt.NewColor(0.1, 1, 0.5))
	middle.Material().SetDiffuse(0.7)
	middle.Material().SetSpecular(0.9)

	// the smaller green sphere on the right is scaled in half
	right := rt.NewSphere()
	right.SetTransform(rt.Translation(1.5, 0.5, -0.5).Multiply(
		rt.Scaling(0.5, 0.5, 0.5),
	))
	right.Material().SetDiffuse(0.7)
	right.Material().SetSpecular(0.3)
	right.Material().SetColor(rt.NewColor(1, 0, 0))

	// the smallest sphere is scaled by a third, before being translated
	left := rt.NewSphere()
	left.SetTransform(rt.Translation(-1.5, 0.33, -0.75).Multiply(
		rt.Scaling(0.33, 0.33, 0.33),
	))
	left.Material().SetColor(rt.NewColor(1, 0.8, 0.1))
	left.Material().SetDiffuse(0.7)
	left.Material().SetSpecular(0.3)

	// create the world
	world := rt.NewWorld()

	// add the shapes
	world.AddShapes(floor, leftWall, rightWall, middle, right, left)

	// The light source is white shining from above and to the left
	world.SetLight(rt.NewPointLight(rt.NewPoint(-10, 10, -10), rt.NewColor(1, 1, 1)))

	// set up our camera
	cam := rt.NewCamera(1024, 768, rt.Pi/3)
	camFrom := rt.NewPoint(0, 1.5, -5)
	//camFrom = rt.NewPoint(10, 15.5, -25)
	camTo := rt.NewPoint(0, 1, 0)
	camUp := rt.NewVector(0, 1, 0)
	cam.SetTransform(rt.ViewTransform(camFrom, camTo, camUp))

	start := time.Now()
	canvas := cam.Render(world)
	dur := time.Since(start)
	fmt.Printf("Total time: %s\n", dur.Truncate(time.Millisecond))
	pixels := cam.HSize() * cam.VSize()
	durPixel := float64(dur) / float64(pixels)
	fmt.Printf("Per pixel : %s\n", time.Duration(durPixel))

	err := image.WritePNGTo(canvas, "camera.png")
	if err != nil {
		log.Fatal(err)
	}
}
