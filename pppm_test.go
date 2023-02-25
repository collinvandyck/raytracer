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
	})
	t.Run("Constructing the PPM pixel data", func(t *testing.T) {
		cv := NewCanvas(5, 3)
		c1 := NewColor(1.5, 0, 0)
		c2 := NewColor(0, 0.5, 0)
		c3 := NewColor(-0.5, 0, 1)

		cv.WritePixel(0, 0, c1)
		cv.WritePixel(2, 1, c2)
		cv.WritePixel(4, 2, c3)

		buf := new(bytes.Buffer)
		err := writePPM(cv, buf)
		require.NoError(t, err)

		br := bufio.NewScanner(buf)
		require.True(t, br.Scan())
		require.EqualValues(t, "P3", br.Text())
		require.True(t, br.Scan())
		require.EqualValues(t, "5 3", br.Text())
		require.True(t, br.Scan())
		require.EqualValues(t, "255", br.Text())

		require.True(t, br.Scan())
		require.Equal(t, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0", br.Text())
		require.True(t, br.Scan())
		require.Equal(t, "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0", br.Text())
		require.True(t, br.Scan())
		require.Equal(t, "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255", br.Text())

		require.Empty(t, br.Bytes())
	})
}
