package foldersutils

import (
	"os"
)

// Receives a folder's path and returns if it is empty.
// If it's not a folder, returns false
func IsEmpty(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if !info.IsDir() {
		return false, nil
	}

	content, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}

	return len(content) == 0, nil
}
