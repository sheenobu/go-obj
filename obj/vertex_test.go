package obj

import (
	"bytes"
	"testing"
)

var vNullIndex = int64(-1)

var vertexReadTests = []struct {
	Items  stringList
	Error  string
	Vertex Vertex
}{
	{stringList{"1", "1", "1" /*-----------------------*/}, "" /*-------------------------------*/, Vertex{vNullIndex, 1, 1, 1}},
	{stringList{"1", "1" /*----------------------------*/}, "Vertex: item length is incorrect" /**/, Vertex{vNullIndex, 0, 0, 0}},
	{stringList{"1.000000", "-1.000000", "-1.000000" /**/}, "" /*-------------------------------*/, Vertex{vNullIndex, 1, -1, -1}},
	{stringList{"0.999999", "-1.000000", "-1.000001" /**/}, "" /*-------------------------------*/, Vertex{vNullIndex, 0.999999, -1, -1.000001}},
	{stringList{"x", "-1.000000", "-1.000001" /*-------*/}, "Vertex: unable to parse X coordinate" /*---*/, Vertex{vNullIndex, 0, 0, 0}},
	{stringList{"1.000000", "y", "-1.000001" /*--------*/}, "Vertex: unable to parse Y coordinate" /*---*/, Vertex{vNullIndex, 1, 0, 0}},
	{stringList{"1.000000", "1", "z" /*----------------*/}, "Vertex: unable to parse Z coordinate" /*---*/, Vertex{vNullIndex, 1, 1, 0}},
}

func TestReadVertex(t *testing.T) {

	for _, test := range vertexReadTests {
		v, err := parseVertex(test.Items.ToByteList())

		failed := false

		if test.Error == "" && err != nil {
			failed = true
		} else if err != nil && test.Error != err.Error() {
			failed = true
		}

		if v.X != test.Vertex.X || v.Y != test.Vertex.Y || v.Z != test.Vertex.Z {
			failed = true
		}

		if failed {
			t.Errorf("parseVertex(%s) => %v, '%v', expected %v, '%v'", test.Items, v, err, test.Vertex, test.Error)
		}
	}
}

var vertexWriteTests = []struct {
	Vertex Vertex
	Output string
	Error  string
}{
	{Vertex{vNullIndex, 1, 1, 1}, "1.000000 1.000000 1.000000", ""},
	{Vertex{vNullIndex, -1, 1, 1}, "-1.000000 1.000000 1.000000", ""},
	{Vertex{vNullIndex, -1.000001, 0.999999, 1}, "-1.000001 0.999999 1.000000", ""},
}

func TestWriteVertex(t *testing.T) {

	for _, test := range vertexWriteTests {
		var buf bytes.Buffer
		err := writeVertex(&test.Vertex, &buf)

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
			t.Errorf("writeVertex(%v, wr) => '%v', '%v', expected '%v', '%v'",
				test.Vertex, body, err, test.Output, test.Error)
		}
	}

}
