# now command
Display current date time with both CLI and GUI interfaces.

## Overview
The `now` command is a versatile tool that displays the current local time in your specified format. It supports both command-line interface (CLI) and graphical user interface (GUI) modes, and includes timezone information.

## Features
- Flexible time format customization
- GUI mode with real-time updates
- Timezone information display
- Command-line interface for quick access
- Live format preview in GUI mode

## Installation
```sh
# Clone the repository
git clone https://github.com/y16ra/now.git
cd now

# Build the application
make build
```

## Usage

### CLI Mode
Display the current time in CLI mode:
```sh
# Using make
make run

# Direct execution
./bin/now

# With custom format
./bin/now -f "2006-01-02 15:04:05"
```

### GUI Mode
Launch the interactive GUI interface:
```sh
# Using make
make run-gui

# Direct execution
./bin/now -gui

# With custom initial format
./bin/now -gui -f "15:04:05"
```

### Command Line Options
- `-gui`: Launch in GUI mode
- `-f string`: Specify time format (default: "2006-01-02 15:04:05")

### Format Examples
- `2006-01-02 15:04:05` : 2025-02-08 17:00:28 (JST)
- `2006-01-02` : 2025-02-08
- `15:04:05` : 17:00:28
- `2006年01月02日 15時04分05秒` : 2025年02月08日 17時00分28秒

## Build Commands
The project includes a Makefile with the following targets:

- `make build`: Build the application
- `make run`: Run in CLI mode
- `make run-gui`: Run in GUI mode
- `make test`: Run tests
- `make clean`: Clean build artifacts
- `make bundle`: Create macOS application bundle (macOS only)
- `make deps`: Update dependencies
- `make help`: Show available commands

## GUI Features
The GUI interface provides:
- Real-time time display with automatic updates
- Live format preview
- Timezone information display
- Interactive format input field
- Clean and simple interface

## Code Structure
- `main.go`: Main application logic and GUI implementation
- `internal/timeutil`: Time formatting utilities
- `Makefile`: Build and run commands

## Requirements
- Go 1.21 or later
- Fyne GUI toolkit (automatically installed via go modules)

## Platform Support
- Primary support for macOS (includes code signing)
- Should work on other platforms (Linux, Windows) but may require additional setup

## Contributing
Please refer to the commit message template below for contributing to this project.

### Commit Message Template
```
<type>: <subject>

<body>
```

### Types
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools

### Example
```
feat: add GUI mode with timezone display

Added a new GUI mode using Fyne toolkit that displays current time
with timezone information and supports real-time updates.
```
