package obj

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

// A Normal is a vertex normal
type Normal struct {
	Index int64
	X     float64
	Y     float64
	Z     float64
}

func parseNormal(items [][]byte) (n Normal, err error) {
	if len(items) != 3 {
		err = errors.New("item length is incorrect")
		return
	}

	//TODO: check all, merge error types

	if n.X, err = strconv.ParseFloat(string(items[0]), 64); err != nil {
		err = errors.New("unable to parse X coordinate")
		return
	}
	if n.Y, err = strconv.ParseFloat(string(items[1]), 64); err != nil {
		err = errors.New("unable to parse Y coordinate")
		return
	}
	if n.Z, err = strconv.ParseFloat(string(items[2]), 64); err != nil {
		err = errors.New("unable to parse Z coordinate")
		return
	}

	return
}

func writeNormal(n *Normal, w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf("%0.4f %0.4f %0.4f", n.X, n.Y, n.Z)))
	return err
}
