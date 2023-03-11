package rt

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

func NewWorld() *World {
	return &World{}
}

func NewDefaultWorld() *World {
	l1 := NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1))
	s1 := NewSphere()
	m1 := NewBlankMaterial()
	m1.SetColor(NewColor(0.8, 1.0, 0.6))
	m1.SetDiffuse(0.7)
	m1.SetSpecular(0.2)
	s1.SetMaterial(m1)
	s2 := NewSphere()
	s2.SetTransform(Scaling(0.5, 0.5, 0.5))
	return &World{
		shapes: []Shape{s1, s2},
		light:  l1,
	}
}

type World struct {
	shapes []Shape
	light  *PointLight
}

func (w *World) Shapes() []Shape {
	return w.shapes
}

func (w *World) Light() *PointLight {
	return w.light
}
