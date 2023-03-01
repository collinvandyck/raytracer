package rt

type Sphere struct {
}

// todo: do i need to force an allocation here?
func NewSphere() Sphere {
	return Sphere{}
}
