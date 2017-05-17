package obj

import (
	"bufio"
	"io"
	"strings"
)

// Reader is responsible for reading the Object
type Reader interface {
	Read() (*Object, error)
}

// NewReader creates a new reader frrouter the given io reader
func NewReader(r io.Reader) Reader {
	sr := &stdReader{
		r:      r,
		router: make(objectRouter),
	}
	sr.router["#"] = commentHandler
	sr.router["o"] = objectHandler
	sr.router["v"] = vertexHandler
	sr.router["vn"] = normalHandler
	sr.router["vt"] = textureHandler
	sr.router["f"] = faceHandler
	return sr
}

type stdReader struct {
	r      io.Reader
	router objectRouter
}

func (r *stdReader) Read() (*Object, error) {
	buf := bufio.NewReader(r.r)

	var o Object

	lineNumber := int64(1)
	for {
		line, err := buf.ReadBytes('\n')
		if err == io.EOF {
			return &o, nil
		}
		if err != nil && err != io.EOF {
			return nil, err
		}

		if err := r.readLine(line[0:len(line)-1], lineNumber, &o); err != nil {
			return nil, err
		}

		lineNumber++
	}
}

func (r *stdReader) readLine(line []byte, lineNumber int64, o *Object) error {
	if len(line) == 0 {
		return nil
	}

	tokens := splitByToken(line, ' ')

	if _, err := r.router.Route(o, tokens...); err != nil {
		return wrapLineNumber(lineNumber, err)
	}

	return nil
}

func commentHandler(o *Object, rest ...[]byte) error {
	return nil
}

func objectHandler(o *Object, rest ...[]byte) error {
	o.Name = string(rest[0])
	return nil
}

func vertexHandler(o *Object, rest ...[]byte) error {
	v, err := parseVertex(rest)
	if err != nil {
		return wrapParseErrors("vertex (v)", err)
	}
	o.Vertices = append(o.Vertices, v)
	return nil
}

func normalHandler(o *Object, rest ...[]byte) error {
	vn, err := parseNormal(rest)
	if err != nil {
		return wrapParseErrors("vertexNormal (vn)", err)
	}
	o.Normals = append(o.Normals, vn)
	return nil
}

func textureHandler(o *Object, rest ...[]byte) error {
	vt, err := parseTextCoord(rest)
	if err != nil {
		return wrapParseErrors("textureCoordinate (vt)", err)
	}

	o.Textures = append(o.Textures, vt)
	return nil
}

func faceHandler(o *Object, rest ...[]byte) error {
	f, err := parseFace(rest, o)
	if err != nil {
		return wrapParseErrors("face (f)", err)
	}
	o.Faces = append(o.Faces, f)
	return nil
}

type objectRouter map[string]func(*Object, ...[]byte) error

// Route returns true if the list of tokens has been routed, false if it has been skipped
func (router objectRouter) Route(o *Object, tokens ...[]byte) (bool, error) {
	typ := strings.TrimSpace(string(tokens[0]))
	r, ok := router[typ]
	if !ok {
		return false, nil
	}
	rest := tokens[1:]

	err := r(o, rest...)
	if err != nil {
		return false, err
	}

	return true, nil
}
