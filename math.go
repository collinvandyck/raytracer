package rt

import "math"

const EPSILON = 0.00001

func floatsEqual(v1, v2 float) bool {
	return math.Abs(v1-v2) < EPSILON
}
