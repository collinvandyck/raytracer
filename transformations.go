package rt

// from: where you want the eye to be in the scene
// to  : where you want to look
// up  : which direction is "up"
func ViewTransform(from Point, to Point, up Vector) Matrix {
	return MatrixIdentity4x4
}
