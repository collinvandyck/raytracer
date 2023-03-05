package rt

type Intersections []Intersection

func NewIntersections(intersections ...Intersection) Intersections {
	return intersections
}

type Intersection struct {
	value Value
	shape Shape
}

func NewIntersection(t Value, s Shape) Intersection {
	return Intersection{
		shape: s,
		value: t,
	}
}

func (i Intersection) Value() Value {
	return i.value
}

func (i Intersection) Shape() Shape {
	return i.shape
}

func (i Intersection) Equal(o Intersection) bool {
	if !floatsEqual(i.value, o.value) {
		return false
	}
	if i.shape == nil {
		return o.shape == nil
	}
	return i.shape == o.shape
}
