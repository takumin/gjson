package filelist_test

import (
	"sort"
	"testing"
	"testing/fstest"

	"github.com/takumin/gjson/internal/filelist"
)

func TestFilelist(t *testing.T) {
	testFS := fstest.MapFS{
		"test.go":              &fstest.MapFile{Data: []byte("")},
		"test.txt":             &fstest.MapFile{Data: []byte("")},
		"test.inc.txt":         &fstest.MapFile{Data: []byte("")},
		"test.exc.txt":         &fstest.MapFile{Data: []byte("")},
		"foo/test.go":          &fstest.MapFile{Data: []byte("")},
		"foo/test.txt":         &fstest.MapFile{Data: []byte("")},
		"foo/test.inc.txt":     &fstest.MapFile{Data: []byte("")},
		"foo/test.exc.txt":     &fstest.MapFile{Data: []byte("")},
		"foo/bar/test.go":      &fstest.MapFile{Data: []byte("")},
		"foo/bar/test.txt":     &fstest.MapFile{Data: []byte("")},
		"foo/bar/test.inc.txt": &fstest.MapFile{Data: []byte("")},
		"foo/bar/test.exc.txt": &fstest.MapFile{Data: []byte("")},
	}

	tests := []struct {
		name     string
		root     string
		includes []string
		excludes []string
		want     []string
	}{
		{
			name:     "Include go, Exclude None",
			root:     ".",
			includes: []string{"**/*.go"},
			excludes: []string{},
			want:     []string{"test.go", "foo/test.go", "foo/bar/test.go"},
		},
		{
			name:     "Include txt, Exclude exc.txt",
			root:     ".",
			includes: []string{"**/*.txt"},
			excludes: []string{"**/*.exc.txt"},
			want:     []string{"test.txt", "test.inc.txt", "foo/test.txt", "foo/test.inc.txt", "foo/bar/test.txt", "foo/bar/test.inc.txt"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filelist.Filelist(testFS, tt.root, tt.includes, tt.excludes)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			sort.Strings(got)
			sort.Strings(tt.want)

			if len(got) != len(tt.want) {
				t.Errorf("Length of result slice differs, got %d entries, want %d entries", len(got), len(tt.want))
				return
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Filelist() got = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
