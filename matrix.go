package rt

import (
	"bufio"
	"strconv"
	"strings"
)

type Matrix [][]float

func NewMatrix(rows, cols int) Matrix {
	m := make(Matrix, rows)
	for i := range m {
		m[i] = make([]float, cols)
	}
	return m
}

func NewMatrixFromTable(table string) (res Matrix) {
	table = strings.TrimSpace(table)
	s := bufio.NewScanner(strings.NewReader(table))
	rows := make([][]float, 0)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if !strings.HasPrefix(line, "|") {
			continue
		}
		parts := strings.Split(line, "|")
		row := make([]float, 0, 4)
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if len(part) == 0 {
				continue
			}
			num, err := strconv.ParseFloat(part, 64)
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}
		if len(rows) > 0 {
			if len(row) != len(rows[0]) {
				panic("inconsistent rows")
			}
		}
		rows = append(rows, row)
	}
	res = Matrix(rows)
	return
}

func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Cols() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

func (m Matrix) Get(row int, column int) float {
	return m[row][column]
}

func (m Matrix) Equal(o Matrix) bool {
	if !m.sameDimensions(o) {
		return false
	}
	for r := 0; r < len(m); r++ {
		r1 := m[r]
		r2 := o[r]
		for x := 0; x < len(r1); x++ {
			if !floatsEqual(r1[x], r2[x]) {
				return false
			}
		}
	}
	return true
}

func (m Matrix) Multiply(o Matrix) (res Matrix) {
	if !m.sameDimensions(o) {
		panic("can't multiply matrices with different dimensions")
	}
	res = NewMatrix(m.Rows(), m.Cols())
	for ri := 0; ri < m.Rows(); ri++ {
		for ci := 0; ci < m.Cols(); ci++ {
			for lhsi := 0; lhsi < m.Cols(); lhsi++ {
				lhs := m.Get(ri, lhsi)
				rhs := o.Get(lhsi, ci)
				res[ri][ci] += (lhs * rhs)
			}
		}
	}
	return
}

func (m Matrix) MultiplyTuple4(t1 Tuple4) (res Tuple4) {
	if m.Rows() != 4 || m.Cols() != 4 {
		panic("must be a 4x4")
	}
	vals := [4]float{}
	for ri := 0; ri < m.Rows(); ri++ {
		vals[ri] += m.Get(ri, 0) * t1.x
		vals[ri] += m.Get(ri, 1) * t1.y
		vals[ri] += m.Get(ri, 2) * t1.z
		vals[ri] += m.Get(ri, 3) * t1.w
	}
	return NewTuple(vals[0], vals[1], vals[2], vals[3])
}

func (m Matrix) sameDimensions(o Matrix) bool {
	return m.Rows() == o.Rows() && m.Cols() == o.Cols()
}
