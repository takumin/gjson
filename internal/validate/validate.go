package validate

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type rdjsonlPosition struct {
	Line   int `json:"line,omitempty"`
	Column int `json:"column,omitempty"`
}

type rdjsonlRange struct {
	Start *rdjsonlPosition `json:"start,omitempty"`
	End   *rdjsonlPosition `json:"end,omitempty"`
}

type rdjsonlLocation struct {
	Path  string        `json:"path,omitempty"`
	Range *rdjsonlRange `json:"range,omitempty"`
}

type rdjsonl struct {
	Message  string           `json:"message,omitempty"`
	Location *rdjsonlLocation `json:"location,omitempty"`
	Severity string           `json:"severity,omitempty"`
}

type position struct {
	Filename string
	Offset   int
	Line     int
	Column   int
}

func (pos *position) String() string {
	return fmt.Sprintf("%s:%d:%d", pos.Filename, pos.Line, pos.Column)
}

func getPosition(file string, offset int) (pos *position, err error) {
	fd, err := os.OpenFile(filepath.Clean(file), os.O_RDONLY, os.ModePerm)
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

	pos = &position{
		Filename: filepath.Clean(file),
		Offset:   offset,
		Line:     l,
		Column:   offset - c,
	}

	return pos, nil
}

func Validate(file string) ([]byte, error) {
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

	rdjsonl := rdjsonl{
		Message: errjson.Error(),
		Location: &rdjsonlLocation{
			Path: file,
			Range: &rdjsonlRange{
				Start: &rdjsonlPosition{
					Line:   pos.Line,
					Column: pos.Column,
				},
			},
		},
		Severity: "ERROR",
	}

	buf, err := json.Marshal(rdjsonl)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
