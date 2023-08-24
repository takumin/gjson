package report

import "encoding/json"

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

func ReviewdogDiagnosticJSONLines(path, message string, line, column int) ([]byte, error) {
	res := &rdjsonl{
		Message: message,
		Location: &rdjsonlLocation{
			Path: path,
		},
		Severity: "ERROR",
	}
	if line != 0 {
		res.Location.Range = &rdjsonlRange{
			Start: &rdjsonlPosition{
				Line: line,
			},
		}
		if column != 0 {
			res.Location.Range.Start.Column = column
		}
	}
	buf, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
