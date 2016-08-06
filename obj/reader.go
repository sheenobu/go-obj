package obj

import (
	"bufio"
	"io"
)

// Reader is responsible for reading the Object
type Reader interface {
	Read() (*Object, error)
}

// NewReader creates a new reader from the given io reader
func NewReader(r io.Reader) Reader {
	return &stdReader{r}
}

type stdReader struct {
	r io.Reader
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

	//TODO: cyclomic complexity is 11. Would a 'router' be better here?

	if len(line) == 0 {
		return nil
	}

	tokens := splitByToken(line, ' ')
	rest := tokens[1:]

	switch string(tokens[0]) {
	case "#":
		// skip comments
		return nil
	case "o":
		o.Name = string(tokens[1])
	case "v":
		v, err := parseVertex(rest)
		if err != nil {
			return wrapParseError(err, lineNumber, "vertex (v)")
		}
		o.Vertices = append(o.Vertices, v)
		return nil
	case "vn":
		vn, err := parseNormal(rest)
		if err != nil {
			return wrapParseError(err, lineNumber, "vertexNormal (vn)")
		}
		o.Normals = append(o.Normals, vn)
		return nil
	case "vt":
		vt, err := parseTextCoord(rest)
		if err != nil {
			return wrapParseError(err, lineNumber, "textureCoordinate (vt)")
		}

		o.Textures = append(o.Textures, vt)
	case "f":
		f, err := parseFace(rest, o)
		if err != nil {
			return wrapParseError(err, lineNumber, "face (f)")
		}
		o.Faces = append(o.Faces, f)
		return nil
	default:
		//	fmt.Printf("ignoring token: %s\n", tokens[0])
		return nil
	}

	return nil
}
