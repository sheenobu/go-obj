package obj

import (
	"bytes"
	"fmt"
	"testing"
)

var fNullIndex = int64(-1)

var faceReadTests = []struct {
	Items stringList
	Error string
	Face  Face
}{
	{stringList{"12//1"}, "", Face{fNullIndex, []*Point{&Point{&Vertex{12, 1, 1, 1}, &Normal{1, 1, 2, 3}, nil}}}},
}

func TestReadFace(t *testing.T) {

	var dummyObject Object
	dummyObject.Vertices = make([]Vertex, 14)
	dummyObject.Vertices[11] = Vertex{12, 1, 1, 1}
	dummyObject.Normals = make([]Normal, 3)
	dummyObject.Normals[0] = Normal{1, 1, 2, 3}

	for _, test := range faceReadTests {
		name := fmt.Sprintf("parseFace(%s)", test.Items)

		t.Run(name, func(t *testing.T) {
			f, err := parseFace(test.Items.ToByteList(), &dummyObject)

			failed := false
			failed = failed || !compareErrors(err, test.Error)
			failed = failed || len(f.Points) != len(test.Face.Points)

			if !failed {
				for pidx, p := range f.Points {
					failed = failed || !comparePoints(p, test.Face.Points[pidx])
					if failed {
						break
					}
				}
			}

			if failed {
				t.Errorf("got %v, '%v', expected %v, '%v'", f, err, test.Face, test.Error)
			}
		})
	}
}

var faceWriteTests = []struct {
	Face   Face
	Output string
	Error  string
}{
	{
		Face: Face{fNullIndex,
			[]*Point{
				&Point{
					Vertex: &Vertex{12, 0, 0, 0},
					Normal: &Normal{2, 0, 0, 0},
				},
				&Point{
					Vertex: &Vertex{13, 0, 0, 0},
					Normal: &Normal{2, 0, 0, 0},
				},
			},
		},
		Output: "12//2 13//2",
		Error:  "",
	},
}

func TestWriteFace(t *testing.T) {

	for _, test := range faceWriteTests {
		name := fmt.Sprintf("writeFace(%v, wr)", test.Face)

		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			err := writeFace(&test.Face, &buf)
			body := string(buf.Bytes())

			failed := false
			failed = failed || !compareErrors(err, test.Error)
			failed = failed || test.Output != body

			if failed {
				t.Errorf("got '%v', '%v', expected '%v', '%v'",
					body, err, test.Output, test.Error)
			}
		})
	}

}
