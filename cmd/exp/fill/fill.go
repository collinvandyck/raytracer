package main

import (
	"fmt"
	"os"
	"rt"
)

func main() {
	cv := rt.NewCanvas(900, 550)
	cv.Fill(rt.NewColor(0, 0, 1))
	f, err := os.Create("fill.ppm")
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
