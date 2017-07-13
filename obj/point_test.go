package obj

import (
	"bytes"
	"fmt"
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
	{"-2/-2/-4" /*-*/, "" /*----------*/, Point{&Vertex{1, 9, 9, 9}, &Normal{2, 1, 3, 4}, &TextureCoord{3, 9, 1, 2}}},
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
		name := fmt.Sprintf("parsePoint(%s)", test.Items)
		t.Run(name, func(t *testing.T) {
			p, err := parsePoint(test.Items, &dummyObject)

			failed := false
			failed = failed || !compareErrors(err, test.Error)
			failed = failed || !comparePoints(&test.Point, p)

			if failed {
				t.Errorf("got %v, '%v', expected %v, '%v'", p, err, test.Point, test.Error)
			}

		})
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
		name := fmt.Sprintf("writePoint(%v, wr)", test.Point)
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			err := writePoint(&test.Point, &buf)

			body := string(buf.Bytes())

			failed := false
			failed = failed || !compareErrors(err, test.Error)
			failed = failed || test.Output != body

			if failed {
				t.Errorf("got '%v', '%v', expected '%v', '%v'", body, err, test.Output, test.Error)
			}
		})
	}

}
