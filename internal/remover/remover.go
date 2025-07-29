package remover

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindMatchingFiles finds all files matching the given pattern in the specified directory
func FindMatchingFiles(dir, pattern string, recursive bool) ([]string, error) {
	var matchingFiles []string

	if recursive {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				matched, err := filepath.Match(pattern, info.Name())
				if err != nil {
					return fmt.Errorf("invalid pattern '%s': %w", pattern, err)
				}
				if matched {
					matchingFiles = append(matchingFiles, path)
				}
			}
			return nil
		})
		return matchingFiles, err
	} else {
		// Non-recursive: only check files in the specified directory
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, fmt.Errorf("failed to read directory: %w", err)
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				matched, err := filepath.Match(pattern, entry.Name())
				if err != nil {
					return nil, fmt.Errorf("invalid pattern '%s': %w", pattern, err)
				}
				if matched {
					fullPath := filepath.Join(dir, entry.Name())
					matchingFiles = append(matchingFiles, fullPath)
				}
			}
		}
	}

	return matchingFiles, nil
}

// DeleteFiles deletes the specified files and returns the count of successfully deleted files
func DeleteFiles(files []string) (int, error) {
	var deletedCount int
	var errors []string

	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			errors = append(errors, fmt.Sprintf("failed to delete %s: %v", file, err))
		} else {
			deletedCount++
		}
	}

	if len(errors) > 0 {
		return deletedCount, fmt.Errorf("some files could not be deleted:\n%s", strings.Join(errors, "\n"))
	}

	return deletedCount, nil
}
