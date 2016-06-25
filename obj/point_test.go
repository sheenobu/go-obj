package obj

import (
	"bytes"
	"testing"
)

var pointReadTests = []struct {
	Items string
	Error string
	Point Point
}{
	{"1/3/2" /*-*/, "" /*----------*/, Point{&Vertex{1, 9, 9, 9}, &Normal{2, 1, 3, 4}, &TextureCoord{3, 9, 1, 2}}},
	{"1//2" /*--*/, "" /*----------*/, Point{&Vertex{1, 9, 9, 9}, &Normal{2, 1, 3, 4}, nil}},
	{"1/3" /*---*/, "" /*----------*/, Point{&Vertex{1, 9, 9, 9}, nil, &TextureCoord{3, 9, 1, 2}}},
	{"1" /*-----*/, "" /*----------*/, Point{&Vertex{1, 9, 9, 9}, nil, nil}},
}

func TestReadPoint(t *testing.T) {

	var dummyObject Object
	dummyObject.Vertices = make([]Vertex, 2)
	dummyObject.Vertices[0] = Vertex{1, 9, 9, 9}
	dummyObject.Normals = make([]Normal, 5)
	dummyObject.Normals[1] = Normal{2, 1, 3, 4}
	dummyObject.Textures = make([]TextureCoord, 4)
	dummyObject.Textures[2] = TextureCoord{3, 9, 1, 2}

	for _, test := range pointReadTests {

		p, err := parsePoint([]byte(test.Items), &dummyObject)

		failed := false

		if test.Error == "" && err != nil {
			failed = true
		} else if err != nil && test.Error != err.Error() {
			failed = true
		}

		if p != nil {
			if !compareVertices(test.Point.Vertex, p.Vertex) {
				failed = true
			}

			if !compareNormals(test.Point.Normal, p.Normal) {
				failed = true
			}

			if !compareTextureCoords(test.Point.Texture, p.Texture) {
				failed = true
			}
		} else {
			failed = true
		}

		if failed {
			t.Errorf("parsePoint(%s) => %v, '%v', expected %v, '%v'", test.Items, p, err, test.Point, test.Error)
		}
	}
}

var pointWriteTests = []struct {
	Point  Point
	Output string
	Error  string
}{
	{Point{&Vertex{1, 0, 0, 0}, nil, nil}, "1", ""},
	{Point{&Vertex{1, 0, 0, 0}, &Normal{2, 0, 0, 0}, nil}, "1//2", ""},
	{Point{&Vertex{1, 0, 0, 0}, &Normal{2, 0, 0, 0}, &TextureCoord{3, 0, 0, 0}}, "1/3/2", ""},
	{Point{&Vertex{1, 0, 0, 0}, nil, &TextureCoord{3, 0, 0, 0}}, "1/3", ""},
}

func TestWritePoint(t *testing.T) {

	for _, test := range pointWriteTests {
		var buf bytes.Buffer
		err := writePoint(&test.Point, &buf)

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
			t.Errorf("writePoint(%v, wr) => '%v', '%v', expected '%v', '%v'",
				test.Point, body, err, test.Output, test.Error)
		}
	}

}
