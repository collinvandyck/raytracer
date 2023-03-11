package rt

import (
	"fmt"
	"math"
	"sort"
)

type Computations struct {
	shape   Shape
	point   Point
	eyev    Vector
	normalv Vector
	value   Value
}

func PrepareComputations(i Intersection, ray Ray) *Computations {
	res := &Computations{}

	// copy the intersection's properties, for convenience
	res.value = i.Value()
	res.shape = i.Shape()

	// precompute useful values
	res.point = ray.Position(res.value)
	res.eyev = ray.Direction().Negate()
	res.normalv = res.shape.NormalAt(res.point)

	return res
}

func (c *Computations) Value() Value {
	return c.value
}

func (c *Computations) Shape() Shape {
	return c.shape
}

func (c *Computations) Point() Point {
	return c.point
}

func (c *Computations) Eye() Vector {
	return c.eyev
}

func (c *Computations) Normal() Vector {
	return c.normalv
}

func IntersectWorld(world *World, ray Ray) Intersections {
	res := NewIntersections()
	shapes := world.Shapes()
	for i := 0; i < len(shapes); i++ {
		shape := shapes[i]
		xs := shape.Intersect(ray)
		res.AddAll(xs)
	}
	res.Sort()
	return res
}

func IntersectSphere(sphere *Sphere, ray Ray) Intersections {
	// transform the ray into object coordinates
	ray = ray.Transform(sphere.GetInverseTransform())

	// the vector from the sphere's center to the ray origin
	sphereToRay := ray.Origin().SubtractPoint(sphere.Point())

	a := ray.Direction().Dot(ray.Direction())
	b := 2 * ray.Direction().Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		// there was no intersection
		return nil
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	i1 := NewIntersection(t1, sphere)
	i2 := NewIntersection(t2, sphere)
	return NewIntersections(i1, i2)
}

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
	res := Intersections(xs)
	res.Sort()
	return res
}

func (xs Intersections) Sort() {
	if len(xs) == 0 {
		return
	}
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].Value() < xs[j].Value()
	})
}

func (i *Intersections) Add(x Intersection) {
	*i = append(*i, x)
}

func (i *Intersections) AddAll(xs Intersections) {
	*i = append(*i, xs...)
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
