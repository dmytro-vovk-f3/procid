//go:build linux

package detector

import (
	"path/filepath"
)

func FindPID(binaryName string) (string, error) {
	files, err := filepath.Glob("/proc/*/exe")
	if err != nil {
		return "", err
	}

	for i := range files {
		if files[i][6] < '0' || files[i][6] > '9' {
			continue
		}

		file, err := filepath.EvalSymlinks(files[i])
		if err != nil {
			continue
		}

		if file == binaryName {
			return files[i][6 : len(files[i])-4], nil
		}
	}

	return "", nil
}
