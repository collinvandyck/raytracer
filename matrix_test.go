package rt

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatrix(t *testing.T) {
	t.Run("Constructing and testing a 4x4 matrix from table", func(t *testing.T) {
		m := NewMatrixFromTable(`
		    +---------------------------+
			| 1    | 2    | 3    | 4    |
			| 5.5  | 6.5  | 7.5  | 8.5  |
			| 9    | 10   | 11   | 12   |
			| 13.5 | 14.5 | 15.5 | 16.5 |
		    +---------------------------+
		`)
		require.Equal(t, 1.0, m.Get(0, 0))
	})
}

func TestSplitScanner(t *testing.T) {
	val := "  | 1 | 2 | 3 | "
	s := bufio.NewScanner(strings.NewReader(val))
	s.Buffer(make([]byte, 2), bufio.MaxScanTokenSize)
	s.Split(splitCells())
	for s.Scan() {
		t.Log(s.Text())
	}
}
