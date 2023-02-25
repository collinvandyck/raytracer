package rt

import (
	"bufio"
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

func NewMatrixFromTable(table string) Matrix {
	s := bufio.NewScanner(strings.NewReader(table))
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if !strings.HasPrefix(line, "|") {
			continue
		}
	}
	return Matrix{}
}

func (m Matrix) Get(row int, column int) float {
	return m[row][column]
}
