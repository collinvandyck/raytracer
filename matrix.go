package rt

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]float

var MatrixIdentity4x4 = Matrix{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

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

func (m Matrix) Set(row int, column int, val float) {
	m[row][column] = val
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
			var val float
			for lhsi := 0; lhsi < m.Cols(); lhsi++ {
				lhs := m.Get(ri, lhsi)
				rhs := o.Get(lhsi, ci)
				val += (lhs * rhs)
			}
			res.Set(ri, ci, val)
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

func (m Matrix) Transpose() Matrix {
	rows, cols := m.Rows(), m.Cols()
	res := NewMatrix(cols, rows)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			res.Set(c, r, m.Get(r, c))
		}
	}
	return res
}

func (m Matrix) Submatrix(skipr, skipc int) Matrix {
	if m.Rows() <= 1 || m.Cols() <= 1 {
		panic("matrix must have dimension of at least2")
	}
	res := NewMatrix(m.Rows()-1, m.Cols()-1)
	for ri := 0; ri < m.Rows(); ri++ {
		if ri == skipr {
			continue
		}
		for ci := 0; ci < m.Cols(); ci++ {
			if ci == skipc {
				continue
			}
			resri := ri
			if resri > skipr {
				resri -= 1
			}
			resci := ci
			if resci > skipc {
				resci -= 1
			}
			res.Set(resri, resci, m.Get(ri, ci))
		}
	}
	return res
}

func (m Matrix) Determinant() float {
	rows, cols := m.Rows(), m.Cols()
	if rows != 2 || cols != 2 {
		panic("only 2x2 matrixes supported")
	}
	return m.Get(0, 0)*m.Get(1, 1) - m.Get(0, 1)*m.Get(1, 0)
}

func (m Matrix) Minor(row, col int) float {
	sm := m.Submatrix(row, col)
	return sm.Determinant()
}

func (m Matrix) Cofactor(row, col int) float {
	return 0
}

func (m Matrix) String() string {
	if m.Empty() {
		return "<empty>"
	}
	rows := make([][]string, 0)
	for ri := 0; ri < m.Rows(); ri++ {
		row := make([]string, 0)
		for ci := 0; ci < m.Cols(); ci++ {
			str := fmt.Sprintf("%.1f", m.Get(ri, ci))
			row = append(row, str)
		}
		rows = append(rows, row)
	}
	widths := make([]int, 0)
	for ci := 0; ci < len(rows[0]); ci++ {
		width := 0
		for ri := 0; ri < len(rows); ri++ {
			cw := len(rows[ri][ci])
			if cw > width {
				width = cw
			}
		}
		widths = append(widths, width)
	}
	widthSum := 0
	for _, w := range widths {
		widthSum += w
	}
	header := "+"
	header += strings.Repeat("-", widthSum-1)
	header += strings.Repeat("-", m.Cols()*3) // cell borders
	header += "+"
	buf := new(bytes.Buffer)
	buf.WriteString(header + "\n")
	for _, row := range rows {
		buf.WriteString("| ")
		for col, val := range row {
			pad := widths[col] - len(val)
			if pad > 0 {
				val = strings.Repeat(" ", pad) + val
			}
			buf.WriteString(val + " | ")
		}
		buf.WriteString("\n")
	}
	buf.WriteString(header + "\n")
	return buf.String()
}

func (m Matrix) Empty() bool {
	return m.Rows() == 0 || m.Cols() == 0
}

func (m Matrix) sameDimensions(o Matrix) bool {
	return m.Rows() == o.Rows() && m.Cols() == o.Cols()
}
