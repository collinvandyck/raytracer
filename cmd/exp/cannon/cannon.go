package main

import (
	"fmt"
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
	p := projectile{
		point: rt.NewPoint(0, 1, 0),
		veloc: rt.NewVector(1, 8, 0).Normalize(),
	}
	e := env{
		grav: rt.NewVector(0, -0.1, 0),
		wind: rt.NewVector(-0.01, 0, 0),
	}
	report := func() {
		fmt.Println(p.point.Y())
	}
	for {
		report()
		p = tick(e, p)
		if p.point.Y() <= 0 {
			break
		}
	}
}
