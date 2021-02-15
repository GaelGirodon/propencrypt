package fileutil

import (
	"errors"
	"os"
	"path/filepath"
)

// ExistsRegular checks that the given file exists and is a regular file.
func ExistsRegular(filename string) bool {
	stat, err := os.Stat(filename)
	return err == nil && stat.Mode().IsRegular()
}

// Expand returns the full list of files by interpreting glob patterns.
func Expand(filenames []string) (files []string, err error) {
	for _, f := range filenames {
		if matches, err := filepath.Glob(f); err != nil {
			return nil, errors.New("invalid file name pattern '" + f + "'")
		} else {
			files = append(files, matches...)
		}
	}
	return files, nil
}
