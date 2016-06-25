package obj

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

// A Face is a collection of Vertices and a Normal grouped together
type Face struct {
	Index    int64
	Vertices []Vertex
	Normal   Normal //TODO: assumption, each face has one normal
}

func parseFace(items [][]byte, o *Object) (f Face, err error) {

	var lastVertexNormal int64 = -1
	var vertexIndex int64
	var vertexNormalIndex int64

	for _, i := range items {
		vertexItems := splitByToken(i, '/')
		if vertexIndex, err = strconv.ParseInt(string(vertexItems[0]), 10, 64); err != nil {
			return
		}
		if vertexNormalIndex, err = strconv.ParseInt(string(vertexItems[2]), 10, 64); err != nil {
			return
		}

		f.Vertices = append(f.Vertices, o.Vertices[vertexIndex-1])

		if lastVertexNormal > 0 {
			if lastVertexNormal != vertexNormalIndex {
				err = errors.New("Face: Mismatched normal for face!")
				return
			}
		}

		lastVertexNormal = vertexNormalIndex
	}

	f.Normal = o.Normals[lastVertexNormal-1]

	return
}

func writeFace(f *Face, w io.Writer) error {
	for idx, v := range f.Vertices {
		if _, err := w.Write([]byte(fmt.Sprintf("%d//%d", v.Index, f.Normal.Index))); err != nil {
			return err
		}
		if len(f.Vertices) != idx+1 {
			if _, err := w.Write([]byte{' '}); err != nil {
				return err
			}
		}
	}

	return nil
}
