package rt

type Matrix [][]float

func NewMatrix(rows, cols int) Matrix {
	m := make(Matrix, rows)
	for i := range m {
		m[i] = make([]float, cols)
	}
	return m
}

func NewMatrixFromTable(table string) Matrix {
	return Matrix{}
}

func (m Matrix) Get(row int, column int) float {
	return m[row][column]
}
