package parser

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
)

type ParseError struct {
	Filename string
	Offset   int
	Line     int
	Column   int
	Message  string
}

func getPosition(file string, offset int) (pos *ParseError, err error) {
	fd, err := os.OpenFile(filepath.Clean(file), os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := fd.Close(); cerr != nil {
			err = cerr
		}
	}()

	r := bufio.NewScanner(fd)
	n := 0
	r.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)
		n += advance
		return advance, token, err
	})

	l := 0
	c := 0
	for r.Scan() {
		l++
		if n >= offset {
			break
		}
		c = n
	}

	pos = &ParseError{
		Filename: filepath.Clean(file),
		Offset:   offset,
		Line:     l,
		Column:   offset - c,
	}

	return pos, nil
}

func Parse(file string) (*ParseError, error) {
	data, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return nil, err
	}
	if json.Valid(data) {
		return nil, nil
	}

	var offset int64
	var errjson error
	var val struct{}
	if err := json.Unmarshal(data, &val); err != nil {
		switch err := err.(type) {
		case *json.UnmarshalTypeError:
			offset = err.Offset
			errjson = err
		case *json.SyntaxError:
			offset = err.Offset
			errjson = err
		default:
			return nil, err
		}
	}

	pos, err := getPosition(file, int(offset))
	if err != nil {
		return nil, err
	}
	pos.Message = errjson.Error()

	return pos, nil
}
