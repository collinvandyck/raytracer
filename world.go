package rt

func ColorAt(world *World, ray Ray) Color {
	xs := IntersectWorld(world, ray)
	hit, ok := Hit(xs)
	if !ok {
		return black
	}
	comps := PrepareComputations(hit, ray)
	color := ShadeHit(world, comps)
	return color
}

func ShadeHit(world *World, computations *Computations) Color {
	var (
		material = computations.Shape().Material()
		light    = world.Light()
		position = computations.Point()
		eye      = computations.Eye()
		normal   = computations.Normal()
	)
	color := Lighting(material, light, position, eye, normal)
	return color
}

func NewWorld() *World {
	return &World{}
}

func NewDefaultWorld() *World {
	l1 := NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1))
	s1 := NewSphere()
	m1 := DefaultMaterial()
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

func (w *World) ColorAt(ray Ray) Color {
	return ColorAt(w, ray)
}

func (w *World) Shapes() []Shape {
	return w.shapes
}

func (w *World) Shape(i int) Shape {
	return w.shapes[i]
}

func (w *World) Light() *PointLight {
	return w.light
}

func (w *World) SetLight(pl *PointLight) {
	w.light = pl
}
