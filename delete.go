package main

import (
	"flag"
	"fmt"
	"os"
)

// RunDelete executes the delete subcommand
func RunDelete(args []string) int {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
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

	if err := DeleteSecret(config, *name); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	fmt.Printf("Successfully deleted secret: %s\n", *name)
	return 0
}
