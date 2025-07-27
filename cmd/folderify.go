package cmd

import (
	"flag"
	"fmt"
	"os"

	"filekit/internal/folderify"
)

// ExecuteFolderify handles the folderify command
func ExecuteFolderify(args []string) {
	fs := flag.NewFlagSet("folderify", flag.ExitOnError)
	recursive := fs.Bool("recursive", false, "Process subdirectories recursively")

	fs.Parse(args)

	// Get the directory to process (default to current directory)
	dir := "."
	if fs.NArg() > 0 {
		dir = fs.Arg(0)
	}

	count, err := folderify.ProcessDirectory(dir, *recursive)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully processed %d files\n", count)
}
