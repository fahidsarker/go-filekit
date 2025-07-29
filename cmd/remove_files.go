package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"filekit/internal/remover"
)

// ExecuteRemoveFiles handles the remove-files command
func ExecuteRemoveFiles(args []string) {
	fs := flag.NewFlagSet("remove-files", flag.ExitOnError)
	pattern := fs.String("pattern", "", "File pattern to match (e.g., '*.rar', '*.tmp')")
	recursive := fs.Bool("recursive", false, "Process directories recursively")

	fs.Parse(args)

	if *pattern == "" {
		fmt.Println("Error: -pattern flag is required")
		fmt.Println("Usage: filekit remove-files <directory> -pattern=\"*.rar\" [-recursive]")
		os.Exit(1)
	}

	// Get the directory to process (default to current directory)
	dir := "."
	if fs.NArg() > 0 {
		dir = fs.Arg(0)
	}

	// Convert to absolute path
	absDir, err := filepath.Abs(dir)
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		os.Exit(1)
	}

	// Check if directory exists
	if _, err := os.Stat(absDir); os.IsNotExist(err) {
		fmt.Printf("Error: Directory does not exist: %s\n", absDir)
		os.Exit(1)
	}

	// Find matching files
	fmt.Printf("Searching for files matching pattern '%s' in directory: %s\n", *pattern, absDir)
	if *recursive {
		fmt.Println("Recursive mode enabled")
	}

	files, err := remover.FindMatchingFiles(absDir, *pattern, *recursive)
	if err != nil {
		fmt.Printf("Error finding files: %v\n", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Printf("No files found matching pattern '%s'\n", *pattern)
		return
	}

	// Show confirmation
	fmt.Printf("\nFound %d file(s) matching pattern '%s':\n", len(files), *pattern)
	for i, file := range files {
		if i < 10 { // Show first 10 files
			fmt.Printf("  %s\n", file)
		} else if i == 10 {
			fmt.Printf("  ... and %d more files\n", len(files)-10)
			break
		}
	}

	// Ask for confirmation
	fmt.Printf("\nAre you sure you want to delete these %d file(s)? (y/N): ", len(files))
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	response = strings.TrimSpace(strings.ToLower(response))
	if response != "y" && response != "yes" {
		fmt.Println("Operation cancelled")
		return
	}

	// Delete files
	deletedCount, err := remover.DeleteFiles(files)
	if err != nil {
		fmt.Printf("Error during deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully deleted %d file(s)\n", deletedCount)
}
