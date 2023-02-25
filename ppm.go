package rt

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func WritePPM(canvas Canvas, dst io.Writer) error {
	bw := bufio.NewWriter(dst)
	writer := ppmWriter{
		canvas: canvas,
		writer: bw,
		max:    255,
	}
	defer bw.Flush()
	return writer.write()
}

func WritePPMTo(canvas Canvas, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err = WritePPM(canvas, f); err != nil {
		return err
	}
	return f.Close()
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
		writevalue := func(vi int) {
			vs := fmt.Sprintf("%d", vi)
			switch {
			case l == 0:
			case 1+l+len(vs) > 70:
				l = 0
				write("\n")
			default:
				write(" ")
			}
			write(vs)
		}
		for ci := 0; ci < w.canvas.Width(); ci++ {
			px := w.canvas.PixelAt(ci, ri)
			writevalue(w.scale(px.Red()))
			writevalue(w.scale(px.Green()))
			writevalue(w.scale(px.Blue()))
		}
		w.writeString("\n")
	}
	w.writeString("\n")
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
