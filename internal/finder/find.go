package finder

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var ErrEmptyPath = errors.New("empty path")

func Resolve(path string) (string, error) {
	if path == "" {
		return "", ErrEmptyPath
	}

	fullPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	s, err := os.Stat(fullPath)
	if err != nil {
		return "", err
	}

	if s.IsDir() {
		return "", fmt.Errorf("%s points to a directory", path)
	}

	return filepath.EvalSymlinks(fullPath)
}
