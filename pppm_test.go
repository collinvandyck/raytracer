package rt

import (
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
	})
}
