package obj

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// A Point is a single point on a face
type Point struct {
	Vertex  *Vertex
	Normal  *Normal
	Texture *TextureCoord
}

func parseIndex(i string, length int) (idx int64, err error) {
	idx, err = strconv.ParseInt(i, 10, 64)
	if err != nil {
		return
	}
	if idx < 0 {
		// Negative indices are relative to the end
		idx = int64(length) + idx
	} else {
		// Positive indices start at 1
		idx = idx - 1
	}
	return
}

func parsePoint(i string, o *Object) (p *Point, err error) {
	p = &Point{}

	var vertexNormalIndex int64
	var vertexTextureIndex int64
	var vertexIndex int64

	vertexItems := strings.Split(i, "/")

	if vertexIndex, err = parseIndex(vertexItems[0], len(o.Vertices)); err != nil {
		return
	}

	p.Vertex = &o.Vertices[vertexIndex]

	if len(vertexItems) > 1 && len(vertexItems[1]) != 0 {
		if vertexTextureIndex, err = parseIndex(vertexItems[1], len(o.Textures)); err != nil {
			return
		}
		p.Texture = &o.Textures[vertexTextureIndex]
	}

	if len(vertexItems) > 2 && len(vertexItems[2]) != 0 {
		if vertexNormalIndex, err = parseIndex(vertexItems[2], len(o.Normals)); err != nil {
			return
		}
		p.Normal = &o.Normals[vertexNormalIndex]
	}

	return
}

func writePoint(p *Point, w io.Writer) (err error) {
	if _, err = w.Write([]byte(fmt.Sprintf("%d", p.Vertex.Index))); err != nil {
		return
	}

	if p.Texture != nil {
		if _, err = w.Write([]byte(fmt.Sprintf("/%d", p.Texture.Index))); err != nil {
			return
		}
	} else if p.Normal != nil {
		if _, err = w.Write([]byte("/")); err != nil {
			return
		}
	}

	if p.Normal != nil {
		if _, err = w.Write([]byte(fmt.Sprintf("/%d", p.Normal.Index))); err != nil {
			return
		}
	}

	return
}
