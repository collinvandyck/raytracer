package main

import (
	"fmt"
	"os"
	"rt"
)

type projectile struct {
	point rt.Point
	veloc rt.Vector
}

type env struct {
	grav rt.Vector
	wind rt.Vector
}

func tick(env env, p projectile) projectile {
	newPoint := p.point.AddVector(p.veloc)
	newVeloc := p.veloc.AddVector(env.grav.AddVector(env.wind))
	return projectile{
		point: newPoint,
		veloc: newVeloc,
	}
}

func main() {
	proj := projectile{
		point: rt.NewPoint(0, 1, 0),
		veloc: rt.NewVector(1, 1.8, 0).Normalize().MultiplyBy(11.25),
	}
	env := env{
		grav: rt.NewVector(0, -0.1, 0),
		wind: rt.NewVector(-0.01, 0, 0),
	}
	cv := rt.NewCanvas(900, 550)
	report := func() {
		pjp := proj.point
		plp := rt.NewPoint(pjp.X(), 550-pjp.Y(), 0)
		color := rt.NewColor(1, 0, 0)
		cv.WritePixel(int(plp.X()), int(plp.Y()), color)
	}
	for {
		report()
		proj = tick(env, proj)
		if proj.point.Y() <= 0 {
			break
		}
	}
	f, err := os.Create("cannon.ppm")
	check(err)
	err = rt.WritePPM(cv, f)
	check(err)
	err = f.Close()
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
