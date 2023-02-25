package rt

import (
	"fmt"
	"io"
)

func writePPM(canvas Canvas, dst io.Writer) error {
	writer := ppmWriter{
		canvas: canvas,
		writer: dst,
	}
	return writer.write()
}

type ppmWriter struct {
	canvas Canvas
	writer io.Writer
}

func (w ppmWriter) write() error {
	err := w.writeHeader()
	if err != nil {
		return err
	}
	err = w.writeBody()
	if err != nil {
		return err
	}
	return nil
}

func (w ppmWriter) writeHeader() error {
	width, height := w.canvas.Width(), w.canvas.Height()
	header := fmt.Sprintf("P3\n%d %d\n255\n", width, height)
	err := w.writeString(header)
	return err
}

func (w ppmWriter) writeBody() error {
	width, height := w.canvas.Width(), w.canvas.Height()
	header := fmt.Sprintf("P3\n%d %d\n255", width, height)
	err := w.writeString(header)
	return err
}

func (w ppmWriter) writeString(v string) error {
	_, err := w.writer.Write([]byte(v))
	return err
}
