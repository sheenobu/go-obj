package obj

import "github.com/pkg/errors"

func wrapLineNumber(lineNumber int64, err error) error {
	return errors.Wrapf(err, "error at line %d", lineNumber)
}

func wrapParseErrors(itemType string, err error) error {
	return errors.Wrapf(err, "error parsing %s", itemType)
}
