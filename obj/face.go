package obj

import "io"

// A Face is a list of points
type Face struct {
	Index  int64
	Points []*Point
}

func parseFace(items [][]byte, o *Object) (f Face, err error) {
	var p *Point

	for _, i := range items {

		p, err = parsePoint(i, o)
		if err != nil {
			return
		}

		f.Points = append(f.Points, p)
	}

	return
}

func writeFace(f *Face, w io.Writer) error {
	for idx, p := range f.Points {
		if err := writePoint(p, w); err != nil {
			return err
		}

		if idx != len(f.Points)-1 {
			if _, err := w.Write([]byte{' '}); err != nil {
				return err
			}
		}
	}

	return nil
}
