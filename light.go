package rt

type PointLight struct {
	position  Point
	intensity Color
}

func NewPointLight(position Point, intensity Color) PointLight {
	return PointLight{
		position:  position,
		intensity: intensity,
	}
}

func (p PointLight) Position() Point {
	return p.position
}

func (p PointLight) Intensity() Color {
	return p.intensity
}
