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

#### 5. unrar

Extracts RAR files to their containing directories. Can process a single RAR file or recursively process all RAR files in a directory.

```bash
filekit unrar <rar_file_or_directory> [-r]
```

**Flags:**
- `-r`: Process directories recursively (only applicable when target is a directory)

**Arguments:**
- `rar_file_or_directory`: Path to a RAR file or directory containing RAR files (required)

**Examples:**
```bash
# Extract a single RAR file
filekit unrar archive.rar

# Extract all RAR files in current directory
filekit unrar .

# Recursively extract all RAR files in a directory and its subdirectories
filekit unrar /path/to/archives -r

# Extract all RAR files in a specific directory (non-recursive)
filekit unrar /downloads/
```

**Requirements:**
- The `unrar` utility must be installed on your system
- **macOS**: `brew install unrar`
- **Ubuntu/Debian**: `sudo apt install unrar`
- **CentOS/RHEL**: `sudo yum install unrar` or `sudo dnf install unrar`

**unrar behavior:**
- Extracts files to the same directory as the RAR file
- Overwrites existing files (uses `-o+` flag)
- Processes only `.rar` files (case-insensitive)
- Shows progress for each file being extracted
- Reports total number of files extracted

**Use cases:**
- Batch extract downloaded archive collections
- Process backup archives
- Extract game or software distributions
- Automate archive extraction workflows

#### 6. remove-files

Removes files matching a specified pattern from a directory. Shows confirmation before deletion for safety.

```bash
filekit remove-files <directory> -pattern="*.ext" [-recursive]
```

**Flags:**
- `-pattern`: File pattern to match (e.g., '*.rar', '*.tmp', '*.log') (required)
- `-recursive`: Process directories recursively (optional)

**Arguments:**
- `directory`: Directory to process (optional, defaults to current directory)

**Examples:**
```bash
# Remove all .rar files from current directory
filekit remove-files . -pattern="*.rar"

# Remove all .tmp files recursively from a directory
filekit remove-files /path/to/cleanup -pattern="*.tmp" -recursive

# Remove all .log files from current directory (directory argument omitted)
filekit remove-files -pattern="*.log"

# Remove all backup files with .bak extension recursively
filekit remove-files /project -pattern="*.bak" -recursive
```

**remove-files behavior:**
- Searches for files matching the specified pattern
- Shows a list of files to be deleted (first 10 files, then count for remaining)
- Asks for user confirmation before deletion (y/N prompt)
- Deletes files only after confirmation
- Reports number of successfully deleted files
- Shows errors for files that couldn't be deleted

**Safety features:**
- **Confirmation required**: Always shows what will be deleted and asks for confirmation
- **Pattern validation**: Validates the pattern syntax before searching
- **Error reporting**: Reports which files couldn't be deleted and why
- **Preview mode**: Shows files to be deleted before any action

**Use cases:**
- Clean up temporary files
- Remove old log files
- Delete backup files
- Clean download directories
- Remove compilation artifacts

## Project Structure

```
filekit/
├── main.go                    # Main entry point and command routing
├── cmd/                       # Command handlers
│   ├── replace_in_names.go   # rename-replace command handler
│   ├── create_rand_files.go  # create-rand-files command handler
│   ├── folderify.go          # folderify command handler
│   ├── deep_compare.go       # deep-compare command handler
│   ├── unrar.go              # unrar command handler
│   └── remove_files.go       # remove-files command handler
├── internal/                  # Internal packages (implementation logic)
│   ├── rename/               # File renaming logic
│   │   └── rename.go
│   ├── generator/            # Random file generation logic
│   │   └── generator.go
│   ├── folderify/           # Folderify logic
│   │   └── folderify.go
│   ├── compare/             # Directory comparison logic
│   │   └── compare.go
│   ├── unrar/               # RAR extraction logic
│   │   └── unrar.go
│   └── remover/             # File removal logic
│       └── remover.go
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
