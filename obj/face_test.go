package obj

import (
	"bytes"
	"testing"
)

var fNullIndex = int64(-1)

var faceReadTests = []struct {
	Items stringList
	Error string
	Face  Face
}{
	{stringList{"12//1"}, "", Face{fNullIndex, []Vertex{Vertex{12, 1, 1, 1}}, Normal{1, 1, 2, 3}}},
}

func TestReadFace(t *testing.T) {

	var dummyObject Object
	dummyObject.Vertices = make([]Vertex, 14)
	dummyObject.Vertices[11] = Vertex{12, 1, 1, 1}
	dummyObject.Normals = make([]Normal, 3)
	dummyObject.Normals[0] = Normal{1, 1, 2, 3}

	for _, test := range faceReadTests {
		f, err := parseFace(test.Items.ToByteList(), &dummyObject)

		failed := false

		if !compareErrors(err, test.Error) {
			failed = true
		}

		if len(f.Vertices) == len(test.Face.Vertices) {
			for idx, vL := range f.Vertices {
				vR := test.Face.Vertices[idx]
				if !compareVertices(vL, vR) {
					failed = true
				}
			}
		} else {
			failed = true
		}

		n := f.Normal
		if !compareNormals(f.Normal, test.Face.Normal) {
			failed = true
		}

		if failed {
			t.Errorf("parseFace(%s) => %v, '%v', expected %v, '%v'", test.Items, n, err, test.Face, test.Error)
		}
	}
}

var faceWriteTests = []struct {
	Face   Face
	Output string
	Error  string
}{
	{Face{fNullIndex, []Vertex{{12, 0, 0, 0}, {13, 0, 0, 0}}, Normal{2, 0, 0, 0}}, "12//2 13//2", ""},
}

func TestWriteFace(t *testing.T) {

	for _, test := range faceWriteTests {
		var buf bytes.Buffer
		err := writeFace(&test.Face, &buf)

		failed := false

		body := string(buf.Bytes())
		if test.Output != body {
			failed = true
		}

		if test.Error == "" && err != nil {
			failed = true
		} else if err != nil && test.Error != err.Error() {
			failed = true
		}

		if failed {
			t.Errorf("writeFace(%v, wr) => '%v', '%v', expected '%v', '%v'",
				test.Face, body, err, test.Output, test.Error)
		}
	}

}
