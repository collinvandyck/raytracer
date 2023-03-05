package rt

type Shape interface {
	Equal(o Shape) bool
}

type Sphere struct {
}

// todo: do i need to force an allocation here?
func NewSphere() Sphere {
	return Sphere{}
}

func (s Sphere) Equal(o Shape) bool {
	if _, ok := o.(Sphere); !ok {
		return false
	}
	return true
}

func (s Sphere) Point() Point {
	return NewPoint(0, 0, 0)
}

func (s Sphere) Intersection(ts ...float) Intersection {
	res := NewIntersection(ts...)
	res.SetShape(s)
	return res
}

func (s Sphere) String() string {
	return "Sphere()"
}
