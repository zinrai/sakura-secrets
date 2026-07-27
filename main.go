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

	var err error
	switch os.Args[1] {
	case "list":
		err = runList(os.Args[2:])
	case "put":
		err = runPut(os.Args[2:])
	case "delete":
		err = runDelete(os.Args[2:])
	case "version":
		runVersion()
	default:
		fmt.Fprintf(os.Stderr, "Unknown subcommand: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: sakura-secrets <subcommand> [options]

Subcommands:
  list     List all secrets in a Vault
  put      Create or update a secret
  delete   Delete a secret
  version  Print version

Use "sakura-secrets <subcommand> -h" for more information about a subcommand.
`)
}
