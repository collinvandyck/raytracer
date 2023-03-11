package rt

import "math"

type Material struct {
	color     Color
	ambient   Value // typically between 0 and 1
	diffuse   Value // typically between 0 and 1
	specular  Value // typically between 0 and 1
	shininess Value
}

func Lighting(m Material, light *PointLight, position Point, eye Vector, normal Vector) Color {
	// combine the material color with the light's color/intensity
	effectiveColor := m.Color().Multiply(light.Intensity())

	// find the direction to the light source.
	lightv := light.Position().SubtractPoint(position).Normalize()

	// compute the ambient contribution
	ambient := effectiveColor.MultiplyBy(m.Ambient())

	// lightDotNormal is the cosine of the angle between the light vector
	// and the normal vector. A negative value means that the light is on
	// the other side of the surface.
	lightDotNormal := lightv.Dot(normal)

	diffuse := black
	specular := black

	if lightDotNormal >= 0 {

		// compute the diffuse contribution
		diffuse = effectiveColor.MultiplyBy(m.Diffuse()).MultiplyBy(lightDotNormal)

		// reflectDotEye is the cosine of the angle between the reflection vector
		// and the eye vector. A negative value means that the light reflects
		// away from the eye
		reflectv := lightv.Negate().Reflect(normal)
		reflectDotEye := reflectv.Dot(eye)

		if reflectDotEye > 0 {
			// compute the specular contribution
			factor := math.Pow(reflectDotEye, m.Shininess())
			specular = light.Intensity().MultiplyBy(m.Specular()).MultiplyBy(factor)
		}
	}
	return ambient.Add(diffuse).Add(specular)
}

func NewBlankMaterial() Material {
	return Material{}
}

func NewMaterial(color Color, ambient, diffuse, specular, shininess Value) Material {
	return Material{
		color:     color,
		ambient:   ambient,
		diffuse:   diffuse,
		specular:  specular,
		shininess: shininess,
	}
}

func DefaultMaterial() Material {
	color := NewColor(1, 1, 1)
	ambient := 0.1
	diffuse := 0.9
	specular := 0.9
	shininess := 200.0
	return NewMaterial(color, ambient, diffuse, specular, shininess)
}

func (m Material) Color() Color {
	return m.color
}

func (m *Material) SetColor(c Color) {
	m.color = c
}

func (m Material) Ambient() Value {
	return m.ambient
}

func (m *Material) SetAmbient(v Value) {
	m.ambient = v
}

func (m Material) Diffuse() Value {
	return m.diffuse
}

func (m *Material) SetDiffuse(v Value) {
	m.diffuse = v
}

func (m Material) Specular() Value {
	return m.specular
}

func (m *Material) SetSpecular(v Value) {
	m.specular = v
}

func (m Material) Shininess() Value {
	return m.shininess
}

func (m *Material) SetShininess(v Value) {
	m.shininess = v
}

func (m Material) Equal(o Material) bool {
	if !m.Color().Equal(o.Color()) {
		return false
	}
	if !floatsEqual(m.Ambient(), o.Ambient()) {
		return false
	}
	if !floatsEqual(m.Diffuse(), o.Diffuse()) {
		return false
	}
	if !floatsEqual(m.Specular(), o.Specular()) {
		return false
	}
	if !floatsEqual(m.Shininess(), o.Shininess()) {
		return false
	}
	return true
}
