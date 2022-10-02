package filelist

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Filelist(dir string, includes, excludes []string) ([]string, error) {
	list := make([]string, 0, 65536)
	err := filepath.WalkDir(filepath.Clean(dir), func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
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
