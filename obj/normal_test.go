package obj

import (
	"bytes"
	"testing"
)

var nNullIndex = int64(-1)

var normalReadTests = []struct {
	Items  stringList
	Error  string
	Normal Normal
}{
	{stringList{"1", "1", "1" /*-----------------------*/}, "" /*-------------------------------*/, Normal{nNullIndex, 1, 1, 1}},
	{stringList{"1", "1" /*----------------------------*/}, "Normal: item length is incorrect" /**/, Normal{nNullIndex, 0, 0, 0}},
	{stringList{"1.0000", "-1.0000", "-1.0000" /*------*/}, "" /*-------------------------------*/, Normal{nNullIndex, 1, -1, -1}},
	{stringList{"0.9999", "-1.0000", "-1.0001" /*------*/}, "" /*-------------------------------*/, Normal{nNullIndex, 0.9999, -1, -1.0001}},
	{stringList{"x", "-1.000000", "-1.000001" /*-------*/}, "Normal: unable to parse X coordinate" /*---*/, Normal{nNullIndex, 0, 0, 0}},
	{stringList{"1.0000", "y", "-1.0001" /*------------*/}, "Normal: unable to parse Y coordinate" /*---*/, Normal{nNullIndex, 1, 0, 0}},
	{stringList{"1.0000", "1", "z" /*------------------*/}, "Normal: unable to parse Z coordinate" /*---*/, Normal{nNullIndex, 1, 1, 0}},
}

func TestReadNormal(t *testing.T) {

	for _, test := range normalReadTests {
		n, err := parseNormal(test.Items.ToByteList())

		failed := false

		if test.Error == "" && err != nil {
			failed = true
		} else if err != nil && test.Error != err.Error() {
			failed = true
		}

		if n.X != test.Normal.X || n.Y != test.Normal.Y || n.Z != test.Normal.Z {
			failed = true
		}

		if failed {
			t.Errorf("parseNormal(%s) => %v, '%v', expected %v, '%v'", test.Items, n, err, test.Normal, test.Error)
		}
	}
}

var normalWriteTests = []struct {
	Normal Normal
	Output string
	Error  string
}{
	{Normal{nNullIndex, 1, 1, 1}, "1.0000 1.0000 1.0000", ""},
	{Normal{nNullIndex, -1, 1, 1}, "-1.0000 1.0000 1.0000", ""},
	{Normal{nNullIndex, -1.0001, 0.9999, 1}, "-1.0001 0.9999 1.0000", ""},
}

func TestWriteNormal(t *testing.T) {

	for _, test := range normalWriteTests {
		var buf bytes.Buffer
		err := writeNormal(&test.Normal, &buf)

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
			t.Errorf("writeNormal(%v, wr) => '%v', '%v', expected '%v', '%v'",
				test.Normal, body, err, test.Output, test.Error)
		}
	}

}
