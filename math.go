package rt

import (
	"math"
	"strconv"
)

const (
	EPSILON = 0.00001
	Pi      = math.Pi
	Sqrt2   = math.Sqrt2
)

func floatsEqual(v1, v2 Value) bool {
	return math.Abs(v1-v2) < EPSILON
}

func formatFloat(val Value) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

/*
func radians(degrees float) float {
	return (degrees / 180) * math.Pi
}
*/
