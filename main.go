package main

import (
	"fmt"
	"os"

	"github.com/NahomAnteneh/veco/cmd"
)

type Command struct {
	Name        string
	Description string
	Execute     func([]string) error
}

var commands = map[string]Command{
	"init": {
		Name:        "init",
		Description: "Initialize a new Veco repository",
		Execute:     cmd.CmdInit,
	},
	"add": {
		Name:        "add",
		Description: "Stage files for commit",
		Execute:     cmd.CmdAdd,
	},
	"commit": {
		Name:        "commit",
		Description: "Commit staged changes",
		Execute:     cmd.CmdCommit,
	},
	"status": {
		Name:        "status",
		Description: "Show working tree status",
		Execute:     cmd.CmdStatus,
	},
	"log": {
		Name:        "log",
		Description: "Show commit history",
		Execute:     cmd.CmdLog,
	},
	"diff": {
		Name:        "diff",
		Description: "Show changes between commits",
		Execute:     cmd.CmdDiff,
	},
	"branch": {
		Name:        "branch",
		Description: "List, create, or delete branches",
		Execute:     cmd.CmdBranch,
	},
	"checkout": {
		Name:        "checkout",
		Description: "Switch or create branches",
		Execute:     cmd.CmdCheckout,
	},
	"remote": {
		Name:        "remote",
		Description: "Manage remote repositories",
		Execute:     cmd.CmdRemote,
	},
	"push": {
		Name:        "push",
		Description: "Push changes to remote repository",
		Execute:     cmd.CmdPush,
	},
	"pull": {
		Name:        "pull",
		Description: "Pull changes from remote repository",
		Execute:     cmd.CmdPull,
	},
	"fetch": {
		Name:        "fetch",
		Description: "Download objects from remote repository",
		Execute:     cmd.CmdFetch,
	},
	"merge": {
		Name:        "merge",
		Description: "Merge branches or commits",
		Execute:     cmd.CmdMerge,
	},
	"clone": {
		Name:        "clone",
		Description: "Clone a repository into a new directory",
		Execute:     cmd.CmdClone,
	},
	"reset": {
		Name:        "reset",
		Description: "Reset the working directory",
		Execute:     cmd.CmdReset,
	},
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmd, exists := commands[cmdName]
	if !exists {
		fmt.Printf("Unknown command: %s\n", cmdName)
		printUsage()
		os.Exit(1)
	}

	args := os.Args[2:]
	if err := cmd.Execute(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: veco <command> [arguments]")
	fmt.Println("\nAvailable commands:")
	for _, cmd := range commands {
		fmt.Printf("  %-12s %s\n", cmd.Name, cmd.Description)
	}
}
