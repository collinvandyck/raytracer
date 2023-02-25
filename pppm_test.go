package rt

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPPM(t *testing.T) {
	t.Run("Constructing the PPM header", func(t *testing.T) {
		c1 := NewCanvas(5, 3)
		buf := new(bytes.Buffer)
		err := writePPM(c1, buf)
		require.NoError(t, err)

		br := bufio.NewScanner(buf)
		require.True(t, br.Scan())
		require.EqualValues(t, "P3", br.Text())
		require.True(t, br.Scan())
		require.EqualValues(t, "5 3", br.Text())
		require.True(t, br.Scan())
		require.EqualValues(t, "255", br.Text())

		require.Empty(t, buf.Bytes())
	})
}
