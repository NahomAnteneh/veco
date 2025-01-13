# Veco - Delta-based Version Control System

Veco is a modern, delta-based version control system designed for efficient storage and version management of both text and binary files.

## Features

- Delta-based storage for efficient space usage
- Support for both text and binary files
- Branch management and merging
- Remote repository support
- Clean and modular architecture

## Installation

```bash
go install github.com/NahomAnteneh/veco@latest
```

## Commands

- `init`: Initialize a new repository

  ```bash
  veco init
  ```

- `add`: Stage files for commit

  ```bash
  veco add <file>...
  veco add .  # Stage all changes
  ```

- `commit`: Create a new commit

  ```bash
  veco commit -m "commit message"
  ```

- `status`: Show working tree status

  ```bash
  veco status
  ```

- `log`: Show commit history

  ```bash
  veco log
  ```

- `branch`: Manage branches

  ```bash
  veco branch          # List branches
  veco branch <name>   # Create new branch
  ```

- `checkout`: Switch branches or restore files

  ```bash
  veco checkout <branch>
  veco checkout -b <new-branch>
  ```

- `remote`: Manage remote repositories

  ```bash
  veco remote add <name> <url>
  veco remote remove <name>
  ```

- `push`: Push to remote repository

  ```bash
  veco push <remote> <branch>
  ```

- `pull`: Pull from remote repository
  ```bash
  veco pull <remote> <branch>
  ```

## Architecture

Veco uses a delta-based approach to store changes efficiently:

1. **Object Storage**: Files are stored as content-addressed objects
2. **Delta Compression**: Changes are stored as deltas between versions
3. **Branch Management**: Efficient branching with fast switching
4. **Remote Sync**: Optimized remote repository synchronization

## Development

### Building from source

```bash
git clone https://github.com/NahomAnteneh/veco.git
cd veco
go build
```
