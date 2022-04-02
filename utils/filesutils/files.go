package filesutils

import (
	"os"
	"path/filepath"
)

const fileMode = 0666

func WriteToFile(path string, content []byte) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModeDir)
	if err != nil {
		return err
	}

	return os.WriteFile(path, content, fileMode)
}
