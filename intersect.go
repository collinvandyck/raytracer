package rt

import (
	"fmt"
	"sort"
)

var noIntersection Intersection

func Hit(xs Intersections) (Intersection, bool) {
	for _, x := range xs {
		if x.Value() >= 0 {
			return x, true
		}
	}
	return noIntersection, false
}

type Intersections []Intersection

func NewIntersections(xs ...Intersection) Intersections {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].Value() < xs[j].Value()
	})
	return xs
}

func (i Intersections) Len() int {
	return len(i)
}

func (i Intersections) Get(idx int) Intersection {
	return i[idx]
}

func (i Intersections) Values() []Value {
	res := make([]Value, len(i))
	for x := range i {
		res[x] = i[x].Value()
	}
	return res
}

func (i Intersections) Equal(o Intersections) bool {
	if len(i) != len(o) {
		return false
	}
	for x := range i {
		if !i[x].Equal(o[x]) {
			return false
		}
	}
	return true
}

func (i Intersections) String() string {
	return fmt.Sprintf("%v", []Intersection(i))
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
	return i.shape.EqualShape(o.Shape())
}

func (i Intersection) String() string {
	return fmt.Sprintf("x(t:%s, s:%v)", formatFloat(i.value), i.shape)
}
