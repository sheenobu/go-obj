package obj

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

// A TextureCoord defines a texture coordinate
type TextureCoord struct {
	Index int64
	U     float64
	V     float64
	W     float64
}

func parseTextCoord(items [][]byte) (vt TextureCoord, err error) {
	if len(items) < 2 {
		err = errors.New("TextureCoord: item length is incorrect")
		return
	}
	if vt.U, err = strconv.ParseFloat(string(items[0]), 64); err != nil {
		err = errors.New("TextureCoord: unable to parse U coordinate")
		return
	}
	if vt.V, err = strconv.ParseFloat(string(items[1]), 64); err != nil {
		err = errors.New("TextureCoord: unable to parse V coordinate")
		return
	}
	if len(items) == 3 {
		if vt.W, err = strconv.ParseFloat(string(items[2]), 64); err != nil {
			err = errors.New("TextureCoord: unable to parse W coordinate")
			return
		}
	}

	return
}

func writeTextCoord(vt *TextureCoord, w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf("%0.3f %0.3f %0.3f", vt.U, vt.V, vt.W)))
	return err
}
