package rt

import "fmt"

type PointLight struct {
	position  Point
	intensity Color
}

func NewPointLight(position Point, intensity Color) *PointLight {
	return &PointLight{
		position:  position,
		intensity: intensity,
	}
}

func (p *PointLight) Position() Point {
	return p.position
}

func (p *PointLight) Intensity() Color {
	return p.intensity
}

func (p *PointLight) Equal(o *PointLight) bool {
	if p == nil {
		return o == nil
	}
	if !p.Position().Equal(o.Position()) {
		return false
	}
	if !p.Intensity().Equal(o.Intensity()) {
		return false
	}
	return true
}

func (p *PointLight) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PointLight(pos=%s int=%s)", p.Position(), p.Intensity())
}
