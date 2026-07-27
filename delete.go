package main

import (
	"flag"
	"fmt"
	"os"
)

// runDelete executes the delete subcommand
func runDelete(args []string) error {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
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

	if err := DeleteSecret(op, *name); err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully deleted secret: %s\n", *name)
	return nil
}
