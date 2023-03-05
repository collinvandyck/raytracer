package rt

type Intersection struct {
	ts    []float
	shape Shape
}

func NewIntersection(ts ...float) Intersection {
	return Intersection{
		ts: ts,
	}
}

func (i *Intersection) SetShape(shape Shape) {
	i.shape = shape
}

func (i Intersection) Get() []float {
	return i.ts
}

func (i Intersection) Len() int {
	return len(i.ts)
}

func (i Intersection) Equal(o Intersection) bool {
	if i.Len() != o.Len() {
		return false
	}
	for idx := 0; idx < i.Len(); idx++ {
		if !floatsEqual(i.ts[idx], o.ts[idx]) {
			return false
		}
	}
	if i.sphok != o.sphok {
		return false
	}
	if i.sph != o.sph {
		return false
	}
	return true
}
