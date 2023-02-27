package rt

import (
	"sync"
)

type matrixPool struct {
	pool2x2 sync.Pool
	pool3x3 sync.Pool
	pool4x4 sync.Pool
}

func newMatrixPool() *matrixPool {
	return &matrixPool{
		pool2x2: sync.Pool{
			New: func() any {
				return AllocateMatrix(2, 2)
			},
		},
		pool3x3: sync.Pool{
			New: func() any {
				return AllocateMatrix(3, 3)
			},
		},
		pool4x4: sync.Pool{
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
	default:
		return AllocateMatrix(rows, cols)
	}
}

func (p *matrixPool) Return(m *Matrix) {
	rows, cols := m.Rows(), m.Cols()
	switch {
	case rows == cols && rows == 2:
		m.reset()
		p.pool2x2.Put(m)
	case rows == cols && rows == 3:
		m.reset()
		p.pool3x3.Put(m)
	case rows == cols && rows == 4:
		m.reset()
		p.pool4x4.Put(m)
	}
}
