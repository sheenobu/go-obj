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
	{stringList{"1", "1" /*----------------------------*/}, "item length is incorrect" /**/, Normal{nNullIndex, 0, 0, 0}},
	{stringList{"1.0000", "-1.0000", "-1.0000" /*------*/}, "" /*-------------------------------*/, Normal{nNullIndex, 1, -1, -1}},
	{stringList{"0.9999", "-1.0000", "-1.0001" /*------*/}, "" /*-------------------------------*/, Normal{nNullIndex, 0.9999, -1, -1.0001}},
	{stringList{"x", "-1.000000", "-1.000001" /*-------*/}, "unable to parse X coordinate" /*---*/, Normal{nNullIndex, 0, 0, 0}},
	{stringList{"1.0000", "y", "-1.0001" /*------------*/}, "unable to parse Y coordinate" /*---*/, Normal{nNullIndex, 1, 0, 0}},
	{stringList{"1.0000", "1", "z" /*------------------*/}, "unable to parse Z coordinate" /*---*/, Normal{nNullIndex, 1, 1, 0}},
}

func TestReadNormal(t *testing.T) {

	for _, test := range normalReadTests {
		n, err := parseNormal(test.Items.ToByteList())

		failed := false
		failed = failed || test.Error == "" && err != nil
		failed = failed || err != nil && test.Error != err.Error()
		failed = failed || (n.X != test.Normal.X || n.Y != test.Normal.Y || n.Z != test.Normal.Z)

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
		body := string(buf.Bytes())

		failed := false
		failed = failed || test.Error == "" && err != nil
		failed = failed || err != nil && test.Error != err.Error()
		failed = failed || test.Output != body

		if failed {
			t.Errorf("writeNormal(%v, wr) => '%v', '%v', expected '%v', '%v'",
				test.Normal, body, err, test.Output, test.Error)
		}
	}

}
