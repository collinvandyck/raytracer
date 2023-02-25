package rt

import (
	"bufio"
	"fmt"
	"io"
)

func writePPM(canvas Canvas, dst io.Writer) error {
	bw := bufio.NewWriter(dst)
	writer := ppmWriter{
		canvas: canvas,
		writer: bw,
	}
	return writer.write()
}

type ppmWriter struct {
	canvas Canvas
	writer *bufio.Writer
}

func (w ppmWriter) write() error {
	return w.writeHeader()
}

func (w ppmWriter) writeHeader() error {
	width, height := w.canvas.Width(), w.canvas.Height()
	header := fmt.Sprintf("P3\n%d %d\n255", width, height)
	err := w.writeString(header)
	return err
}

func (w ppmWriter) writeString(v string) error {
	_, err := w.writer.WriteString(v)
	return err
}
