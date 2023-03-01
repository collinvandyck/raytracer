package rt

type Sphere struct {
}

// todo: do i need to force an allocation here?
func NewSphere() Sphere {
	return Sphere{}
}

func (s Sphere) Equal(o Sphere) bool {
	return true
}

func (s Sphere) Point() Point {
	return NewPoint(0, 0, 0)
}

func (s Sphere) Intersection(ts ...float) Intersection {
	res := NewIntersection(ts...)
	res.SetSphere(s)
	return res
}

func (s Sphere) String() string {
	return "Sphere()"
}
