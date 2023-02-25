package rt

import (
	"fmt"
	"io"
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
		for ci := 0; ci < w.canvas.Width(); ci++ {
			px := w.canvas.PixelAt(ci, ri)
			r := w.scale(px.Red())
			g := w.scale(px.Green())
			b := w.scale(px.Blue())
			w.writeString(fmt.Sprintf("%d %d %d", r, g, b))
			if ci < w.canvas.Width()-1 {
				w.writeString(" ")
			}
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
	v := int(val * float(w.max))
	if v < 0 {
		v = 0
	}
	if v > w.max {
		v = w.max
	}
	return v
}
