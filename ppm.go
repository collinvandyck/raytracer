package rt

import (
	"fmt"
	"io"
	"math"
)

func writePPM(canvas Canvas, dst io.Writer) error {
	writer := ppmWriter{
		canvas: canvas,
		writer: dst,
		max:    255,
	}
	return writer.write()
}

type ppmWriter struct {
	canvas Canvas
	writer io.Writer
	max    int
	err    error
}

func (w *ppmWriter) write() error {
	w.writeHeader()
	w.writeBody()
	return w.err
}

func (w *ppmWriter) writeHeader() {
	width, height := w.canvas.Width(), w.canvas.Height()
	header := fmt.Sprintf("P3\n%d %d\n%d\n", width, height, w.max)
	w.writeString(header)
}

func (w *ppmWriter) writeBody() {
	for ri := 0; ri < w.canvas.Height(); ri++ {
		l := 0
		write := func(ps string) {
			w.writeString(ps)
			l += len(ps)
		}
		writePixel := func(vi int) {
			vs := fmt.Sprintf("%d", vi)
			switch {
			// first value, we should be fine to write it
			case l == 0:
				write(vs)
			case 1+l+len(vs) > 70:
				l = 0
				write("\n")
				write(vs)
			default:
				write(" ")
				write(vs)
			}
		}
		for ci := 0; ci < w.canvas.Width(); ci++ {
			px := w.canvas.PixelAt(ci, ri)
			writePixel(w.scale(px.Red()))
			writePixel(w.scale(px.Green()))
			writePixel(w.scale(px.Blue()))
		}
		w.writeString("\n")
	}
}

func (w *ppmWriter) writeString(v string) {
	if w.err != nil {
		return
	}
	_, err := w.writer.Write([]byte(v))
	w.err = err
}

func (w *ppmWriter) scale(val float) int {
	if val == 0 {
		return 0
	}
	v := val * float(w.max)
	if v < 0 {
		return 0
	}
	if v > float(w.max) {
		return w.max
	}
	return int(math.Round(v))
}
