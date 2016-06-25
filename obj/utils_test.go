package obj

import "testing"

var splitTests = []struct {
	Input  []byte
	Token  byte
	Output [][]byte
}{
	{[]byte("hello,world,from,sheena"), ',', [][]byte{[]byte("hello"), []byte("world"), []byte("from"), []byte("sheena")}},
	{[]byte("hello,world"), ',', [][]byte{[]byte("hello"), []byte("world")}},
	{[]byte("hello"), ',', [][]byte{[]byte("hello")}},
	{[]byte("12//4"), '/', [][]byte{[]byte("12"), []byte(""), []byte("4")}},
}

func TestSplitByToken(t *testing.T) {
	for _, test := range splitTests {
		res := splitByToken(test.Input, test.Token)

		failed := false

		if len(res) != len(test.Output) {
			failed = true
		} else {
			for i, item := range res {
				if string(item) != string(test.Output[i]) {
					failed = true
				}
			}
		}

		if failed {
			t.Errorf("splitByToken([]byte(%s), '%c') => %s, expected %s",
				test.Input, test.Token, res, test.Output)
		}
	}

}
