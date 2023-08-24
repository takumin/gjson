package report_test

import (
	"reflect"
	"testing"

	"github.com/takumin/gjson/internal/report"
)

func TestReviewdogDiagnosticJSONLines(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		message string
		line    int
		column  int
		want    []byte
	}{
		{
			name:    "No Line and Column",
			path:    "invalid.yaml",
			message: "invalid",
			line:    0,
			column:  0,
			want:    []byte(`{"message":"invalid","location":{"path":"invalid.yaml"},"severity":"ERROR"}`),
		},
		{
			name:    "Set Line No Column",
			path:    "invalid.yaml",
			message: "invalid",
			line:    1,
			column:  0,
			want:    []byte(`{"message":"invalid","location":{"path":"invalid.yaml","range":{"start":{"line":1}}},"severity":"ERROR"}`),
		},
		{
			name:    "Set Line and Column",
			path:    "invalid.yaml",
			message: "invalid",
			line:    1,
			column:  1,
			want:    []byte(`{"message":"invalid","location":{"path":"invalid.yaml","range":{"start":{"line":1,"column":1}}},"severity":"ERROR"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := report.ReviewdogDiagnosticJSONLines(tt.path, tt.message, tt.line, tt.column)
			if err != nil {
				t.Errorf("expected error to be 'nil', but got '%s'", err)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Error("expected json to be equal, but got not equal")
			}
		})
	}
}
