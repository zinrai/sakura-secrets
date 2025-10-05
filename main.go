package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	subcommand := os.Args[1]
	args := os.Args[2:]

	switch subcommand {
	case "list":
		os.Exit(RunList(args))
	case "put":
		os.Exit(RunPut(args))
	case "delete":
		os.Exit(RunDelete(args))
	default:
		fmt.Fprintf(os.Stderr, "Unknown subcommand: %s\n", subcommand)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: sakura-secrets <subcommand> [options]

Subcommands:
  list    List all secrets in a Vault
  put     Create or update a secret
  delete  Delete a secret

Use "sakura-secrets <subcommand> -h" for more information about a subcommand.
`)
}
