# Filekit

A collection of useful command-line filekit written in Go for file manipulation and management.

## Installation

### Download Pre-built Binaries (Recommended)

Download the latest release for your operating system from the [GitHub Releases](https://github.com/fahidsarker/go-filekit/releases) page. Pre-built binaries are available for:
- Linux (x86_64, ARM64)
- macOS (x86_64, ARM64)  
- Windows (x86_64, ARM64)

After downloading, extract the binary and optionally add it to your PATH.

### Build from Source

If you prefer to build from source:

```bash
go build -o filekit .
```

## Usage

The filekit are executed using the format: `filekit <cmd> <flags> [directory]`

### Available Commands

#### 1. rename-replace

Renames all files in a directory by replacing a target string with a replacement string.

```bash
filekit rename-replace -target="old_string" [-replaceWith="new_string"] [directory]
```

**Flags:**
- `-target`: Target string to replace in filenames (required)
- `-replaceWith`: String to replace target with (optional, defaults to empty string to remove target)

**Arguments:**
- `directory`: Directory to process (optional, defaults to current directory)

**Examples:**
```bash
# Replace "test" with "prod" in all filenames in current directory
filekit rename-replace -target="test" -replaceWith="prod"

# Remove "backup_" prefix from all files in /path/to/files (replaceWith omitted)
filekit rename-replace -target="backup_" /path/to/files

# Explicitly remove "temp_" prefix using empty string
filekit rename-replace -target="temp_" -replaceWith=""
```

#### 2. create-rand-files

Creates random text files with random names at a specified directory depth.

```bash
filekit create-rand-files -depth=<number> -count=<number> [directory]
```

**Flags:**
- `-depth`: Directory depth for file creation (default: 1)
- `-count`: Number of files to create (default: 5)

**Arguments:**
- `directory`: Base directory to create files in (optional, defaults to current directory)

**Examples:**
```bash
# Create 10 random files in current directory
filekit create-rand-files -depth=1 -count=10

# Create 5 random files in nested directories under /tmp
filekit create-rand-files -depth=3 -count=5 /tmp
```

#### 3. folderify

Takes files and creates folders with the same name (minus extension), then moves each file into its corresponding folder.

```bash
filekit folderify [-recursive] [directory]
```

**Flags:**
- `-recursive`: Process subdirectories recursively (optional)

**Arguments:**
- `directory`: Directory to process (optional, defaults to current directory)

**Examples:**
```bash
# Folderify files in current directory only
filekit folderify

# Folderify files recursively in all subdirectories
filekit folderify -recursive

# Folderify files in specific directory
filekit folderify /path/to/files

# Folderify files recursively in specific directory
filekit folderify -recursive /path/to/files
```

**Example of folderify behavior:**
Before:
```
document.pdf
image.jpg
script.sh
```

After:
```
document/
  document.pdf
image/
  image.jpg
script/
  script.sh
```

#### 4. deep-compare

Compares two directories recursively to validate if they have the same structure and files. The comparison checks file names, directory structure, and modification times.

```bash
filekit deep-compare [-verbose] <directory1> <directory2>
```

**Flags:**
- `-verbose`: Show detailed comparison results including specific differences (optional)

**Arguments:**
- `directory1`: First directory to compare (required)
- `directory2`: Second directory to compare (required)

**Examples:**
```bash
# Compare two directories with basic output
filekit deep-compare /path/to/dir1 /path/to/dir2

# Compare with detailed differences shown
filekit deep-compare -verbose /backup/folder /current/folder

# Compare current directory with another directory
filekit deep-compare . /backup/current-dir
```

**deep-compare behavior:**
- ✅ **Identical**: Both directories have the same structure, files, and modification times
- ❌ **Different**: Shows summary of differences:
  - Files/directories only in first directory
  - Files/directories only in second directory  
  - Files with different modification times
  - Type mismatches (file vs directory with same name)

**Use cases:**
- Verify backup integrity
- Check if directories are synchronized
- Validate file migration results
- Compare before/after states

## Project Structure

```
filekit/
├── main.go                    # Main entry point and command routing
├── cmd/                       # Command handlers
│   ├── replace_in_names.go   # rename-replace command handler
│   ├── create_rand_files.go  # create-rand-files command handler
│   ├── folderify.go          # folderify command handler
│   └── deep_compare.go       # deep-compare command handler
├── internal/                  # Internal packages (implementation logic)
│   ├── rename/               # File renaming logic
│   │   └── rename.go
│   ├── generator/            # Random file generation logic
│   │   └── generator.go
│   ├── folderify/           # Folderify logic
│   │   └── folderify.go
│   └── compare/             # Directory comparison logic
│       └── compare.go
├── go.mod                    # Go module definition
└── README.md                # This file
```

## Error Handling

- All commands include proper error handling and validation
- Invalid flags or missing required parameters will show usage information
- File operation errors are reported with descriptive messages

## Development

To add new tools:

1. Add the command case to `main.go`
2. Create a new handler in the `cmd/` package
3. Implement the logic in a new package under `internal/`
4. Update this README with documentation

## License

This project is open source and available under the MIT License.
