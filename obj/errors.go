package obj

import "fmt"

// ParseError represents an error in parsing
type ParseError struct {
	LineNumber int64
	ItemType   string
	Err        []error
}

func (err *ParseError) Error() string {
	return fmt.Sprintf("Parse error at line %d for %s: %s",
		err.LineNumber,
		err.ItemType,
		err.Err)
}

func wrapParseErrors(lineNumber int64, itemType string, err ...error) error {
	return &ParseError{
		LineNumber: lineNumber,
		ItemType:   itemType,
		Err:        err,
	}
}
