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

func (m Matrix) Get(row int, column int) float {
	return m[row][column]
}

func (m Matrix) Equal(o Matrix) bool {
	if len(m) != len(o) {
		return false
	}
	if len(m) == 0 && len(o) == 0 {
		return true
	}
	if len(m[0]) != len(o[0]) {
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
