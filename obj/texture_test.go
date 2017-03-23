package obj

import (
	"bytes"
	"fmt"
	"testing"
)

var tNullIndex = int64(0)

var textureReadTests = []struct {
	Items   stringList
	Error   string
	Texture TextureCoord
}{
	{stringList{"1", "1", "1" /*---------------------*/}, "" /*-------------------------------*/, TextureCoord{tNullIndex, 1, 1, 1}},
	{stringList{"1", "1" /*--------------------------*/}, "" /*-------------------------------*/, TextureCoord{tNullIndex, 1, 1, 0}},
	{stringList{"1" /*------------------------------*/}, "item length is incorrect" /**/, TextureCoord{tNullIndex, 0, 0, 0}},
	{stringList{"1.000", "-1.000", "-1.000" /*-------*/}, "" /*-------------------------------*/, TextureCoord{tNullIndex, 1, -1, -1}},
	{stringList{"0.999", "-1.000", "-1.001" /*-------*/}, "" /*-------------------------------*/, TextureCoord{tNullIndex, 0.999, -1, -1.001}},
	{stringList{"x", "-1.000", "-1.001" /*-----------*/}, "unable to parse U coordinate" /*---*/, TextureCoord{tNullIndex, 0, 0, 0}},
	{stringList{"1.000", "y", "-1.001" /*------------*/}, "unable to parse V coordinate" /*---*/, TextureCoord{tNullIndex, 1, 0, 0}},
	{stringList{"1.000", "1", "z" /*-----------------*/}, "unable to parse W coordinate" /*---*/, TextureCoord{tNullIndex, 1, 1, 0}},
}

func TestReadTexture(t *testing.T) {

	for _, test := range textureReadTests {
		name := fmt.Sprintf("parseTextCoord(%s)", test.Items)
		t.Run(name, func(t *testing.T) {
			n, err := parseTextCoord(test.Items.ToByteList())

			failed := false
			failed = failed || !compareErrors(err, test.Error)
			failed = failed || !compareTextureCoords(&n, &test.Texture)

			if failed {
				t.Errorf("got %v, '%v', expected %v, '%v'", n, err, test.Texture, test.Error)
			}
		})
	}
}

var textureWriteTests = []struct {
	Texture TextureCoord
	Output  string
	Error   string
}{
	{TextureCoord{tNullIndex, 1, 1, 1}, "1.000 1.000 1.000", ""},
	{TextureCoord{tNullIndex, -1, 1, 1}, "-1.000 1.000 1.000", ""},
	{TextureCoord{tNullIndex, -1.001, 0.999, 1}, "-1.001 0.999 1.000", ""},
}

func TestWriteTexture(t *testing.T) {

	for _, test := range textureWriteTests {
		name := fmt.Sprintf("writeTextCoord(%v, wr)", test.Texture)
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			err := writeTextCoord(&test.Texture, &buf)
			body := string(buf.Bytes())

			failed := false
			failed = failed || (test.Error == "" && err != nil)
			failed = failed || (err != nil && test.Error != err.Error())
			failed = failed || (test.Output != body)

			if failed {
				t.Errorf("got '%v', '%v', expected '%v', '%v'",
					body, err, test.Output, test.Error)
			}
		})
	}
}
