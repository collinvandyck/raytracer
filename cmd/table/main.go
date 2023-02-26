package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"rt"
	"strconv"
	"strings"
)

func main() {
	bs, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewScanner(bytes.NewReader(bs))
	var rows [][]float64
	for r.Scan() {
		line := r.Text()
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "|") || !strings.HasSuffix(line, "|") {
			continue
		}
		line = line[1 : len(line)-1]
		parts := strings.Split(line, "|")
		var row []float64
		for _, part := range parts {
			v, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, v)
		}
		rows = append(rows, row)
	}
	if len(rows) == 0 {
		log.Fatal("No rows")
	}
	m := rt.Matrix(rows)
	fmt.Println(m)

}
