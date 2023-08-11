package filelist

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Filelist(fsys fs.FS, root string, includes, excludes []string) ([]string, error) {
	list := make([]string, 0, 65536)

	err := fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		for _, v := range excludes {
			matched, err := filepath.Match(fmt.Sprintf("*.%s", v), filepath.Base(path))
			if err != nil {
				return err
			}
			if matched {
				return nil
			}
		}

		for _, v := range includes {
			matched, err := filepath.Match(fmt.Sprintf("*.%s", v), filepath.Base(path))
			if err != nil {
				return err
			}
			if matched {
				list = append(list, path)
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
