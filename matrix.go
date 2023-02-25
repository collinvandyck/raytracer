package rt

import (
	"bufio"
	"fmt"
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

func NewMatrixFromTable(table string) Matrix {
	s := bufio.NewScanner(strings.NewReader(table))
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if !strings.HasPrefix(line, "|") {
			continue
		}
		split := strings.Split(line, "|")
		values := []float{}
		for _, part := range split {
			part = strings.TrimSpace(part)
			if len(part) == 0 {
				continue
			}
			val, err := strconv.ParseFloat(part, 64)
			if err != nil {
				panic(err)
			}
			values = append(values, val)
		}
		fmt.Println(values)
	}
	return Matrix{}
}

func (m Matrix) Get(row int, column int) float {
	return m[row][column]
}

func splitCells() bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		words := bufio.ScanWords
		advance, token, err = words(data, atEOF)
		if err != nil {
			return
		}
		if len(token) == 1 && token[0] == '|' {
			token = nil
		}
		return
	}
}
