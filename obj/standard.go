package obj

import "io"

// StandardSet is the standard set of wavefront object types. Not all are
// implemented but all are allowed within a `StandardReader`
var StandardSet = []string{"o", "g", "s", "mtlib", "usemtl", "v", "vn", "vp", "#"}

// NewStandardReader returns a reader which supports a set of
// given types. Any others generate errors.
func NewStandardReader(r io.Reader) Reader {
	return NewReader(r, WithRestrictedTypes(StandardSet...))
}
