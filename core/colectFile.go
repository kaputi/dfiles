package core

import (
	"os"
	"path/filepath"
	"slices"
)

// collectFiles recursively collects files from the directory and appends them to the slice
func CollectFiles(path string) ([]string, error) {
	tracked := GetTraked()
	files := []string{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && !slices.Contains(tracked, RemoveHome(path)) {
			files = append(files, path)
		}

		// TODO: check ignored files

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
