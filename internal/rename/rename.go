package rename

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReplaceInFilenames renames files in the given directory by replacing target string with replaceWith
func ReplaceInFilenames(dir, target, replaceWith string) (int, error) {
	count := 0

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Get the directory and filename
		dirPath := filepath.Dir(path)
		filename := info.Name()

		// Check if filename contains target string
		if strings.Contains(filename, target) {
			newFilename := strings.ReplaceAll(filename, target, replaceWith)
			newPath := filepath.Join(dirPath, newFilename)

			// Rename the file
			err := os.Rename(path, newPath)
			if err != nil {
				return fmt.Errorf("failed to rename %s to %s: %v", path, newPath, err)
			}

			fmt.Printf("Renamed: %s -> %s\n", filename, newFilename)
			count++
		}

		return nil
	})

	if err != nil {
		return count, err
	}

	return count, nil
}
