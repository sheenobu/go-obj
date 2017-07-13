package obj

import (
	"bufio"
	"io"
	"strings"
)

// A Handler is a handler for a given line
type Handler func(o *Object, token string, rest ...string) error

// Reader is responsible for reading the Object
type Reader interface {
	Read() (*Object, error)
}

// NewReader creates a new reader for the given io reader
func NewReader(r io.Reader, os ...ReaderOption) Reader {
	sr := &stdReader{
		r:       r,
		router:  make(objectRouter),
		unknown: emptyUnknown,
	}
	sr.router["#"] = commentHandler
	sr.router["o"] = objectHandler
	sr.router["v"] = vertexHandler
	sr.router["vn"] = normalHandler
	sr.router["vt"] = textureHandler
	sr.router["f"] = faceHandler

	for _, o := range os {
		o(sr)
	}

	return sr
}

type stdReader struct {
	r       io.Reader
	router  objectRouter
	unknown Handler
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

		if err := r.readLine(string(line), lineNumber, &o); err != nil {
			return nil, err
		}

		lineNumber++
	}
}

func (r *stdReader) readLine(line string, lineNumber int64, o *Object) error {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return nil
	}

	var tokens []string

	for _, tok := range strings.Split(line, " ") {
		if len(tok) > 0 {
			tokens = append(tokens, tok)
		}
	}

	ok, err := r.router.Route(o, tokens...)
	if err == nil && !ok {
		// not routed
		typ := strings.TrimSpace(tokens[0]) // TODO: duplicate code from router
		rest := tokens[1:]                  // TODO: duplicate code from router
		err = r.unknown(o, typ, rest...)
	}
	if err != nil {
		return wrapLineNumber(lineNumber, err)
	}

	return nil
}

func commentHandler(o *Object, token string, rest ...string) error {
	return nil
}

func objectHandler(o *Object, token string, rest ...string) error {
	o.Name = rest[0]
	return nil
}

func vertexHandler(o *Object, token string, rest ...string) error {
	v, err := parseVertex(rest)
	if err != nil {
		return wrapParseErrors("vertex (v)", err)
	}
	o.Vertices = append(o.Vertices, v)
	return nil
}

func normalHandler(o *Object, token string, rest ...string) error {
	vn, err := parseNormal(rest)
	if err != nil {
		return wrapParseErrors("vertexNormal (vn)", err)
	}
	o.Normals = append(o.Normals, vn)
	return nil
}

func textureHandler(o *Object, token string, rest ...string) error {
	vt, err := parseTextCoord(rest)
	if err != nil {
		return wrapParseErrors("textureCoordinate (vt)", err)
	}

	o.Textures = append(o.Textures, vt)
	return nil
}

func faceHandler(o *Object, token string, rest ...string) error {
	f, err := parseFace(rest, o)
	if err != nil {
		return wrapParseErrors("face (f)", err)
	}
	o.Faces = append(o.Faces, f)
	return nil
}

type objectRouter map[string]Handler

// Route returns true if the list of tokens has been routed, false if it has been skipped
func (router objectRouter) Route(o *Object, tokens ...string) (bool, error) {
	typ := strings.TrimSpace(tokens[0])
	r, ok := router[typ]
	if !ok {
		return false, nil
	}
	rest := tokens[1:]

	err := r(o, typ, rest...)
	if err != nil {
		return false, err
	}

	return true, nil
}

func emptyUnknown(o *Object, token string, rest ...string) error {
	return nil
}
