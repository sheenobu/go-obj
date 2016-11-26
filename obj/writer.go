package obj

import (
	"errors"
	"fmt"
	"io"
)

// Writer is responsible for writing an Object
type Writer interface {
	Write(o *Object) error
}

// NewWriter creates a new reader from the given io reader
func NewWriter(w io.Writer) Writer {
	return &stdWriter{w}
}

type stdWriter struct {
	w io.Writer
}

func (w *stdWriter) Write(o *Object) error {
	if o == nil {
		return errors.New("No Object given")
	}

	w.w.Write([]byte(fmt.Sprintf("o %s\n", o.Name)))

	for _, v := range o.Vertices {
		if err := writeVertex(&v, w.w); err != nil {
			return err
		}
	}

	for _, n := range o.Normals {
		if err := writeNormal(&n, w.w); err != nil {
			return err
		}
	}

	for _, tc := range o.Textures {
		if err := writeTextCoord(&tc, w.w); err != nil {
			return err
		}
	}

	for _, f := range o.Faces {
		if err := writeFace(&f, w.w); err != nil {
			return err
		}
	}

	return nil
}
