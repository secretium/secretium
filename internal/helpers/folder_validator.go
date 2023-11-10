package helpers

import (
	"os"
	"path/filepath"
)

// isExistInFolder searches for a file or folder by the given name in the current folder.
func IsExistInFolder(name string, isFolder bool) bool {
	// Check if file or folder exists.
	_, err := os.Stat(filepath.Clean(name))
	if err != nil {
		return false
	}

	// Check if it is a directory.
	info, err := os.Lstat(filepath.Clean(name))
	if err != nil {
		return false
	}

	return info.IsDir() == isFolder
}
