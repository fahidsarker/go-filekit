# Filekit

A collection of useful command-line tools written in Go for file manipulation and management.

## Installation

```bash
go build -o tools .
```

## Usage

The tools are executed using the format: `tools <cmd> <flags> [directory]`

### Available Commands

#### 1. rename-replace

Renames all files in a directory by replacing a target string with a replacement string.

```bash
tools rename-replace -target="old_string" [-replaceWith="new_string"] [directory]
```

**Flags:**
- `-target`: Target string to replace in filenames (required)
- `-replaceWith`: String to replace target with (optional, defaults to empty string to remove target)

**Arguments:**
- `directory`: Directory to process (optional, defaults to current directory)

**Examples:**
```bash
# Replace "test" with "prod" in all filenames in current directory
tools rename-replace -target="test" -replaceWith="prod"

# Remove "backup_" prefix from all files in /path/to/files (replaceWith omitted)
tools rename-replace -target="backup_" /path/to/files

# Explicitly remove "temp_" prefix using empty string
tools rename-replace -target="temp_" -replaceWith=""
```

#### 2. create-rand-files

Creates random text files with random names at a specified directory depth.

```bash
tools create-rand-files -depth=<number> -count=<number> [directory]
```

**Flags:**
- `-depth`: Directory depth for file creation (default: 1)
- `-count`: Number of files to create (default: 5)

**Arguments:**
- `directory`: Base directory to create files in (optional, defaults to current directory)

**Examples:**
```bash
# Create 10 random files in current directory
tools create-rand-files -depth=1 -count=10

# Create 5 random files in nested directories under /tmp
tools create-rand-files -depth=3 -count=5 /tmp
```

#### 3. folderify

Takes files and creates folders with the same name (minus extension), then moves each file into its corresponding folder.

```bash
tools folderify [-recursive] [directory]
```

**Flags:**
- `-recursive`: Process subdirectories recursively (optional)

**Arguments:**
- `directory`: Directory to process (optional, defaults to current directory)

**Examples:**
```bash
# Folderify files in current directory only
tools folderify

# Folderify files recursively in all subdirectories
tools folderify -recursive

# Folderify files in specific directory
tools folderify /path/to/files

# Folderify files recursively in specific directory
tools folderify -recursive /path/to/files
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

## Project Structure

```
filekit/
├── main.go                    # Main entry point and command routing
├── cmd/                       # Command handlers
│   ├── replace_in_names.go   # rename-replace command handler
│   ├── create_rand_files.go  # create-rand-files command handler
│   └── folderify.go          # folderify command handler
├── internal/                  # Internal packages (implementation logic)
│   ├── rename/               # File renaming logic
│   │   └── rename.go
│   ├── generator/            # Random file generation logic
│   │   └── generator.go
│   └── folderify/           # Folderify logic
│       └── folderify.go
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
