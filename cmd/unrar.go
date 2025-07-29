package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"filekit/internal/unrar"
)

// ExecuteUnrar handles the unrar command
func ExecuteUnrar(args []string) {
	fs := flag.NewFlagSet("unrar", flag.ExitOnError)
	recursive := fs.Bool("r", false, "Process directories recursively")

	fs.Parse(args)

	// Check if unrar utility is installed
	if err := unrar.CheckUnrarInstalled(); err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("To install unrar on macOS: brew install unrar")
		fmt.Println("To install unrar on Ubuntu/Debian: sudo apt install unrar")
		os.Exit(1)
	}

	// Get the target (file or directory)
	if fs.NArg() == 0 {
		fmt.Println("Error: Please specify a RAR file or directory containing RAR files")
		fmt.Println("Usage: filekit unrar <rar_file_or_directory> [-r]")
		os.Exit(1)
	}

	target := fs.Arg(0)

	// Check if target exists
	info, err := os.Stat(target)
	if os.IsNotExist(err) {
		fmt.Printf("Error: Target does not exist: %s\n", target)
		os.Exit(1)
	}

	if info.IsDir() {
		// Process directory
		fmt.Printf("Processing directory: %s\n", target)
		if *recursive {
			fmt.Println("Recursive mode enabled")
		}

		count, err := unrar.ProcessDirectory(target, *recursive)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if count == 0 {
			fmt.Println("No RAR files found to extract")
		} else {
			fmt.Printf("Successfully extracted %d RAR file(s)\n", count)
		}
	} else {
		// Process single file
		absPath, err := filepath.Abs(target)
		if err != nil {
			fmt.Printf("Error getting absolute path: %v\n", err)
			os.Exit(1)
		}

		err = unrar.UnrarFile(absPath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Extraction completed successfully")
	}
}
