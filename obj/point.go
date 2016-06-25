package obj

import (
	"fmt"
	"io"
	"strconv"
)

// A Point is a single point on a face
type Point struct {
	Vertex  *Vertex
	Normal  *Normal
	Texture *TextureCoord
}

func parsePoint(i []byte, o *Object) (p *Point, err error) {
	p = &Point{}

	var vertexNormalIndex int64
	var vertexTextureIndex int64
	var vertexIndex int64

	vertexItems := splitByToken(i, '/')

	if vertexIndex, err = strconv.ParseInt(string(vertexItems[0]), 10, 64); err != nil {
		return
	}

	p.Vertex = &o.Vertices[vertexIndex-1]

	if len(vertexItems) > 1 && len(vertexItems[1]) != 0 {
		if vertexTextureIndex, err = strconv.ParseInt(string(vertexItems[1]), 10, 64); err != nil {
			return
		}
		p.Texture = &o.Textures[vertexTextureIndex-1]
	}

	if len(vertexItems) > 2 && len(vertexItems[2]) != 0 {
		if vertexNormalIndex, err = strconv.ParseInt(string(vertexItems[2]), 10, 64); err != nil {
			return
		}
		p.Normal = &o.Normals[vertexNormalIndex-1]
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
