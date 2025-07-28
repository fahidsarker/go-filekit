package cmd

import (
	"flag"
	"fmt"
	"os"

	"filekit/internal/compare"
)

// ExecuteDeepCompare handles the deep-compare command
func ExecuteDeepCompare(args []string) {
	fs := flag.NewFlagSet("deep-compare", flag.ExitOnError)
	verbose := fs.Bool("verbose", false, "Show detailed comparison results")

	fs.Parse(args)

	// Need exactly 2 directories to compare
	if fs.NArg() != 2 {
		fmt.Println("Error: deep-compare requires exactly 2 directories to compare")
		fmt.Println("Usage: tools deep-compare [-verbose] <directory1> <directory2>")
		os.Exit(1)
	}

	dir1 := fs.Arg(0)
	dir2 := fs.Arg(1)

	result, err := compare.DeepCompare(dir1, dir2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if *verbose {
		compare.PrintResult(result)
	} else {
		if result.Identical {
			fmt.Println("✅ Directories are identical!")
		} else {
			fmt.Println("❌ Directories have differences:")
			if len(result.OnlyInDir1) > 0 {
				fmt.Printf("  - %d items only in first directory\n", len(result.OnlyInDir1))
			}
			if len(result.OnlyInDir2) > 0 {
				fmt.Printf("  - %d items only in second directory\n", len(result.OnlyInDir2))
			}
			if len(result.ModTimeDiffs) > 0 {
				fmt.Printf("  - %d files with different modification times\n", len(result.ModTimeDiffs))
			}
			fmt.Println("Use -verbose flag for detailed comparison")
		}
	}

	fmt.Printf("Total files: %d, Total directories: %d\n", result.TotalFiles, result.TotalDirs)
}
