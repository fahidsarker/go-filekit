package main

import (
	"fmt"
	"os"

	"filekit/cmd"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "rename-replace":
		cmd.ExecuteReplaceInNames(args)
	case "create-rand-files":
		cmd.ExecuteCreateRandFiles(args)
	case "folderify":
		cmd.ExecuteFolderify(args)
	case "deep-compare":
		cmd.ExecuteDeepCompare(args)
	case "unrar":
		cmd.ExecuteUnrar(args)
	case "remove-files":
		cmd.ExecuteRemoveFiles(args)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: tools <cmd> <flags> [directory]")
	fmt.Println("")
	fmt.Println("Available commands:")
	fmt.Println("  rename-replace -target=\"\" [-replaceWith=\"\"] [directory]")
	fmt.Println("    Renames all files by replacing target string with replaceWith string (or removes target if replaceWith not specified)")
	fmt.Println("")
	fmt.Println("  create-rand-files -depth=num -count=num [directory]")
	fmt.Println("    Creates random txt files with random names in the specified directory")
	fmt.Println("")
	fmt.Println("  folderify [-recursive] [directory]")
	fmt.Println("    Creates folders with file names (minus extension) and moves files into them")
	fmt.Println("    Use -recursive to process subdirectories")
	fmt.Println("")
	fmt.Println("  deep-compare [-verbose] <directory1> <directory2>")
	fmt.Println("    Compares two directories recursively by structure, file names, and modification times")
	fmt.Println("    Use -verbose for detailed comparison results")
	fmt.Println("")
	fmt.Println("  unrar <rar_file_or_directory> [-r]")
	fmt.Println("    Extracts RAR files to their containing directories")
	fmt.Println("    Use -r for recursive processing when target is a directory")
	fmt.Println("")
	fmt.Println("  remove-files <directory> -pattern=\"*.ext\" [-recursive]")
	fmt.Println("    Removes files matching the specified pattern")
	fmt.Println("    Shows confirmation before deletion. Use -recursive to process subdirectories")
}
