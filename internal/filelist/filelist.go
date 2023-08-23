package filelist

import (
	"io/fs"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
)

func Filelist(fsys fs.FS, root string, includes, excludes []string) ([]string, error) {
	list := make([]string, 0, 65536)
	root = filepath.Clean(root)

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		path = filepath.Clean(path)

		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		for _, v := range excludes {
			matched, err := doublestar.PathMatch(v, path)
			if err != nil {
				return err
			}
			if matched {
				return nil
			}
		}

		for _, v := range includes {
			matched, err := doublestar.PathMatch(v, path)
			if err != nil {
				return err
			}
			if matched {
				list = append(list, filepath.Join(root, path))
				return nil
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}
