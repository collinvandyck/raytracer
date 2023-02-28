package rt

import (
	"sync"
)

type matrixPool struct {
	pool2x2 *sync.Pool
	pool3x3 *sync.Pool
	pool4x4 *sync.Pool
}

func newMatrixPool() *matrixPool {
	return &matrixPool{
		pool2x2: &sync.Pool{
			New: func() any {
				return AllocateMatrix(2, 2)
			},
		},
		pool3x3: &sync.Pool{
			New: func() any {
				return AllocateMatrix(3, 3)
			},
		},
		pool4x4: &sync.Pool{
			New: func() any {
				return AllocateMatrix(4, 4)
			},
		},
	}
}

func (p *matrixPool) New(rows, cols int) *Matrix {
	switch {
	case rows == cols && rows == 2:
		return p.pool2x2.Get().(*Matrix)
	case rows == cols && rows == 3:
		return p.pool3x3.Get().(*Matrix)
	case rows == cols && rows == 4:
		return p.pool4x4.Get().(*Matrix)
	}
	return AllocateMatrix(rows, cols)
}

func (p *matrixPool) Put(m *Matrix) {
	rows, cols := m.Rows(), m.Cols()
	var pool *sync.Pool
	switch {
	case rows == cols && rows == 2:
		pool = p.pool2x2
	case rows == cols && rows == 3:
		pool = p.pool3x3
	case rows == cols && rows == 4:
		pool = p.pool4x4
	}
	if pool != nil {
		zero := []float{0, 0, 0, 0, 0, 0, 0, 0}
		for i := 0; i < len(m.vals); i++ {
			copy(m.vals[i], zero)
		}
		pool.Put(m)
	}
}
