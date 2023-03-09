package rt

type Shape interface {
	EqualShape(o Shape) bool
	NormalAt(p Point) Vector
	Material() Material
}
