package rt

type Material struct {
	color     Color
	ambient   Value // typically between 0 and 1
	diffuse   Value // typically between 0 and 1
	specular  Value // typically between 0 and 1
	shininess Value
}

func Lighting(m Material, light PointLight, position Point, eye Vector, normal Vector) Color {
	return NewColor(1, 1, 1)
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

func (m Material) GetColor() Color {
	return m.color
}

func (m *Material) SetColor(c Color) {
	m.color = c
}

func (m Material) GetAmbient() Value {
	return m.ambient
}

func (m *Material) SetAmbient(v Value) {
	m.ambient = v
}

func (m Material) GetDiffuse() Value {
	return m.diffuse
}

func (m *Material) SetDiffuse(v Value) {
	m.diffuse = v
}

func (m Material) GetSpecular() Value {
	return m.specular
}

func (m *Material) SetSpecular(v Value) {
	m.specular = v
}

func (m Material) GetShininess() Value {
	return m.shininess
}

func (m *Material) SetShininess(v Value) {
	m.shininess = v
}

func (m Material) Equal(o Material) bool {
	if !m.GetColor().Equal(o.GetColor()) {
		return false
	}
	if !floatsEqual(m.GetAmbient(), o.GetAmbient()) {
		return false
	}
	if !floatsEqual(m.GetDiffuse(), o.GetDiffuse()) {
		return false
	}
	if !floatsEqual(m.GetSpecular(), o.GetSpecular()) {
		return false
	}
	if !floatsEqual(m.GetShininess(), o.GetShininess()) {
		return false
	}
	return true
}
