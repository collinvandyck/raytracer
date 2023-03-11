package rt

func NewWorld() *World {
	return &World{}
}

type World struct {
	shapes []Shape
	lights []PointLight
}

func (w *World) Shapes() []Shape {
	return w.shapes
}

func (w *World) Lights() []PointLight {
	return w.lights
}
