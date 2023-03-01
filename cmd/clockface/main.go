package main

import (
	"fmt"
	"log"
	"rt"
)

func main() {
	color := rt.NewColor(1, 1, 0)
	c := rt.NewCanvas(500, 500)
	c.SetPointSize(4)

	tCenter := rt.Translation(float64(c.Width())/2, float64(c.Height())/2, 0)

	writePoint := func(point rt.Point) {
		p1 := tCenter.MultiplyPoint(point)
		c.WritePoint(p1, color)
		fmt.Println("Writing point", p1)
	}

	origin := rt.NewPoint(0, 0, 0)
	writePoint(origin)

	// start at 12:00
	start := rt.NewPoint(0, 180, 0)
	writePoint(start)

	err := rt.WritePPMTo(c, "clockface.ppm")
	if err != nil {
		log.Fatal(err)
	}
}
