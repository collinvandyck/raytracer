package rt

import "fmt"

// from: where you want the eye to be in the scene
// to  : where you want to look
// up  : which direction is "up"
func ViewTransform(from Point, to Point, up Vector) Matrix {
	var (
		forward  = to.SubtractPoint(from)
		normalUp = up.Normalize()
		left     = forward.Cross(normalUp)
		trueUp   = left.Cross(forward)
	)

	// orientation
	o := NewMatrix(4, 4)
	o.Set(0, 0, left.X())
	o.Set(0, 1, left.Y())
	o.Set(0, 2, left.Z())
	o.Set(1, 0, trueUp.X())
	o.Set(1, 1, trueUp.Y())
	o.Set(1, 2, trueUp.Z())
	o.Set(2, 0, -forward.X())
	o.Set(2, 1, -forward.Y())
	o.Set(2, 2, -forward.Z())
	o.Set(3, 3, 1)

	fmt.Printf("orientation:\n%s", o)

	// translate to move the scene into place
	res := o.Multiply(Translation(-from.X(), -from.Y(), -from.Z()))
	return res
}
