package rt

import "math"

const (
	EPSILON = 0.00001
	Pi      = math.Pi
)

func floatsEqual(v1, v2 float) bool {
	return math.Abs(v1-v2) < EPSILON
}

func radians(degrees float) float {
	return (degrees / 180) * math.Pi
}
