# now command
Print current date time.

## Overview
The `now` command is a simple tool that displays the current local time in the specified format.

## Usage
By executing the following command, the current date and time will be displayed in the specified format.

```sh
go run main.go -f "format string"
```

### Format Examples
- `2006-01-02 15:04:05` : 2023-10-05 14:45:30
- `2006-01-02` : 2023-10-05
- `15:04:05` : 14:45:30

## Code Explanation
The `main.go` file contains the code to get the current local time and display it in the specified format.

## Commit Message Template
Use the following template for commit messages:

```
<type>: <subject>

<body>
```

### Types
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools and libraries such as documentation generation

### Example
```
feat: add new time format option

Added a new option to the `now` command to support custom time formats.
