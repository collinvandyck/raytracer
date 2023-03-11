package rt

func NormalAtSphere(s1 *Sphere, p1 Point) Vector {
	mt := s1.GetTransposedInverseTransform()         // transposed inverse transform
	op := s1.GetInverseTransform().MultiplyPoint(p1) // object space
	on := op.SubtractPoint(Origin)                   // object normal
	wn := mt.MultiplyVector(on)                      // world normal
	wn.SetW(0)                                       // correct transposition side effect
	nv := wn.Normalize()                             // normalize
	return nv
}

type Sphere struct {
	transform         Matrix
	inverse           Matrix
	transposedInverse Matrix
	material          Material
}

func NewSphere() *Sphere {
	return &Sphere{
		material: DefaultMaterial(),
	}
}

func (s *Sphere) Intersect(r Ray) Intersections {
	return IntersectSphere(s, r)
}

func (s *Sphere) NormalAt(point Point) Vector {
	return NormalAtSphere(s, point)
}

func (s *Sphere) GetTransposedInverseTransform() Matrix {
	if s.transposedInverse.Empty() {
		m := s.GetInverseTransform()
		s.transposedInverse = m.Transpose()
	}
	return s.transposedInverse
}

func (s *Sphere) GetInverseTransform() Matrix {
	if s.inverse.Empty() {
		m := s.GetTransform()
		s.inverse = m.Inverse()
	}
	return s.inverse
}

func (s *Sphere) GetTransform() Matrix {
	if s.transform.Empty() {
		return MatrixIdentity4x4
	}
	return s.transform
}

func (s *Sphere) SetTransform(m Matrix) {
	s.transform = m
	s.inverse = emptyMatrix
	s.transposedInverse = emptyMatrix
}

func (s *Sphere) EqualShape(o Shape) bool {
	os, ok := o.(*Sphere)
	if !ok {
		return false
	}
	return s.Equal(os)
}

func (s *Sphere) Equal(o *Sphere) bool {
	return true
}

func (s *Sphere) Point() Point {
	return NewPoint(0, 0, 0)
}

func (s *Sphere) Material() Material {
	return s.material
}

func (s *Sphere) SetMaterial(m Material) {
	s.material = m
}

func (s *Sphere) String() string {
	return "Sphere()"
}
