package rt

type Intersection struct {
	ts    []value
	shape Shape
}

func NewIntersection(ts ...value) Intersection {
	return Intersection{
		ts: ts,
	}
}

func (i Intersection) Shape() Shape {
	return i.shape
}

func (i *Intersection) SetShape(shape Shape) {
	i.shape = shape
}

func (i Intersection) Get() []value {
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
	if i.shape == nil {
		return o.shape == nil
	}
	return i.shape.Equal(o.shape)
}
