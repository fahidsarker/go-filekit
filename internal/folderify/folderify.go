package folderify

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ProcessDirectory processes files in a directory, creating folders and moving files into them
func ProcessDirectory(dir string, recursive bool) (int, error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return 0, fmt.Errorf("failed to get absolute path: %v", err)
	}

	if recursive {
		return processRecursively(absDir)
	}

	return processSingleDirectory(absDir)
}

// processRecursively processes directories recursively
func processRecursively(dir string) (int, error) {
	totalCount := 0

	// First, collect all directories and files to avoid infinite recursion
	var dirs []string
	var files []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, fmt.Errorf("failed to read directory %s: %v", dir, err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			dirs = append(dirs, fullPath)
		} else {
			files = append(files, fullPath)
		}
	}

	// Process files in current directory
	count, err := processFiles(files)
	if err != nil {
		return totalCount, err
	}
	totalCount += count

	// Recursively process subdirectories
	for _, subDir := range dirs {
		count, err := processRecursively(subDir)
		if err != nil {
			return totalCount, err
		}
		totalCount += count
	}

	return totalCount, nil
}

// processSingleDirectory processes only the specified directory (non-recursive)
func processSingleDirectory(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, fmt.Errorf("failed to read directory %s: %v", dir, err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}

	return processFiles(files)
}

// processFiles processes a list of files, creating folders and moving them
func processFiles(files []string) (int, error) {
	count := 0

	for _, filePath := range files {
		err := folderifyFile(filePath)
		if err != nil {
			return count, fmt.Errorf("failed to process file %s: %v", filePath, err)
		}
		count++
	}

	return count, nil
}

// folderifyFile creates a folder for a file and moves the file into it
func folderifyFile(filePath string) error {
	dir := filepath.Dir(filePath)
	filename := filepath.Base(filePath)

	// Get filename without extension
	ext := filepath.Ext(filename)
	nameWithoutExt := strings.TrimSuffix(filename, ext)

	// Skip if filename is empty after removing extension
	if nameWithoutExt == "" {
		return fmt.Errorf("cannot create folder for file with empty name: %s", filename)
	}

	// Create folder with the same name as the file (without extension)
	folderPath := filepath.Join(dir, nameWithoutExt)
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create folder %s: %v", folderPath, err)
	}

	// Move file into the folder
	newFilePath := filepath.Join(folderPath, filename)
	err = os.Rename(filePath, newFilePath)
	if err != nil {
		return fmt.Errorf("failed to move file %s to %s: %v", filePath, newFilePath, err)
	}

	fmt.Printf("Moved: %s -> %s\n", filePath, newFilePath)
	return nil
}
