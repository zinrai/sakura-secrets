package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// RunPut executes the put subcommand
func RunPut(args []string) int {
	fs := flag.NewFlagSet("put", flag.ExitOnError)
	zone := fs.String("zone", "is1a", "Zone name (default: is1a)")
	name := fs.String("name", "", "Secret name (required)")

	fs.Parse(args)

	if *name == "" {
		fmt.Fprintln(os.Stderr, "Error: -name is required")
		fs.Usage()
		return 1
	}

	config, err := LoadConfig(*zone)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	secretValue, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		return 1
	}

	if len(secretValue) == 0 {
		fmt.Fprintln(os.Stderr, "Error: no input provided")
		return 1
	}

	if err := CreateSecret(config, *name, string(secretValue)); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	fmt.Printf("Successfully created/updated secret: %s\n", *name)
	return 0
}
