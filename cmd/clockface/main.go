package main

import (
	"log"
	"rt"
)

func main() {
	color := rt.NewColor(1, 1, 0)
	c := rt.NewCanvas(500, 500)
	c.SetPointSize(4)

	var (
		cx = float64(c.Width()) / 2
		cy = float64(c.Height()) / 2
		tc = rt.Translation(cx, cy, 0) // shift (0,0) to center of canvas
		ts = rt.Scaling(1, -1, 1)      // flip the y coord
		t1 = tc.Multiply(ts)           // flip the y coord, and then shift to center
	)

	// transforms and then writes the point to the canvas
	write := func(point rt.Point) {
		p1 := t1.MultiplyPoint(point)
		c.WritePoint(p1, color)
	}

	origin := rt.NewPoint(0, 0, 0)
	write(origin)

	// current at 12:00 (noon)
	current := rt.NewPoint(0, 180, 0)
	write(current)

	// rotate by an hour (1/3 * pi/2) for the remaining hours
	rotate := rt.RotationZ(rt.Pi / -6)
	for i := 1; i < 12; i++ {
		current = rotate.MultiplyPoint(current)
		write(current)
	}

	err := rt.WritePPMTo(c, "clockface.ppm")
	if err != nil {
		log.Fatal(err)
	}
}
