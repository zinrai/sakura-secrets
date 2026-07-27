package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// runPut executes the put subcommand
func runPut(args []string) error {
	fs := flag.NewFlagSet("put", flag.ExitOnError)
	zone := fs.String("zone", "is1a", "Zone name (default: is1a)")
	name := fs.String("name", "", "Secret name (required)")

	fs.Parse(args)

	if *name == "" {
		fs.Usage()
		return fmt.Errorf("-name is required")
	}

	vaultID, err := LoadVaultID()
	if err != nil {
		return err
	}

	op, err := NewSecretOp(*zone, vaultID)
	if err != nil {
		return err
	}

	secretValue, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %w", err)
	}

	if len(secretValue) == 0 {
		return fmt.Errorf("no input provided")
	}

	if err := CreateSecret(op, *name, string(secretValue)); err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created/updated secret: %s\n", *name)
	return nil
}
