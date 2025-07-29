package unrar

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// UnrarFile extracts a single RAR file to the same directory
func UnrarFile(rarPath string) error {
	// Check if the file exists
	if _, err := os.Stat(rarPath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", rarPath)
	}

	// Check if it's a RAR file
	if !strings.HasSuffix(strings.ToLower(rarPath), ".rar") {
		return fmt.Errorf("file is not a RAR file: %s", rarPath)
	}

	// Get the directory where the RAR file is located
	dir := filepath.Dir(rarPath)

	// Use unrar command to extract
	cmd := exec.Command("unrar", "x", "-o+", rarPath, dir+"/")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Extracting: %s\n", rarPath)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to extract %s: %v", rarPath, err)
	}

	fmt.Printf("Successfully extracted: %s\n", rarPath)
	return nil
}

// ProcessDirectory processes a directory for RAR files
func ProcessDirectory(dirPath string, recursive bool) (int, error) {
	count := 0

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If not recursive and the path is in a subdirectory, skip it
		if !recursive && filepath.Dir(path) != dirPath {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Process RAR files
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".rar") {
			err := UnrarFile(path)
			if err != nil {
				fmt.Printf("Warning: %v\n", err)
			} else {
				count++
			}
		}

		return nil
	})

	if err != nil {
		return count, fmt.Errorf("error processing directory: %v", err)
	}

	return count, nil
}

// CheckUnrarInstalled checks if unrar command is available
func CheckUnrarInstalled() error {
	cmd := exec.Command("unrar")
	err := cmd.Run()

	// unrar returns exit code 10 when run without arguments, which is expected
	if exitError, ok := err.(*exec.ExitError); ok {
		if exitError.ExitCode() == 10 {
			return nil // unrar is installed
		}
	}

	// Check if the error is because the command was not found
	if err != nil && strings.Contains(err.Error(), "executable file not found") {
		return fmt.Errorf("unrar command not found - please install unrar utility")
	}

	return nil
}
