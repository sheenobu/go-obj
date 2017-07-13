package obj

import (
	"fmt"

	"github.com/pkg/errors"
)

// A ReaderOption is a functional option
// which updates the reader object
type ReaderOption func(r *stdReader)

// WithType adds a new `Handler` to the `Reader`, denoted by
// the key and descriptive name
func WithType(k string, desc string, h Handler) ReaderOption {
	return func(r *stdReader) {
		r.router[k] = parseErrorHandler(desc, h)
	}
}

// WithUnknown adds a new `Handler` to the `Reader` for unknown lines/token types
func WithUnknown(h Handler) ReaderOption {
	return func(r *stdReader) {
		r.unknown = parseErrorHandler("unknown element", h)
	}
}

// ErrorHandler is a handler which returns an error
func ErrorHandler(o *Object, token string, rest ...string) error {
	return errors.New("error from error handler")
}

// WithRestrictedTypes registeres an unknown handler
// that errors out only if the key is outside the given
// list.
func WithRestrictedTypes(typ ...string) ReaderOption {
	var m = make(map[string]bool)
	for _, t := range typ {
		m[t] = true
	}
	return WithUnknown(func(o *Object, token string, rest ...string) error {
		_, ok := m[token]
		if ok {
			return nil
		}
		return errors.New("element type restricted")
	})
}

func parseErrorHandler(desc string, h Handler) Handler {
	return func(o *Object, token string, rest ...string) error {
		err := h(o, token, rest...)
		if err != nil {
			return wrapParseErrors(fmt.Sprintf("%s (%s)", desc, token), err)
		}
		return nil
	}
}
