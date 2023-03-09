package rt

type Material struct {
	color     Color
	ambient   Value // typically between 0 and 1
	diffuse   Value // typically between 0 and 1
	specular  Value // typically between 0 and 1
	shininess Value
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

func (m Material) GetAmbient() Value {
	return m.ambient
}

func (m Material) GetDiffuse() Value {
	return m.diffuse
}

func (m Material) GetSpecular() Value {
	return m.specular
}

func (m Material) GetShininess() Value {
	return m.shininess
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
