package cmd

import (
	"flag"
	"fmt"
	"os"

	"filekit/internal/generator"
)

// ExecuteCreateRandFiles handles the create-rand-files command
func ExecuteCreateRandFiles(args []string) {
	fs := flag.NewFlagSet("create-rand-files", flag.ExitOnError)
	depth := fs.Int("depth", 1, "Directory depth for file creation")
	count := fs.Int("count", 5, "Number of files to create")

	fs.Parse(args)

	if *depth < 1 {
		fmt.Println("Error: depth must be at least 1")
		os.Exit(1)
	}

	if *count < 1 {
		fmt.Println("Error: count must be at least 1")
		os.Exit(1)
	}

	// Get the directory to work in (default to current directory)
	dir := "."
	if fs.NArg() > 0 {
		dir = fs.Arg(0)
	}

	err := generator.CreateRandomFiles(dir, *depth, *count)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created %d random files at depth %d in directory %s\n", *count, *depth, dir)
}
