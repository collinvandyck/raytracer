package main

import (
	"log"
	"rt"
)

func main() {
	c := rt.NewCanvas(500, 500)
	err := rt.WritePPMTo(c, "clockface.ppm")
	if err != nil {
		log.Fatal(err)
	}
}
