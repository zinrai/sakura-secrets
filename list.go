package main

import (
	"flag"
	"fmt"
)

// runList executes the list subcommand
func runList(args []string) error {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	zone := fs.String("zone", "is1a", "Zone name (default: is1a)")

	fs.Parse(args)

	vaultID, err := LoadVaultID()
	if err != nil {
		return err
	}

	op, err := NewSecretOp(*zone, vaultID)
	if err != nil {
		return err
	}

	secrets, err := ListSecrets(op)
	if err != nil {
		return err
	}

	fmt.Printf("Total: %d secrets\n\n", len(secrets))
	for _, secret := range secrets {
		fmt.Printf("Name: %s (Version: %d)\n", secret.Name, secret.LatestVersion)
	}

	return nil
}
