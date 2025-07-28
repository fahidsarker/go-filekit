package compare

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// FileInfo represents basic file information for comparison
type FileInfo struct {
	Name    string
	ModTime time.Time
	IsDir   bool
}

// ComparisonResult represents the result of directory comparison
type ComparisonResult struct {
	Identical    bool
	OnlyInDir1   []string
	OnlyInDir2   []string
	ModTimeDiffs []string
	TotalFiles   int
	TotalDirs    int
}

// DeepCompare compares two directories and their contents recursively
func DeepCompare(dir1, dir2 string) (*ComparisonResult, error) {
	absDir1, err := filepath.Abs(dir1)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for %s: %v", dir1, err)
	}

	absDir2, err := filepath.Abs(dir2)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for %s: %v", dir2, err)
	}

	// Check if both directories exist
	if _, err := os.Stat(absDir1); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory does not exist: %s", absDir1)
	}
	if _, err := os.Stat(absDir2); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory does not exist: %s", absDir2)
	}

	result := &ComparisonResult{
		Identical:    true,
		OnlyInDir1:   []string{},
		OnlyInDir2:   []string{},
		ModTimeDiffs: []string{},
	}

	err = compareDirectories(absDir1, absDir2, "", result)
	if err != nil {
		return nil, err
	}

	// Set identical to false if there are any differences
	if len(result.OnlyInDir1) > 0 || len(result.OnlyInDir2) > 0 || len(result.ModTimeDiffs) > 0 {
		result.Identical = false
	}

	return result, nil
}

// compareDirectories recursively compares two directories
func compareDirectories(dir1, dir2, relativePath string, result *ComparisonResult) error {
	entries1, err := os.ReadDir(dir1)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %v", dir1, err)
	}

	entries2, err := os.ReadDir(dir2)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %v", dir2, err)
	}

	// Create maps for easier comparison
	files1 := make(map[string]FileInfo)
	files2 := make(map[string]FileInfo)

	// Process entries from dir1
	for _, entry := range entries1 {
		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("failed to get file info for %s: %v", entry.Name(), err)
		}

		fileInfo := FileInfo{
			Name:    entry.Name(),
			ModTime: info.ModTime(),
			IsDir:   entry.IsDir(),
		}
		files1[entry.Name()] = fileInfo

		if entry.IsDir() {
			result.TotalDirs++
		} else {
			result.TotalFiles++
		}
	}

	// Process entries from dir2
	for _, entry := range entries2 {
		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("failed to get file info for %s: %v", entry.Name(), err)
		}

		fileInfo := FileInfo{
			Name:    entry.Name(),
			ModTime: info.ModTime(),
			IsDir:   entry.IsDir(),
		}
		files2[entry.Name()] = fileInfo
	}

	// Find files/dirs only in dir1
	for name, info := range files1 {
		fullPath := filepath.Join(relativePath, name)
		if _, exists := files2[name]; !exists {
			result.OnlyInDir1 = append(result.OnlyInDir1, fullPath)
		} else {
			// File exists in both, check if it's a directory and compare mod times
			info2 := files2[name]

			// Check if both are the same type (file vs directory)
			if info.IsDir != info2.IsDir {
				result.OnlyInDir1 = append(result.OnlyInDir1, fullPath+" (type mismatch)")
				result.OnlyInDir2 = append(result.OnlyInDir2, fullPath+" (type mismatch)")
				continue
			}

			// Compare modification times (only for files, not directories)
			// Allow up to 2 seconds difference in modification times
			if !info.IsDir && !modTimesEqual(info.ModTime, info2.ModTime, 2*time.Second) {
				diff := fmt.Sprintf("%s (dir1: %s, dir2: %s)",
					fullPath,
					info.ModTime.Format("2006-01-02 15:04:05"),
					info2.ModTime.Format("2006-01-02 15:04:05"))
				result.ModTimeDiffs = append(result.ModTimeDiffs, diff)
			}

			// If both are directories, recursively compare them
			if info.IsDir {
				err := compareDirectories(
					filepath.Join(dir1, name),
					filepath.Join(dir2, name),
					fullPath,
					result,
				)
				if err != nil {
					return err
				}
			}
		}
	}

	// Find files/dirs only in dir2
	for name := range files2 {
		fullPath := filepath.Join(relativePath, name)
		if _, exists := files1[name]; !exists {
			result.OnlyInDir2 = append(result.OnlyInDir2, fullPath)
		}
	}

	return nil
}

// modTimesEqual checks if two modification times are equal within a given tolerance
func modTimesEqual(t1, t2 time.Time, tolerance time.Duration) bool {
	diff := t1.Sub(t2)
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}

// PrintResult prints the comparison result in a formatted way
func PrintResult(result *ComparisonResult) {
	if result.Identical {
		fmt.Println("✅ Directories are identical!")
		fmt.Printf("Total files: %d, Total directories: %d\n", result.TotalFiles, result.TotalDirs)
		return
	}

	fmt.Println("❌ Directories have differences:")
	fmt.Printf("Total files: %d, Total directories: %d\n", result.TotalFiles, result.TotalDirs)
	fmt.Println()

	if len(result.OnlyInDir1) > 0 {
		fmt.Println("📁 Files/directories only in first directory:")
		for _, item := range result.OnlyInDir1 {
			fmt.Printf("  - %s\n", item)
		}
		fmt.Println()
	}

	if len(result.OnlyInDir2) > 0 {
		fmt.Println("📁 Files/directories only in second directory:")
		for _, item := range result.OnlyInDir2 {
			fmt.Printf("  - %s\n", item)
		}
		fmt.Println()
	}

	if len(result.ModTimeDiffs) > 0 {
		fmt.Println("⏰ Files with different modification times:")
		for _, diff := range result.ModTimeDiffs {
			fmt.Printf("  - %s\n", diff)
		}
		fmt.Println()
	}
}
