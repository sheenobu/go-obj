package obj

import (
	"bytes"
	"testing"
)

func TestReadObject(t *testing.T) {
	r := NewReader(bytes.NewBuffer([]byte(objectBody)))

	_, err := r.Read()
	if err != nil {
		t.Errorf("Expected success, got err: '%s'", err)
		return
	}
}

func TestReadBleh(t *testing.T) {
	r := NewReader(bytes.NewBuffer([]byte(blehObject)))

	_, err := r.Read()
	if err != nil {
		t.Errorf("Expected success, got err: '%s'", err)
		return
	}
}

var readLineTests = []struct {
	Line  string
	Error string
}{
	{"", ""},
	{"#", ""},
	{" #", ""},

	{"vn x", "Parse error at line 0 for vertexNormal (vn): [item length is incorrect]"},
	{"vt x", "Parse error at line 0 for textureCoordinate (vt): [item length is incorrect]"},
	{"v 0 0 0", ""},
	{"v x", "Parse error at line 0 for vertex (v): [item length is incorrect]"},
	{"v 0 x 0", "Parse error at line 0 for vertex (v): [unable to parse Y coordinate]"},

	{"vn 0 0 0", ""},

	{"f 1", ""},

	//TODO: better errors
	{"f x", "Parse error at line 0 for face (f): [strconv.ParseInt: parsing \"x\": invalid syntax]"},
	{"f 1/x/1", "Parse error at line 0 for face (f): [strconv.ParseInt: parsing \"x\": invalid syntax]"},
	{"f 1/1/y", "Parse error at line 0 for face (f): [strconv.ParseInt: parsing \"y\": invalid syntax]"},
}

func TestReadLine(t *testing.T) {
	var o Object

	o.Vertices = make([]Vertex, 10)
	o.Textures = make([]TextureCoord, 10)
	o.Normals = make([]Normal, 10)

	r := NewReader(nil).(*stdReader)

	for _, test := range readLineTests {
		err := r.readLine([]byte(test.Line), 0, &o)
		failed := false

		if err == nil && test.Error != "" {
			failed = true
		} else if err != nil && test.Error != err.Error() {
			failed = true
		}

		if failed {
			t.Errorf("readLine('%s', _) => '%s', expected '%s'", test.Line, err, test.Error)
		}
	}

}

var objectBody = `
# Blender v2.77 (sub 0) OBJ File: ''
# www.blender.org
mtllib untitled.mtl
o Cube
v 1.000000 -1.000000 -1.000000
v 1.000000 -1.000000 1.000000
v -1.000000 -1.000000 1.000000
v -1.000000 -1.000000 -1.000000
v 1.000000 1.000000 -0.999999
v 0.999999 1.000000 1.000001
v -1.000000 1.000000 1.000000
v -1.000000 1.000000 -1.000000
v 0.642888 1.000000 -0.642887
v 0.642887 1.000000 0.642888
v -0.642887 1.000000 -0.642887
v -0.642888 1.000000 0.642887
v 0.642888 0.758351 -0.642887
v 0.642887 0.758351 0.642888
v -0.642887 0.758351 -0.642887
v -0.642888 0.758351 0.642887
v 0.682532 -0.682532 1.000000
v -0.682532 -0.682532 1.000000
v 0.682531 0.682532 1.000000
v -0.682532 0.682532 1.000000
v 0.682532 -0.682532 0.718246
v -0.682532 -0.682532 0.718246
v 0.682532 0.682532 0.718247
v -0.682532 0.682532 0.718246
v 1.000000 -0.802290 -0.802290
v 1.000000 -0.802290 0.802291
v 1.000000 0.802290 -0.802290
v 0.999999 0.802290 0.802291
v 0.707367 -0.802290 -0.802290
v 0.707367 -0.802290 0.802290
v 0.707368 0.802290 -0.802290
v 0.707367 0.802290 0.802291
v -1.000000 -0.692795 0.692795
v -1.000000 -0.692795 -0.692796
v -1.000000 0.692795 0.692795
v -1.000000 0.692795 -0.692796
v -1.293880 -0.692796 0.692795
v -1.293880 -0.692796 -0.692796
v -1.293880 0.692795 0.692795
v -1.293880 0.692795 -0.692796
vn 0.0000 -1.0000 0.0000
vn 0.0000 1.0000 0.0000
vn 1.0000 0.0000 0.0000
vn -0.0000 -0.0000 1.0000
vn -1.0000 -0.0000 -0.0000
vn 0.0000 0.0000 -1.0000
usemtl Material
s off
f 1//1 2//1 3//1 4//1
f 7//2 6//2 10//2 12//2
f 5//3 6//3 28//3 27//3
f 3//4 2//4 17//4 18//4
f 8//5 4//5 34//5 36//5
f 5//6 1//6 4//6 8//6
f 12//6 10//6 14//6 16//6
f 6//2 5//2 9//2 10//2
f 8//2 7//2 12//2 11//2
f 5//2 8//2 11//2 9//2
f 13//2 15//2 16//2 14//2
f 10//5 9//5 13//5 14//5
f 11//3 12//3 16//3 15//3
f 9//4 11//4 15//4 13//4
f 17//5 19//5 23//5 21//5
f 6//4 7//4 20//4 19//4
f 7//4 3//4 18//4 20//4
f 2//4 6//4 19//4 17//4
f 21//4 23//4 24//4 22//4
f 20//3 18//3 22//3 24//3
f 18//2 17//2 21//2 22//2
f 19//1 20//1 24//1 23//1
f 27//1 28//1 32//1 31//1
f 6//3 2//3 26//3 28//3
f 1//3 5//3 27//3 25//3
f 2//3 1//3 25//3 26//3
f 29//3 31//3 32//3 30//3
f 25//4 27//4 31//4 29//4
f 28//6 26//6 30//6 32//6
f 26//2 25//2 29//2 30//2
f 36//6 34//6 38//6 40//6
f 4//5 3//5 33//5 34//5
f 3//5 7//5 35//5 33//5
f 7//5 8//5 36//5 35//5
f 37//5 39//5 40//5 38//5
f 34//1 33//1 37//1 38//1
f 35//2 36//2 40//2 39//2
f 33//4 35//4 39//4 37//4
`
