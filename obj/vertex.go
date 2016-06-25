package obj

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Vertex represents a OBJ Vertex
type Vertex struct {
	Index int64
	X     float64
	Y     float64
	Z     float64
}

func parseVertex(items [][]byte) (v Vertex, err error) {
	if len(items) != 3 {
		err = errors.New("Vertex: item length is incorrect")
		return
	}
	if v.X, err = strconv.ParseFloat(string(items[0]), 64); err != nil {
		err = errors.New("Vertex: unable to parse X coordinate")
		return
	}
	if v.Y, err = strconv.ParseFloat(string(items[1]), 64); err != nil {
		err = errors.New("Vertex: unable to parse Y coordinate")
		return
	}
	if v.Z, err = strconv.ParseFloat(string(items[2]), 64); err != nil {
		err = errors.New("Vertex: unable to parse Z coordinate")
		return
	}

	return
}

func writeVertex(v *Vertex, w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf("%f %f %f", v.X, v.Y, v.Z)))
	return err
}
