package apachelog

import (
	"io"
	"errors"
	"strings"
	"bufio"
)

// CombinedParser creates a new parser that reads from r and that parses log
// entries using the Apache Combined Log format.
func CombinedParserWithFilePos(r io.ReadSeeker,start int64) (*Parser, error) {
	if _, err := r.Seek(start, 0); err != nil {
		return nil,err
	}
	return CustomParserWithPos(r, CombinedLogFromat,start)
}

// CommonParser creates a new parser that reads from r and that parses log entries
// using the Apache Common Log format.
func CommonParserWithFilePos(r io.ReadSeeker,start int64) (*Parser, error) {
	if _, err := r.Seek(start, 0); err != nil {
		return nil,err
	}
	return CustomParserWithPos(r, CommonLogFormat,start)
}
func CustomParserWithPos(r io.Reader, format string,start int64) (*Parser, error) {
	if r == nil {
		return nil, errors.New("reader is nil")
	}
	fn, err := makeStateFn(strings.Split(format, " "))
	if err != nil {
		return nil, err
	}
	return &Parser{
		br: bufio.NewReader(r),
		fn: fn,
		Pos:start,
	}, nil
}