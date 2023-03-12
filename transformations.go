package rt

import "fmt"

// from: where you want the eye to be in the scene
// to  : where you want to look
// up  : which direction is "up"
func ViewTransform(from Point, to Point, up Vector) Matrix {
	fmt.Printf("VT from:      %s\n", from)
	fmt.Printf("VT to:        %s\n", to)
	fmt.Printf("VT up:        %s\n", up)
	var (
		forward  = to.SubtractPoint(from).Normalize()
		normalUp = up.Normalize()
		left     = forward.Cross(normalUp)
		trueUp   = left.Cross(forward)
	)

	fmt.Printf("VT forward:   %s\n", forward)
	fmt.Printf("VT normalup:  %s\n", normalUp)
	fmt.Printf("VT left:      %s\n", left)
	fmt.Printf("VT trueup:    %s\n", trueUp)

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

	fmt.Printf("Orientation:\n%s", o)

	// translate to move the scene into place
	res := o.Multiply(Translation(-from.X(), -from.Y(), -from.Z()))
	return res
}
