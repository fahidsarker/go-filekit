package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"filekit/internal/rename"
)

// ExecuteReplaceInNames handles the rename-replace command
func ExecuteReplaceInNames(args []string) {
	fs := flag.NewFlagSet("rename-replace", flag.ExitOnError)
	target := fs.String("target", "", "Target string to replace in filenames")
	replaceWith := fs.String("replaceWith", "", "String to replace target with (optional, defaults to empty string to remove target)")

	fs.Parse(args)

	if *target == "" {
		fmt.Println("Error: -target flag is required")
		fs.Usage()
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

	count, err := rename.ReplaceInFilenames(absDir, *target, *replaceWith)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully renamed %d files\n", count)
}
