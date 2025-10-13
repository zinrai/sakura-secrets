package main

import (
	"flag"
	"fmt"
	"os"
)

// RunList executes the list subcommand
func RunList(args []string) int {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	zone := fs.String("zone", "is1a", "Zone name (default: is1a)")

	fs.Parse(args)

	config, err := LoadConfig(*zone)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	result, err := ListSecrets(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	fmt.Printf("Total: %d secrets\n\n", result.Total)
	for _, secret := range result.Secrets {
		fmt.Printf("Name: %s (Version: %d)\n", secret.Name, secret.LatestVersion)
	}

	return 0
}
