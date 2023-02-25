package rt

type Matrix [][]float

func (m Matrix) Get(row int, column int) float {
	return m[row][column]
}
