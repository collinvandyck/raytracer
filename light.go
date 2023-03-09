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

func (p PointLight) GetPosition() Point {
	return p.position
}

func (p PointLight) GetIntensity() Color {
	return p.intensity
}
